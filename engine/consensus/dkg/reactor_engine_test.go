package dkg_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/engine/consensus/dkg"
	"github.com/onflow/flow-go/model/encodable"
	"github.com/onflow/flow-go/model/flow"
	dkgmodule "github.com/onflow/flow-go/module/dkg"
	module "github.com/onflow/flow-go/module/mock"
	"github.com/onflow/flow-go/state/protocol/events/gadgets"
	protocol "github.com/onflow/flow-go/state/protocol/mock"
	storage "github.com/onflow/flow-go/storage/mock"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/onflow/flow-go/utils/unittest/mocks"
)

// TestEpochSetup ensures that, upon receiving an EpochSetup event, the engine
// correctly creates a new DKGController and registers phase transitions based
// on the views specified in the current epoch, as well as regular calls to the
// DKG smart-contract.
//
// The EpochSetup event is received at view 100.

// The current epoch is configured with DKG phase transitions at views 150, 200,
// and 250. In between phase transitions, the controller calls the DKG
// smart-contract every 10 views.
//
// VIEWS
// setup      : 100
// polling    : 110 120 130 140 150
// Phase1Final: 150
// polling    : 160 170 180 190 200
// Phase2Final: 200
// polling    : 210 220 230 240 250
// Phase3Final: 250
func TestEpochSetup(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	currentCounter := rand.Uint64()
	nextCounter := currentCounter + 1
	committee := unittest.IdentityListFixture(10)
	myIndex := 5
	me := new(module.Local)
	me.On("NodeID").Return(committee[myIndex].NodeID)

	// create a block for each view of interest
	blocks := make(map[uint64]*flow.Header)
	var view uint64
	for view = 100; view <= 250; view += dkg.DefaultPollStep {
		header := unittest.BlockHeaderFixture()
		header.View = view
		blocks[view] = &header
	}
	firstBlock := blocks[100]

	// expectedPrivKey is the expected private share produced by the dkg run. We
	// will mock the controller to return this value, and we will check it
	// against the value that gets inserted in the DB at the end.
	expectedPrivKey := unittest.NetworkingPrivKeyFixture()

	currentEpoch := new(protocol.Epoch)
	currentEpoch.On("Counter").Return(currentCounter, nil)
	currentEpoch.On("DKGPhase1FinalView").Return(uint64(150), nil)
	currentEpoch.On("DKGPhase2FinalView").Return(uint64(200), nil)
	currentEpoch.On("DKGPhase3FinalView").Return(uint64(250), nil)
	nextEpoch := new(protocol.Epoch)
	nextEpoch.On("Counter").Return(nextCounter, nil)
	nextEpoch.On("InitialIdentities").Return(committee, nil)

	epochQuery := mocks.NewEpochQuery(t, currentCounter)
	epochQuery.Add(currentEpoch)
	epochQuery.Add(nextEpoch)
	snapshot := new(protocol.Snapshot)
	snapshot.On("Epochs").Return(epochQuery)
	state := new(protocol.State)
	state.On("AtBlockID", firstBlock.ID()).Return(snapshot)

	// ensure that an attempt is made to insert the expected dkg private share
	// for the next epoch.
	keyStorage := new(storage.BeaconPrivateKeys)
	keyStorage.On("InsertMyBeaconPrivateKey", mock.Anything, mock.Anything).Run(
		func(args mock.Arguments) {
			epochCounter := args.Get(0).(uint64)
			require.Equal(t, nextCounter, epochCounter)
			dkgPriv := args.Get(1).(*encodable.RandomBeaconPrivKey)
			require.Equal(t, expectedPrivKey, dkgPriv.PrivateKey)
		}).
		Return(nil).
		Once()

	// we will ensure that the controller state transitions get called
	// appropriately
	controller := new(module.DKGController)
	controller.On("Run").Return(nil).Once()
	controller.On("EndPhase1").Return(nil).Once()
	controller.On("EndPhase2").Return(nil).Once()
	controller.On("End").Return(nil).Once()
	controller.On("Poll", mock.Anything).Return(nil).Times(15)
	controller.On("GetArtifacts").Return(expectedPrivKey, nil, nil).Once()
	controller.On("SubmitResult").Return(nil).Once()

	factory := new(module.DKGControllerFactory)
	factory.On("Create",
		dkgmodule.CanonicalInstanceID(firstBlock.ChainID, nextCounter),
		committee,
		mock.Anything,
	).Return(controller, nil)

	loggerCalls := 0
	logger := hookedLogger(&loggerCalls)

	viewEvents := gadgets.NewViews()
	engine := dkg.NewReactorEngine(
		logger,
		me,
		state,
		keyStorage,
		factory,
		viewEvents,
	)

	engine.EpochSetupPhaseStarted(currentCounter, firstBlock)

	for view = 100; view <= 250; view += dkg.DefaultPollStep {
		viewEvents.BlockFinalized(blocks[view])
	}

	// check that the appropriate callbacks were registered
	time.Sleep(50 * time.Millisecond)
	controller.AssertExpectations(t)
	keyStorage.AssertExpectations(t)
	// logger shouldn't be called in the happy path
	require.Equal(t, 0, loggerCalls)
}

// TestReactorEngine_EpochCommittedPhaseStarted ensures that we are logging
// a warning message whenever we have a mismatch between the locally produced DKG keys
// and the keys produced by the DKG smart contract.
func TestReactorEngine_EpochCommittedPhaseStarted(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	currentCounter := rand.Uint64()
	nextCounter := currentCounter + 1
	me := new(module.Local)

	id := unittest.IdentifierFixture()
	me.On("NodeID").Return(id)

	// privKey represents private key generated by DKG for the next epoch.
	privKey := unittest.StakingPrivKeyFixture()

	// priv.PublicKey() will return a public key
	// that does not match the public key for the priv key generated above and cause a warning
	// to be logged.
	priv := unittest.RandomBeaconPriv()

	keyStorage := new(storage.BeaconPrivateKeys)
	keyStorage.On("RetrieveMyBeaconPrivateKey", currentCounter+1).Return(priv, nil)
	factory := new(module.DKGControllerFactory)

	nextDKG := new(protocol.DKG)
	nextDKG.On("KeyShare", id).Return(privKey.PublicKey(), nil)

	currentEpoch := new(protocol.Epoch)
	currentEpoch.On("Counter").Return(currentCounter, nil)

	nextEpoch := new(protocol.Epoch)
	nextEpoch.On("Counter").Return(nextCounter, nil)
	nextEpoch.On("DKG").Return(nextDKG, nil)

	epochQuery := mocks.NewEpochQuery(t, currentCounter)
	epochQuery.Add(currentEpoch)
	epochQuery.Add(nextEpoch)

	firstBlock := unittest.BlockHeaderFixture()
	firstBlock.View = 100

	snapshot := new(protocol.Snapshot)
	snapshot.On("Epochs").Return(epochQuery)

	state := new(protocol.State)
	state.On("Final").Return(snapshot)

	viewEvents := gadgets.NewViews()

	hookCalls := 0
	logger := hookedLogger(&hookCalls)

	engine := dkg.NewReactorEngine(
		logger,
		me,
		state,
		keyStorage,
		factory,
		viewEvents,
	)

	engine.EpochCommittedPhaseStarted(currentCounter, &firstBlock)

	require.Equal(t, 1, hookCalls)
}

// utility function to track the number of calls to a logger
func hookedLogger(calls *int) zerolog.Logger {
	hook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
		if level == zerolog.WarnLevel {
			*calls++
		}
	})
	return zerolog.New(os.Stdout).Level(zerolog.WarnLevel).Hook(hook)
}

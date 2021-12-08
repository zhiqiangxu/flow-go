package verification

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/consensus/hotstuff/mocks"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/consensus/hotstuff/signature"
	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/model/encoding"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/local"
	modulemock "github.com/onflow/flow-go/module/mock"
	storagemock "github.com/onflow/flow-go/storage/mock"
	"github.com/onflow/flow-go/utils/unittest"
)

// Test that when DKG key is available for a view, a signed block can pass the validation
// the sig include both staking sig and random beacon sig.
func TestCombinedSignWithDKGKey(t *testing.T) {
	identities := unittest.IdentityListFixture(4, unittest.WithRole(flow.RoleConsensus))

	// prepare data
	dkgKey := unittest.DKGParticipantPriv()
	dkgKey.NodeID = identities[0].NodeID
	pk := dkgKey.RandomBeaconPrivKey.PublicKey()
	signerID := dkgKey.NodeID
	view := uint64(20)

	fblock := unittest.BlockFixture()
	fblock.Header.ProposerID = signerID
	fblock.Header.View = view
	block := model.BlockFromFlow(fblock.Header, 10)

	epochCounter := uint64(3)
	epochLookup := &modulemock.EpochLookup{}
	epochLookup.On("EpochForViewWithFallback", view).Return(epochCounter, nil)

	keys := &storagemock.DKGKeys{}
	// there is DKG key for this epoch
	keys.On("RetrieveMyDKGPrivateInfo", epochCounter).Return(dkgKey, true, nil)

	beaconKeyStore := signature.NewEpochAwareRandomBeaconKeyStore(epochLookup, keys)

	stakingPriv := unittest.StakingPrivKeyFixture()
	nodeID := unittest.IdentityFixture()
	nodeID.NodeID = signerID
	nodeID.StakingPubKey = stakingPriv.PublicKey()

	me, err := local.New(nodeID, stakingPriv)
	require.NoError(t, err)
	signer := NewCombinedSigner(me, beaconKeyStore)

	dkg := &mocks.DKG{}
	dkg.On("KeyShare", signerID).Return(pk, nil)

	committee := &mocks.Committee{}
	committee.On("DKG", mock.Anything).Return(dkg, nil)

	packer := signature.NewConsensusSigDataPacker(committee)
	verifier := NewCombinedVerifier(committee, packer)

	// check that a created proposal can be verified by a verifier
	proposal, err := signer.CreateProposal(block)
	require.NoError(t, err)

	vote := proposal.ProposerVote()
	valid, err := verifier.VerifyVote(nodeID, vote.SigData, proposal.Block)
	require.NoError(t, err)
	require.Equal(t, true, valid)

	// check that a created proposal's signature is a combined staking sig and random beacon sig
	msg := MakeVoteMessage(block.View, block.BlockID)
	stakingSig, err := stakingPriv.Sign(msg, crypto.NewBLSKMAC(encoding.ConsensusVoteTag))
	require.NoError(t, err)

	beaconSig, err := dkgKey.RandomBeaconPrivKey.Sign(msg, crypto.NewBLSKMAC(encoding.RandomBeaconTag))
	require.NoError(t, err)

	expectedSig := signature.EncodeDoubleSig(stakingSig, beaconSig)
	require.Equal(t, expectedSig, proposal.SigData)

	// vote should be valid
	vote, err = signer.CreateVote(block)
	require.NoError(t, err)

	voteValid, err := verifier.VerifyVote(nodeID, vote.SigData, block)
	require.NoError(t, err)
	require.Equal(t, true, voteValid)

	// vote on different bock should be invalid
	blockWrongID := *block
	blockWrongID.BlockID[0]++
	_, err = verifier.VerifyVote(nodeID, vote.SigData, &blockWrongID)
	require.Error(t, err)

	// vote with a wrong view should be invalid
	blockWrongView := *block
	blockWrongView.View++
	_, err = verifier.VerifyVote(nodeID, vote.SigData, &blockWrongView)
	require.Error(t, err)

	// vote by different signer should be invalid
	wrongVoter := identities[1]
	wrongVoter.StakingPubKey = unittest.StakingPrivKeyFixture().PublicKey()
	_, err = verifier.VerifyVote(wrongVoter, vote.SigData, block)
	require.Error(t, err)

	// vote with changed signature should be invalid
	wrongSigData := *vote
	wrongSigData.SigData[4]++
	_, err = verifier.VerifyVote(nodeID, wrongSigData.SigData, block)
	require.Error(t, err)
}

// Test that when DKG key is not available for a view, a signed block can pass the validation
// the sig only include staking sig
func TestCombinedSignWithNoDKGKey(t *testing.T) {
	// prepare data
	dkgKey := unittest.DKGParticipantPriv()
	pk := dkgKey.RandomBeaconPrivKey.PublicKey()
	signerID := dkgKey.NodeID
	view := uint64(20)

	fblock := unittest.BlockFixture()
	fblock.Header.ProposerID = signerID
	fblock.Header.View = view
	block := model.BlockFromFlow(fblock.Header, 10)

	epochCounter := uint64(3)
	epochLookup := &modulemock.EpochLookup{}
	epochLookup.On("EpochForViewWithFallback", view).Return(epochCounter, nil)

	keys := &storagemock.DKGKeys{}
	// there is no DKG key for this epoch
	keys.On("RetrieveMyDKGPrivateInfo", epochCounter).Return(nil, false, nil)

	beaconKeyStore := signature.NewEpochAwareRandomBeaconKeyStore(epochLookup, keys)

	stakingPriv := unittest.StakingPrivKeyFixture()
	nodeID := unittest.IdentityFixture()
	nodeID.NodeID = signerID
	nodeID.StakingPubKey = stakingPriv.PublicKey()

	me, err := local.New(nodeID, stakingPriv)
	require.NoError(t, err)
	signer := NewCombinedSigner(me, beaconKeyStore)

	dkg := &mocks.DKG{}
	dkg.On("KeyShare", signerID).Return(pk, nil)

	committee := &mocks.Committee{}
	// even if the node failed DKG, and has no random beacon private key,
	// but other nodes, who completed and succeeded DKG, have a public key
	// for this failed node, which can be used to verify signature from
	// this failed node.
	committee.On("DKG", mock.Anything).Return(dkg, nil)

	packer := signature.NewConsensusSigDataPacker(committee)
	verifier := NewCombinedVerifier(committee, packer)

	proposal, err := signer.CreateProposal(block)
	require.NoError(t, err)

	vote := proposal.ProposerVote()
	valid, err := verifier.VerifyVote(nodeID, vote.SigData, proposal.Block)
	require.NoError(t, err)
	require.Equal(t, true, valid)

	// As the proposer does not have a Random Beacon Key, it should sign soley with its staking key.
	// In this case, the SigData should be identical to the staking sig.
	expectedStakingSig, err := stakingPriv.Sign(
		MakeVoteMessage(block.View, block.BlockID),
		crypto.NewBLSKMAC(encoding.ConsensusVoteTag),
	)
	require.NoError(t, err)
	require.Equal(t, expectedStakingSig, crypto.Signature(proposal.SigData))
}

package migrations

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/fxamacker/cbor/v2"

	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/runtime/ast"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/utils"
	"github.com/onflow/flow-go/model/flow"
)

// iterates through registers keeping a map of register sizes
// after it has reached the end it add storage used and storage capacity for each address
func StorageFeesStorageCapacityMigration(payload []ledger.Payload) ([]ledger.Payload, error) {
	l := newLed(payload)
	s := state.NewState(l)

	storageFeesAddress, flowTokenAddress, err := commonAddresses(payload)
	if err != nil {
		return nil, err
	}

	uuidGenerator := fvm.NewUUIDGenerator(state.NewUUIDs(s))
	capacityPayload := make(map[string]ledger.Payload)

	usedFlowTokens := uint64(0)

	for _, p := range payload {
		id, err := keyToRegisterID(p.Key)
		if err != nil {
			return nil, err
		}
		if len([]byte(id.Owner)) != flow.AddressLength {
			// not an address
			continue
		}
		_, ok := capacityPayload[id.Owner]
		if ok {
			// already processed
			continue
		}
		if id.Key != "storage_used" {
			// we need the storage_used register
			continue
		}
		storageUsed, _, err := utils.ReadUint64(p.Value)
		if err != nil {
			return nil, err
		}

		payload, tokens, err := createCapacityPayload(id, storageUsed, storageFeesAddress, flowTokenAddress, *uuidGenerator, nextStorageReservationIDGenerator())
		if err != nil {
			return nil, err
		}
		usedFlowTokens = usedFlowTokens + tokens

		capacityPayload[id.Owner] = payload
	}

	var newPayload []ledger.Payload
	for _, v := range capacityPayload {
		newPayload = append(newPayload, v)
	}

	payload = append(payload, newPayload...)
	return payload, nil
}

func createCapacityPayload(id flow.RegisterID, used uint64,
	storageFeesAddress, flowTokenAddress common.Address,
	uuidGenerator fvm.UUIDGenerator,
	generateStorageReservationID func() interpreter.UInt64Value) (ledger.Payload, uint64, error) {

	reservation := neededFlowReservation(used)

	address := common.BytesToAddress([]byte(id.Owner))

	flowVault, err := createFlowVaultComposite(address, flowTokenAddress, interpreter.NewUFix64ValueWithInteger(reservation), uuidGenerator)
	if err != nil {
		return ledger.Payload{}, 0, err
	}

	uuid, err := uuidGenerator.GenerateUUID()
	if err != nil {
		return ledger.Payload{}, 0, err
	}
	uuidValue := interpreter.UInt64Value(uuid)

	value := &interpreter.CompositeValue{
		Location: runtime.AddressLocation{
			Address: storageFeesAddress,
			Name:    "FlowStorageFees",
		},
		TypeID: ast.TypeID(fmt.Sprintf("A.%s.FlowStorageFees.StorageReservation", storageFeesAddress.Hex())),
		Kind:   common.CompositeKindResource,
		Fields: map[string]interpreter.Value{
			"uuid":                 uuidValue,                      // uint64
			"storageReservationId": generateStorageReservationID(), // uint64
			"ownerAddress":         interpreter.NewAddressValueFromBytes([]byte(id.Owner)),
			"reservedTokens":       flowVault,
		},
		Owner: &address,

		InjectedFields: nil,
		NestedValues:   nil,
		Functions:      nil,
		Destructor:     nil,
	}
	registerID := flow.RegisterID{
		Owner:      string(address.Bytes()),
		Controller: "",
		// StorageReservation resource key. Its the /storage/storageReservation path
		Key: fmt.Sprintf("%s\x1F%s", "storage", "flowStorageReservation"),
	}
	encoded, _, err := interpreter.EncodeValue(value, []string{registerID.Key}, false)
	return ledger.Payload{
		Key:   registerIDToKey(registerID),
		Value: encoded,
	}, reservation, nil
}

func createFlowVaultComposite(address, flowTokenAddress common.Address, amount interpreter.UFix64Value, uuidGenerator fvm.UUIDGenerator) (interpreter.Value, error) {
	uuid, err := uuidGenerator.GenerateUUID()
	if err != nil {
		return nil, err
	}
	uuidValue := interpreter.UInt64Value(uuid)

	return &interpreter.CompositeValue{
		Location: runtime.AddressLocation{
			Address: flowTokenAddress,
			Name:    "FlowToken",
		},
		TypeID: ast.TypeID(fmt.Sprintf("A.%s.FlowToken.Vault", flowTokenAddress.Hex())),
		Kind:   common.CompositeKindResource,
		Fields: map[string]interpreter.Value{
			"uuid":    uuidValue,
			"balance": amount,
		},
		Owner: &address,

		InjectedFields: nil,
		NestedValues:   nil,
		Functions:      nil,
		Destructor:     nil,
	}, nil
}

func nextStorageReservationIDGenerator() func() interpreter.UInt64Value {
	i := math.MaxUint64
	return func() interpreter.UInt64Value {
		defer func() {
			i--
		}()
		return interpreter.UInt64Value(i)
	}
}

func neededFlowReservation(used uint64) uint64 {
	bytesToFlow := uint64(100)  // 1MB per flow
	blockSize := uint64(500000) // 500 kB

	block := (used / blockSize) + 2
	return block * bytesToFlow
}

func commonAddresses(payloads []ledger.Payload) (common.Address, common.Address, error) {
	// find one address and infer the rest
	for _, payload := range payloads {
		id, err := keyToRegisterID(payload.Key)
		if err != nil {
			return common.Address{}, common.Address{}, err
		}
		if len([]byte(id.Owner)) != flow.AddressLength {
			// not an address
			continue
		}
		address := flow.BytesToAddress([]byte(id.Owner))

		if id.Key != state.KeyContractNames {
			continue
		}
		hasContract, err := payloadContainsContract(payload.Value, "FungibleToken")
		if err != nil {
			return common.Address{}, common.Address{}, err
		}
		if !hasContract {
			continue
		}

		if address.HexWithPrefix() == "0x1654653399040a61" {
			// mainnet

			return commonAddressFromHex(flow.Mainnet.Chain().ServiceAddress().HexWithPrefix()),
				commonAddressFromHex("0x1654653399040a61"),
				nil
		} else if address.HexWithPrefix() == "0x9a0766d93b6608b7" {
			// devnet
			return commonAddressFromHex(flow.Testnet.Chain().ServiceAddress().HexWithPrefix()),
				commonAddressFromHex("0x9a0766d93b6608b7"),
				nil
		}
	}
	return common.Address{}, common.Address{}, fmt.Errorf("cannot infer FungibleToken and service address")
}

func commonAddressFromHex(hex string) common.Address {
	a := flow.HexToAddress(hex).Bytes()
	return common.BytesToAddress(a)
}

func payloadContainsContract(value []byte, contract string) (bool, error) {
	identifiers := make([]string, 0)
	if len(value) > 0 {
		buf := bytes.NewReader(value)
		cborDecoder := cbor.NewDecoder(buf)
		err := cborDecoder.Decode(&identifiers)
		if err != nil {
			return false, fmt.Errorf("cannot decode deployed contract names %x: %w", buf, err)
		}
	}
	for _, identifier := range identifiers {
		if contract == identifier {
			return true, nil
		}
	}

	return false, nil
}

package ledger

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/storage/ledger/mtrie"
	"github.com/dapperlabs/flow-go/storage/ledger/wal"
)

type MTrieStorage struct {
	mForest *mtrie.MForest
	wal     *wal.LedgerWAL
	metrics module.LedgerMetrics
}

var maxHeight = 257

// NewMTrieStorage creates a new in-memory trie-backed ledger storage with persistence.
func NewMTrieStorage(dbDir string, cacheSize int, metrics module.LedgerMetrics, reg prometheus.Registerer) (*MTrieStorage, error) {

	w, err := wal.NewWAL(nil, reg, dbDir, cacheSize, maxHeight)

	if err != nil {
		return nil, fmt.Errorf("cannot create LedgerWAL: %w", err)
	}

	mForest, err := mtrie.NewMForest(maxHeight, dbDir, cacheSize, metrics, func(evictedTrie *mtrie.MTrie) error {
		return w.RecordDelete(evictedTrie.RootHash())
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create MForest: %w", err)
	}
	trie := &MTrieStorage{
		mForest: mForest,
		wal:     w,
		metrics: metrics,
	}

	err = w.Replay(
		func(storableNodes []*mtrie.StorableNode, storableTries []*mtrie.StorableTrie) error {
			return trie.mForest.LoadStorables(storableNodes, storableTries)
		},
		func(stateCommitment flow.StateCommitment, keys [][]byte, values [][]byte) error {
			_, err = trie.mForest.Update(keys, values, stateCommitment)
			// _, err := trie.UpdateRegisters(keys, values, stateCommitment)
			return err
		},
		func(stateCommitment flow.StateCommitment) error {
			trie.mForest.RemoveTrie(stateCommitment)
			return nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("cannot restore LedgerWAL: %w", err)
	}

	// TODO update to proper value once https://github.com/dapperlabs/flow-go/pull/3720 is merged
	metrics.ForestApproxMemorySize(0)

	return trie, nil
}

func (f *MTrieStorage) Ready() <-chan struct{} {
	ready := make(chan struct{})
	close(ready)
	return ready
}

func (f *MTrieStorage) Done() <-chan struct{} {
	_ = f.wal.Close()
	done := make(chan struct{})
	close(done)
	return done
}

func (f *MTrieStorage) EmptyStateCommitment() flow.StateCommitment {
	return f.mForest.GetEmptyRootHash()
	// return trie.GetDefaultHashForHeight(f.tree.GetHeight() - 1)
}

// GetRegisters read the values at the given registers at the given flow.StateCommitment
// This is trusted so no proof is generated
func (f *MTrieStorage) GetRegisters(
	registerIDs []flow.RegisterID,
	stateCommitment flow.StateCommitment,
) (
	values []flow.RegisterValue,
	err error,
) {
	start := time.Now()

	values, err = f.mForest.Read(registerIDs, stateCommitment)
	// values, _, err = f.tree.Read(registerIDs, true, stateCommitment)

	f.metrics.ReadValuesNumber(uint64(len(registerIDs)))

	readDuration := time.Since(start)

	f.metrics.ReadDuration(readDuration)

	if len(registerIDs) > 0 {
		durationPerValue := time.Duration(readDuration.Nanoseconds()/int64(len(registerIDs))) * time.Nanosecond
		f.metrics.ReadDurationPerItem(durationPerValue)
	}

	return values, err
}

// UpdateRegisters updates the values at the given registers
// This is trusted so no proof is generated
func (f *MTrieStorage) UpdateRegisters(
	ids []flow.RegisterID,
	values []flow.RegisterValue,
	stateCommitment flow.StateCommitment,
) (
	newStateCommitment flow.StateCommitment,
	err error,
) {
	start := time.Now()

	// TODO: add test case
	if len(ids) != len(values) {
		return nil, fmt.Errorf(
			"length of IDs [%d] does not match values [%d]", len(ids), len(values),
		)
	}

	// TODO: add test case
	if len(ids) == 0 {
		// return current state root unchanged
		return stateCommitment, nil
	}

	f.metrics.UpdateCount()
	f.metrics.UpdateValuesNumber(uint64(len(ids)))

	err = f.wal.RecordUpdate(stateCommitment, ids, values)
	if err != nil {
		return nil, fmt.Errorf("cannot update state, error while writing LedgerWAL: %w", err)
	}

	newStateCommitment, err = f.mForest.Update(ids, values, stateCommitment)
	// newStateCommitment, err = f.tree.Update(ids, values, stateCommitment)
	if err != nil {
		return nil, fmt.Errorf("cannot update state: %w", err)
	}

	// TODO update to proper value once https://github.com/dapperlabs/flow-go/pull/3720 is merged
	f.metrics.ForestApproxMemorySize(0)

	elapsed := time.Since(start)
	f.metrics.UpdateDuration(elapsed)

	if len(ids) > 0 {
		durationPerValue := time.Duration(elapsed.Nanoseconds()/int64(len(ids))) * time.Nanosecond
		f.metrics.UpdateDurationPerItem(durationPerValue)
	}

	return newStateCommitment, nil
}

// GetRegistersWithProof read the values at the given registers at the given flow.StateCommitment
// This is untrusted so a proof is generated
func (f *MTrieStorage) GetRegistersWithProof(
	registerIDs []flow.RegisterID,
	stateCommitment flow.StateCommitment,
) (
	values []flow.RegisterValue,
	proofs []flow.StorageProof,
	err error,
) {

	values, err = f.GetRegisters(registerIDs, stateCommitment)

	// values, _, err = f.tree.Read(registerIDs, true, stateCommitment)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get register values: %w", err)
	}

	batchProof, err := f.mForest.Proofs(registerIDs, stateCommitment)
	// batchProof, err := f.tree.GetBatchProof(registerIDs, stateCommitment)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get proofs: %w", err)
	}

	proofToGo, totalProofLength := mtrie.EncodeBatchProof(batchProof)

	if len(proofToGo) > 0 {
		f.metrics.ProofSize(uint32(totalProofLength / len(proofToGo)))
	}

	return values, proofToGo, err
}

func (f *MTrieStorage) GetRegisterTouches(
	registerIDs []flow.RegisterID,
	stateCommitment flow.StateCommitment,
) (
	[]flow.RegisterTouch,
	error,
) {
	values, proofs, err := f.GetRegistersWithProof(registerIDs, stateCommitment)
	if err != nil {
		return nil, err
	}
	rets := make([]flow.RegisterTouch, 0, len(registerIDs))
	for i, reg := range registerIDs {
		rt := flow.RegisterTouch{
			RegisterID: reg,
			Value:      values[i],
			Proof:      proofs[i],
		}
		rets = append(rets, rt)
	}
	return rets, nil
}

// UpdateRegistersWithProof updates the values at the given registers
// This is untrusted so a proof is generated
func (f *MTrieStorage) UpdateRegistersWithProof(
	ids []flow.RegisterID,
	values []flow.RegisterValue,
	stateCommitment flow.StateCommitment,
) (
	newStateCommitment flow.StateCommitment,
	proofs []flow.StorageProof,
	err error,
) {
	newStateCommitment, err = f.UpdateRegisters(ids, values, stateCommitment)
	if err != nil {
		return nil, nil, err
	}

	_, proofs, err = f.GetRegistersWithProof(ids, newStateCommitment)
	return newStateCommitment, proofs, err
}

// CloseStorage closes the DB
func (f *MTrieStorage) CloseStorage() {
	_ = f.wal.Close()
}

func (f *MTrieStorage) DiskSize() (int64, error) {
	return f.mForest.DiskSize()
}

func (f *MTrieStorage) ForestSize() int {
	return f.mForest.Size()
}

func (f *MTrieStorage) Checkpointer() (*wal.Checkpointer, error) {
	checkpointer, err := f.wal.Checkpointer()
	if err != nil {
		return nil, fmt.Errorf("cannot create checkpointer for compactor: %w", err)
	}
	return checkpointer, nil
}

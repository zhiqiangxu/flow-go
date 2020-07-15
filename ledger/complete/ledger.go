package complete

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/dapperlabs/flow-go/ledger"
	"github.com/dapperlabs/flow-go/ledger/common"
	"github.com/dapperlabs/flow-go/ledger/complete/mtrie"
	"github.com/dapperlabs/flow-go/ledger/complete/mtrie/trie"
	"github.com/dapperlabs/flow-go/ledger/complete/wal"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module"
)

const (
	// RegisterKeySize is the size of a Execution register's key [bytes]
	RegisterKeySize = 32
)

// Ledger (complete) is a fast memory-efficient fork-aware thread-safe trie-based key/value storage.
// Ledger holds an array of registers (key value pairs) and keep tracks of changes over a limited time.
// Each register is referenced by an ID (key) and holds a value (byte slice).
// Ledger provides atomic batched update and read (with or without proofs) operation given a list of keys.
// Every update to the Ledger creates a new state commitment which captures the state of the storage.
// Under the hood, it uses binary merkle tries to generate inclusion and noninclusion proofs.
// Ledger is fork-aware that means any update can applied at any previous statecommitments which forms a tree of tries (forest).
// The forrest is in memory but all changes (e.g. register updates) are captured inside write-ahead-logs for crash recovery reasons.
// In order to limit the memory usage and maintain the performance storage only keeps limited number of
// tries and purge the old ones (LRU-based); in other words Ledger is not designed to be used
// for archival use but make it possible for other software components to reconstruct very old tries using write-ahead logs.
type Ledger struct {
	forest  *mtrie.Forest
	wal     *wal.LedgerWAL
	metrics module.LedgerMetrics
}

const CacheSize = 1000

// NewLedger creates a new in-memory trie-backed ledger storage with persistence.
func NewLedger(dbDir string, capacity int, metrics module.LedgerMetrics, reg prometheus.Registerer) (*Ledger, error) {

	w, err := wal.NewWAL(nil, reg, dbDir, capacity, RegisterKeySize)
	if err != nil {
		return nil, fmt.Errorf("cannot create LedgerWAL: %w", err)
	}

	forest, err := mtrie.NewForest(RegisterKeySize, dbDir, capacity, metrics, func(evictedTrie *trie.MTrie) error {
		return w.RecordDelete(evictedTrie.RootHash())
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create forest: %w", err)
	}
	storage := &Ledger{
		forest:  forest,
		wal:     w,
		metrics: metrics,
	}

	err = w.ReplayOnForest(forest)
	if err != nil {
		return nil, fmt.Errorf("cannot restore LedgerWAL: %w", err)
	}

	// TODO update to proper value once https://github.com/dapperlabs/flow-go/pull/3720 is merged
	metrics.ForestApproxMemorySize(0)

	return storage, nil
}

// Ready implements interface module.ReadyDoneAware
// it starts the EventLoop's internal processing loop.
func (f *Ledger) Ready() <-chan struct{} {
	ready := make(chan struct{})
	close(ready)
	return ready
}

// Done implements interface module.ReadyDoneAware
// it closes all the open write-ahead log files.
func (f *Ledger) Done() <-chan struct{} {
	_ = f.wal.Close()
	done := make(chan struct{})
	close(done)
	return done
}

// EmptyStateCommitment returns the state commitment of an empty ledger
func (f *Ledger) EmptyStateCommitment() flow.StateCommitment {
	return flow.StateCommitment(f.forest.GetEmptyRootHash())
}

// Get read the values of the given keys at the given state commitment
// it returns the values in the same order as given registerIDs and errors (if any)
func (f *Ledger) Get(query *ledger.Query) (values []ledger.Value, err error) {
	start := time.Now()
	paths, err := common.KeysToPaths(query.Keys(), 0)
	if err != nil {
		return nil, err
	}
	trieRead := &ledger.TrieRead{RootHash: ledger.RootHash(query.StateCommitment()), Paths: paths}
	payloads, err := f.forest.Read(trieRead)
	if err != nil {
		return nil, err
	}
	values, err = common.PayloadsToValues(payloads)
	if err != nil {
		return nil, err
	}

	f.metrics.ReadValuesNumber(uint64(len(paths)))
	readDuration := time.Since(start)
	f.metrics.ReadDuration(readDuration)

	if len(paths) > 0 {
		durationPerValue := time.Duration(readDuration.Nanoseconds()/int64(len(paths))) * time.Nanosecond
		f.metrics.ReadDurationPerItem(durationPerValue)
	}

	return values, err
}

// Set updates the ledger given an update
// it returns a new state commitment (state after update) and errors (if any)
func (f *Ledger) Set(update *ledger.Update) (newStateCommitment ledger.StateCommitment, err error) {
	start := time.Now()

	// TODO: add test case
	if update.Size() == 0 {
		// return current state root unchanged
		return update.StateCommitment(), nil
	}

	paths, err := common.KeysToPaths(update.Keys(), 0)
	if err != nil {
		return nil, err
	}

	payloads, err := common.UpdateToPayloads(update)
	if err != nil {
		return nil, err
	}

	f.metrics.UpdateCount()
	f.metrics.UpdateValuesNumber(uint64(len(paths)))

	trieUpdate := &ledger.TrieUpdate{RootHash: ledger.RootHash(update.StateCommitment()), Paths: paths, Payloads: payloads}

	err = f.wal.RecordUpdate(trieUpdate)
	if err != nil {
		return nil, fmt.Errorf("cannot update state, error while writing LedgerWAL: %w", err)
	}

	newRootHash, err := f.forest.Update(trieUpdate)
	if err != nil {
		return nil, fmt.Errorf("cannot update state: %w", err)
	}

	// TODO update to proper value once https://github.com/dapperlabs/flow-go/pull/3720 is merged
	f.metrics.ForestApproxMemorySize(0)

	elapsed := time.Since(start)
	f.metrics.UpdateDuration(elapsed)

	if len(paths) > 0 {
		durationPerValue := time.Duration(elapsed.Nanoseconds()/int64(len(paths))) * time.Nanosecond
		f.metrics.UpdateDurationPerItem(durationPerValue)
	}

	// TODO log info state commitments
	return ledger.StateCommitment(newRootHash), nil
}

// Prove provides proofs for a ledger query and errors (if any)
func (f *Ledger) Prove(query *ledger.Query) (proof ledger.Proof, err error) {

	paths, err := common.KeysToPaths(query.Keys(), 0)
	if err != nil {
		return nil, err
	}

	trieRead := &ledger.TrieRead{RootHash: ledger.RootHash(query.StateCommitment()), Paths: paths}
	batchProof, err := f.forest.Proofs(trieRead)
	if err != nil {
		return nil, fmt.Errorf("could not get proofs: %w", err)
	}

	proofToGo := common.EncodeTrieBatchProof(batchProof)

	if len(paths) > 0 {
		f.metrics.ProofSize(uint32(len(proofToGo) / len(paths)))
	}

	return proofToGo, err
}

// CloseStorage closes the DB
func (f *Ledger) CloseStorage() {
	_ = f.wal.Close()
}

// TODO implement an approximate MemSize method
func (f *Ledger) MemSize() (int64, error) {
	return 0, nil
}

// DiskSize returns the amount of disk space used by the storage (in bytes)
func (f *Ledger) DiskSize() (int64, error) {
	return f.forest.DiskSize()
}

// ForestSize returns the number of tries stored in the forest
func (f *Ledger) ForestSize() int {
	return f.forest.Size()
}

func (f *Ledger) Checkpointer() (*wal.Checkpointer, error) {
	checkpointer, err := f.wal.NewCheckpointer()
	if err != nil {
		return nil, fmt.Errorf("cannot create checkpointer for compactor: %w", err)
	}
	return checkpointer, nil
}

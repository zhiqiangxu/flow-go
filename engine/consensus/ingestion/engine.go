// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package ingestion

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/engine/common/fifoqueue"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/module/component"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/network"
)

// defaultGuaranteeQueueCapacity maximum capacity of pending events queue, everything above will be dropped
const defaultGuaranteeQueueCapacity = 1000

// defaultIngestionEngineWorkers number of goroutines engine will use for processing events
const defaultIngestionEngineWorkers = 3

// Engine represents the ingestion engine, used to funnel collections from a
// cluster of collection nodes to the set of consensus nodes. It represents the
// link between collection nodes and consensus nodes and has a counterpart with
// the same engine ID in the collection node.
type Engine struct {
	*component.ComponentManager
	log               zerolog.Logger         // used to log relevant actions with context
	me                module.Local           // used to access local node information
	con               network.Conduit        // conduit to receive/send guarantees
	core              *Core                  // core logic of processing guarantees
	pendingGuarantees engine.MessageStore    // message store of pending events
	messageHandler    *engine.MessageHandler // message handler for incoming events
}

// New creates a new collection propagation engine.
func New(
	log zerolog.Logger,
	engineMetrics module.EngineMetrics,
	net network.Network,
	me module.Local,
	core *Core,
) (*Engine, error) {

	logger := log.With().Str("ingestion", "engine").Logger()

	guaranteesQueue, err := fifoqueue.NewFifoQueue(
		fifoqueue.WithCapacity(defaultGuaranteeQueueCapacity),
		fifoqueue.WithLengthObserver(func(len int) { core.mempool.MempoolEntries(metrics.ResourceCollectionGuaranteesQueue, uint(len)) }),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create guarantees queue: %w", err)
	}

	pendingGuarantees := &engine.FifoMessageStore{
		FifoQueue: guaranteesQueue,
	}

	handler := engine.NewMessageHandler(
		logger,
		engine.NewNotifier(),
		engine.Pattern{
			Match: func(msg *engine.Message) bool {
				_, ok := msg.Payload.(*flow.CollectionGuarantee)
				if ok {
					engineMetrics.MessageReceived(metrics.EngineConsensusIngestion, metrics.MessageCollectionGuarantee)
				}
				return ok
			},
			Store: pendingGuarantees,
		},
	)

	// initialize the propagation engine with its dependencies
	e := &Engine{
		log:               logger,
		me:                me,
		core:              core,
		pendingGuarantees: pendingGuarantees,
		messageHandler:    handler,
	}

	componentManagerBuilder := component.NewComponentManagerBuilder()

	for i := 0; i < defaultIngestionEngineWorkers; i++ {
		componentManagerBuilder.AddWorker(func(ctx irrecoverable.SignalerContext, ready component.ReadyFunc) {
			ready()
			err := e.loop(ctx)
			if err != nil {
				ctx.Throw(err)
			}
		})
	}

	e.ComponentManager = componentManagerBuilder.Build()

	// register the engine with the network layer and store the conduit
	con, err := net.Register(engine.ReceiveGuarantees, e)
	if err != nil {
		return nil, fmt.Errorf("could not register engine: %w", err)
	}
	e.con = con
	return e, nil
}

// SubmitLocal submits an event originating on the local node.
func (e *Engine) SubmitLocal(event interface{}) {
	err := e.ProcessLocal(event)
	if err != nil {
		e.log.Fatal().Err(err).Msg("internal error processing event")
	}
}

// Submit submits the given event from the node with the given origin ID
// for processing in a non-blocking manner. It returns instantly and logs
// a potential processing error internally when done.
func (e *Engine) Submit(channel network.Channel, originID flow.Identifier, event interface{}) {
	err := e.Process(channel, originID, event)
	if err != nil {
		e.log.Fatal().Err(err).Msg("internal error processing event")
	}
}

// ProcessLocal processes an event originating on the local node.
func (e *Engine) ProcessLocal(event interface{}) error {
	return e.messageHandler.Process(e.me.NodeID(), event)
}

// Process processes the given event from the node with the given origin ID in
// a blocking manner. It returns error only in unexpected scenario.
func (e *Engine) Process(_ network.Channel, originID flow.Identifier, event interface{}) error {
	return e.messageHandler.Process(originID, event)
}

// processAvailableMessages processes the given ingestion engine event.
func (e *Engine) processAvailableMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default: // fall through to business logic
		}

		msg, ok := e.pendingGuarantees.Get()
		if ok {
			originID := msg.OriginID
			err := e.core.OnGuarantee(originID, msg.Payload.(*flow.CollectionGuarantee))
			if err != nil {
				if engine.IsInvalidInputError(err) {
					e.log.Error().Str("origin", originID.String()).Err(err).Msg("received invalid collection guarantee")
					return nil
				}
				if engine.IsOutdatedInputError(err) {
					e.log.Warn().Str("origin", originID.String()).Err(err).Msg("received outdated collection guarantee")
					return nil
				}
				if engine.IsUnverifiableInputError(err) {
					e.log.Warn().Str("origin", originID.String()).Err(err).Msg("received unverifiable collection guarantee")
					return nil
				}
				return fmt.Errorf("processing collection guarantee unexpected err: %w", err)
			}

			continue
		}

		// when there is no more messages in the queue, back to the loop to wait
		// for the next incoming message to arrive.
		return nil
	}
<<<<<<< HEAD

	// NOTE: there are two ways to go about this:
	// - expect the collection nodes to propagate the guarantee to all consensus nodes;
	// - ensure that we take care of propagating guarantees to other consensus nodes.
	// Currently, we go with first option as each collection node broadcasts a guarantee to
	// all consensus nodes. So we expect all collections of a cluster to broadcast a guarantee to
	// all consensus nodes. Even on an unhappy path, as long as only one collection node does it
	// the guarantee must be delivered to all consensus nodes.

	return nil
}

// validateExpiry validates that the collection has not expired w.r.t. the local
// latest finalized block.
func (e *Engine) validateExpiry(guarantee *flow.CollectionGuarantee) error {

	// get the last finalized header and the reference block header
	final, err := e.state.Final().Head()
	if err != nil {
		return fmt.Errorf("could not get finalized header: %w", err)
	}
	ref, err := e.headers.ByBlockID(guarantee.ReferenceBlockID)
	if errors.Is(err, storage.ErrNotFound) {
		return engine.NewUnverifiableInputError("collection guarantee refers to an unknown block (id=%x): %w", guarantee.ReferenceBlockID, err)
	}

	// if head has advanced beyond the block referenced by the collection guarantee by more than 'expiry' number of blocks,
	// then reject the collection
	if ref.Height > final.Height {
		return nil // the reference block is newer than the latest finalized one
	}
	if final.Height-ref.Height > flow.DefaultTransactionExpiry {
		return engine.NewOutdatedInputErrorf("collection guarantee expired ref_height=%d final_height=%d", ref.Height, final.Height)
	}

	return nil
}

// validateGuarantors validates that the guarantors of a collection are valid,
// in that they are all from the same cluster and that cluster is allowed to
// produce the given collection w.r.t. the guarantee's reference block.
func (e *Engine) validateGuarantors(guarantee *flow.CollectionGuarantee) error {

	guarantors := guarantee.SignerIDs

	if len(guarantors) == 0 {
		return engine.NewInvalidInputError("invalid collection guarantee with no guarantors")
	}

	// get the clusters to assign the guarantee and check if the guarantor is part of it
	snapshot := e.state.AtBlockID(guarantee.ReferenceBlockID)
	clusters, err := snapshot.Epochs().Current().Clustering()
	if errors.Is(err, storage.ErrNotFound) {
		return engine.NewUnverifiableInputError("could not get clusters for unknown reference block (id=%x): %w", guarantee.ReferenceBlockID, err)
	}
	if err != nil {
		return fmt.Errorf("internal error retrieving collector clusters: %w", err)
	}
	cluster, _, ok := clusters.ByNodeID(guarantors[0])
	if !ok {
		return engine.NewInvalidInputErrorf("guarantor (id=%s) does not exist in any cluster", guarantors[0])
	}

	// NOTE: Eventually we should check the signatures, ensure a quorum of the
	// cluster, and ensure HotStuff finalization rules. Likely a cluster-specific
	// version of the follower will be a good fit for this. For now, collection
	// nodes independently decide when a collection is finalized and we only check
	// that the guarantors are all from the same cluster.

	// ensure the guarantors are from the same cluster
	clusterLookup := cluster.Lookup()
	for _, guarantorID := range guarantors {
		_, exists := clusterLookup[guarantorID]
		if !exists {
			return engine.NewInvalidInputError("inconsistent guarantors from different clusters")
		}
	}

	return nil
=======
>>>>>>> 02def6ea5f686f5a6c5cfddcc230cc3e66e1d802
}

func (e *Engine) loop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-e.messageHandler.GetNotifier():
			err := e.processAvailableMessages(ctx)
			if err != nil {
				return fmt.Errorf("internal error processing queued message: %w", err)
			}
		}
	}
}

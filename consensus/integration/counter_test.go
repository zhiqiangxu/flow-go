package integration_test

import (
	"time"

	"github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/notifications"
	"github.com/rs/zerolog"
)

type CounterConsumer struct {
	notifications.NoopConsumer
	log       zerolog.Logger
	interval  time.Duration
	next      time.Time
	rate      uint
	total     uint
	finalized func(uint)
}

func NewCounterConsumer(log zerolog.Logger, finalized func(uint)) *CounterConsumer {
	cc := CounterConsumer{
		log:       log,
		interval:  time.Second,
		next:      time.Now().UTC().Add(time.Second),
		rate:      0,
		total:     0,
		finalized: finalized,
	}
	return &cc
}

func (c *CounterConsumer) OnFinalizedBlock(block *model.Block) {

	// count the finalized block
	c.rate++
	c.total++

	// notify stopper of total finalized
	c.finalized(c.total)

	// if we are still before the next printing, skip rest
	now := time.Now().UTC()
	if now.Before(c.next) {
		return
	}

	// otherwise, print number of finalized blocks and reset
	c.log.Info().Dur("interval", c.interval).Uint("rate", c.rate).Uint("total", c.total).Msg("blocks per second")
	c.next = now.Add(c.interval)
	c.rate = 0
}

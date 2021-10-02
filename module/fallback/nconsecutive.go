package fallback

import (
	"fmt"
)

type NConsecutiveStrategy struct {
	nClients            int // # of clients available to select
	currentClientIndex  int // index of currently selected client
	consecutiveFailures int // count of failures without a success
	fallbackAfter       int // config - how many failed attempts we allow before selecting another client
}

func NewNConsecutiveStrategy(nClients int, fallbackAfter int) (*NConsecutiveStrategy, error) {
	if nClients <= 0 {
		return nil, fmt.Errorf("must have >0 clients")
	}
	if fallbackAfter <= 0 {
		return nil, fmt.Errorf("must allow at least one failure before fallback")
	}

	strategy := &NConsecutiveStrategy{
		nClients:            nClients,
		currentClientIndex:  0,
		consecutiveFailures: 0,
		fallbackAfter:       fallbackAfter,
	}
	return strategy, nil
}

func (strategy *NConsecutiveStrategy) ClientIndex() int {
	if strategy.consecutiveFailures >= strategy.fallbackAfter {
		strategy.currentClientIndex = (strategy.currentClientIndex + 1) % strategy.nClients
	}
	return strategy.currentClientIndex
}

func (strategy *NConsecutiveStrategy) Success(_ int) {
	strategy.consecutiveFailures = 0
}

func (strategy *NConsecutiveStrategy) Failure(_ int) {
	strategy.consecutiveFailures++
}

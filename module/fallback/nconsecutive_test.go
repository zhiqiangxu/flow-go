package fallback

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNConsecutiveStrategy(t *testing.T) {

	// must have >0 clients
	_, err := NewNConsecutiveStrategy(0, 3)
	require.Error(t, err)

	// must allow at least one attempt before fallback
	_, err = NewNConsecutiveStrategy(2, 0)
	require.Error(t, err)

	// should be able to instantiate with valid args
	_, err = NewNConsecutiveStrategy(3, 5)
	require.NoError(t, err)
}

func TestSuccess(t *testing.T) {

	// should maintain the same client index after any number of sucesses,
	// regardless of fallbackAfter value

	strategy, err := NewNConsecutiveStrategy(3, 2)
	require.NoError(t, err)

	// should start with client index 0
	firstClientIndex := strategy.ClientIndex()
	assert.Equal(t, 0, firstClientIndex)

	for i := 0; i < 100; i++ {
		clientIndex := strategy.ClientIndex()
		assert.Equal(t, firstClientIndex, clientIndex)
		strategy.Success(clientIndex)
	}
}

func TestFailure(t *testing.T) {

}

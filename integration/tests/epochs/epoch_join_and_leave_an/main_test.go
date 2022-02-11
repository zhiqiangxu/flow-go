package epoch_join_and_leave_an

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onflow/flow-go/integration/tests/epochs"
	"github.com/onflow/flow-go/utils/unittest"
)

func TestEpoch(t *testing.T) {
	unittest.SkipUnless(t, unittest.TEST_RESOURCE_INTENSIVE, "epochs join/leave tests should be run on an machine with adequate resources")
	suite.Run(t, new(epochs.EpochJoinAndLeaveANSuite))
}

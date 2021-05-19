// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PingMetrics is an autogenerated mock type for the PingMetrics type
type PingMetrics struct {
	mock.Mock
}

// NodeReachable provides a mock function with given fields: node, nodeInfo, rtt, version, sealedHeight
func (_m *PingMetrics) NodeReachable(node *flow.Identity, nodeInfo string, rtt time.Duration, version string, sealedHeight uint64) {
	_m.Called(node, nodeInfo, rtt, version, sealedHeight)
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	context "context"

	access "github.com/onflow/flow-go/access"

	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// ExecuteScriptAtBlockHeight provides a mock function with given fields: ctx, blockHeight, script, arguments
func (_m *API) ExecuteScriptAtBlockHeight(ctx context.Context, blockHeight uint64, script []byte, arguments [][]byte) ([]byte, error) {
	ret := _m.Called(ctx, blockHeight, script, arguments)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte, [][]byte) []byte); ok {
		r0 = rf(ctx, blockHeight, script, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, []byte, [][]byte) error); ok {
		r1 = rf(ctx, blockHeight, script, arguments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScriptAtBlockID provides a mock function with given fields: ctx, blockID, script, arguments
func (_m *API) ExecuteScriptAtBlockID(ctx context.Context, blockID flow.Identifier, script []byte, arguments [][]byte) ([]byte, error) {
	ret := _m.Called(ctx, blockID, script, arguments)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier, []byte, [][]byte) []byte); ok {
		r0 = rf(ctx, blockID, script, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier, []byte, [][]byte) error); ok {
		r1 = rf(ctx, blockID, script, arguments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScriptAtLatestBlock provides a mock function with given fields: ctx, script, arguments
func (_m *API) ExecuteScriptAtLatestBlock(ctx context.Context, script []byte, arguments [][]byte) ([]byte, error) {
	ret := _m.Called(ctx, script, arguments)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, [][]byte) []byte); ok {
		r0 = rf(ctx, script, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, [][]byte) error); ok {
		r1 = rf(ctx, script, arguments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: ctx, address
func (_m *API) GetAccount(ctx context.Context, address flow.Address) (*flow.Account, error) {
	ret := _m.Called(ctx, address)

	var r0 *flow.Account
	if rf, ok := ret.Get(0).(func(context.Context, flow.Address) *flow.Account); ok {
		r0 = rf(ctx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Address) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtBlockHeight provides a mock function with given fields: ctx, address, height
func (_m *API) GetAccountAtBlockHeight(ctx context.Context, address flow.Address, height uint64) (*flow.Account, error) {
	ret := _m.Called(ctx, address, height)

	var r0 *flow.Account
	if rf, ok := ret.Get(0).(func(context.Context, flow.Address, uint64) *flow.Account); ok {
		r0 = rf(ctx, address, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Address, uint64) error); ok {
		r1 = rf(ctx, address, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtLatestBlock provides a mock function with given fields: ctx, address
func (_m *API) GetAccountAtLatestBlock(ctx context.Context, address flow.Address) (*flow.Account, error) {
	ret := _m.Called(ctx, address)

	var r0 *flow.Account
	if rf, ok := ret.Get(0).(func(context.Context, flow.Address) *flow.Account); ok {
		r0 = rf(ctx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Address) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHeight provides a mock function with given fields: ctx, height
func (_m *API) GetBlockByHeight(ctx context.Context, height uint64) (*flow.Block, error) {
	ret := _m.Called(ctx, height)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *flow.Block); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByID provides a mock function with given fields: ctx, id
func (_m *API) GetBlockByID(ctx context.Context, id flow.Identifier) (*flow.Block, error) {
	ret := _m.Called(ctx, id)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *flow.Block); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeaderByHeight provides a mock function with given fields: ctx, height
func (_m *API) GetBlockHeaderByHeight(ctx context.Context, height uint64) (*flow.Header, error) {
	ret := _m.Called(ctx, height)

	var r0 *flow.Header
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *flow.Header); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeaderByID provides a mock function with given fields: ctx, id
func (_m *API) GetBlockHeaderByID(ctx context.Context, id flow.Identifier) (*flow.Header, error) {
	ret := _m.Called(ctx, id)

	var r0 *flow.Header
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *flow.Header); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionByID provides a mock function with given fields: ctx, id
func (_m *API) GetCollectionByID(ctx context.Context, id flow.Identifier) (*flow.LightCollection, error) {
	ret := _m.Called(ctx, id)

	var r0 *flow.LightCollection
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *flow.LightCollection); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.LightCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForBlockIDs provides a mock function with given fields: ctx, eventType, blockIDs
func (_m *API) GetEventsForBlockIDs(ctx context.Context, eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error) {
	ret := _m.Called(ctx, eventType, blockIDs)

	var r0 []flow.BlockEvents
	if rf, ok := ret.Get(0).(func(context.Context, string, []flow.Identifier) []flow.BlockEvents); ok {
		r0 = rf(ctx, eventType, blockIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.BlockEvents)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []flow.Identifier) error); ok {
		r1 = rf(ctx, eventType, blockIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForHeightRange provides a mock function with given fields: ctx, eventType, startHeight, endHeight
func (_m *API) GetEventsForHeightRange(ctx context.Context, eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error) {
	ret := _m.Called(ctx, eventType, startHeight, endHeight)

	var r0 []flow.BlockEvents
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64) []flow.BlockEvents); ok {
		r0 = rf(ctx, eventType, startHeight, endHeight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.BlockEvents)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, uint64) error); ok {
		r1 = rf(ctx, eventType, startHeight, endHeight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionResultForBlockID provides a mock function with given fields: ctx, blockID
func (_m *API) GetExecutionResultForBlockID(ctx context.Context, blockID flow.Identifier) (*flow.ExecutionResult, error) {
	ret := _m.Called(ctx, blockID)

	var r0 *flow.ExecutionResult
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *flow.ExecutionResult); ok {
		r0 = rf(ctx, blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.ExecutionResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(ctx, blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlock provides a mock function with given fields: ctx, isSealed
func (_m *API) GetLatestBlock(ctx context.Context, isSealed bool) (*flow.Block, error) {
	ret := _m.Called(ctx, isSealed)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(context.Context, bool) *flow.Block); ok {
		r0 = rf(ctx, isSealed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, isSealed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlockHeader provides a mock function with given fields: ctx, isSealed
func (_m *API) GetLatestBlockHeader(ctx context.Context, isSealed bool) (*flow.Header, error) {
	ret := _m.Called(ctx, isSealed)

	var r0 *flow.Header
	if rf, ok := ret.Get(0).(func(context.Context, bool) *flow.Header); ok {
		r0 = rf(ctx, isSealed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, isSealed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestProtocolStateSnapshot provides a mock function with given fields: ctx
func (_m *API) GetLatestProtocolStateSnapshot(ctx context.Context) ([]byte, error) {
	ret := _m.Called(ctx)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context) []byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkParameters provides a mock function with given fields: ctx
func (_m *API) GetNetworkParameters(ctx context.Context) access.NetworkParameters {
	ret := _m.Called(ctx)

	var r0 access.NetworkParameters
	if rf, ok := ret.Get(0).(func(context.Context) access.NetworkParameters); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(access.NetworkParameters)
	}

	return r0
}

// GetTransaction provides a mock function with given fields: ctx, id
func (_m *API) GetTransaction(ctx context.Context, id flow.Identifier) (*flow.TransactionBody, error) {
	ret := _m.Called(ctx, id)

	var r0 *flow.TransactionBody
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *flow.TransactionBody); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.TransactionBody)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionResult provides a mock function with given fields: ctx, id
func (_m *API) GetTransactionResult(ctx context.Context, id flow.Identifier) (*access.TransactionResult, error) {
	ret := _m.Called(ctx, id)

	var r0 *access.TransactionResult
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *access.TransactionResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *API) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *API) SendTransaction(ctx context.Context, tx *flow.TransactionBody) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *flow.TransactionBody) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

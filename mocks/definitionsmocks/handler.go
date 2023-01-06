// Code generated by mockery v2.16.0. DO NOT EDIT.

package definitionsmocks

import (
	context "context"

	definitions "github.com/hyperledger/firefly/internal/definitions"
	core "github.com/hyperledger/firefly/pkg/core"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// HandleDefinitionBroadcast provides a mock function with given fields: ctx, state, msg, data, tx
func (_m *Handler) HandleDefinitionBroadcast(ctx context.Context, state *core.BatchState, msg *core.Message, data core.DataArray, tx *fftypes.UUID) (definitions.HandlerResult, error) {
	ret := _m.Called(ctx, state, msg, data, tx)

	var r0 definitions.HandlerResult
	if rf, ok := ret.Get(0).(func(context.Context, *core.BatchState, *core.Message, core.DataArray, *fftypes.UUID) definitions.HandlerResult); ok {
		r0 = rf(ctx, state, msg, data, tx)
	} else {
		r0 = ret.Get(0).(definitions.HandlerResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.BatchState, *core.Message, core.DataArray, *fftypes.UUID) error); ok {
		r1 = rf(ctx, state, msg, data, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

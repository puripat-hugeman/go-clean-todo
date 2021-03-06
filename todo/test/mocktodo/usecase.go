// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocktodo

import (
	context "context"

	datamodel "github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: ctx, _a1
func (_m *UseCase) CreateTodo(ctx context.Context, _a1 datamodel.TodoRequestEntity) (*datamodel.TodoCreateEntity, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *datamodel.TodoCreateEntity
	if rf, ok := ret.Get(0).(func(context.Context, datamodel.TodoRequestEntity) *datamodel.TodoCreateEntity); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datamodel.TodoCreateEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, datamodel.TodoRequestEntity) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodos provides a mock function with given fields: ctx
func (_m *UseCase) GetTodos(ctx context.Context) ([]datamodel.TodoGetEntity, error) {
	ret := _m.Called(ctx)

	var r0 []datamodel.TodoGetEntity
	if rf, ok := ret.Get(0).(func(context.Context) []datamodel.TodoGetEntity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]datamodel.TodoGetEntity)
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

// NewUseCase creates a new instance of UseCase. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t testing.TB) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

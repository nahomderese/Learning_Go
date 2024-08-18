// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: c, id
func (_m *TaskRepository) Delete(c context.Context, id primitive.ObjectID) error {
	ret := _m.Called(c, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: c, user
func (_m *TaskRepository) FindAll(c context.Context, user domain.User) ([]domain.Task, error) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) ([]domain.Task, error)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) []domain.Task); ok {
		r0 = rf(c, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(c, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: c, id
func (_m *TaskRepository) FindByID(c context.Context, id primitive.ObjectID) (domain.Task, error) {
	ret := _m.Called(c, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (domain.Task, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Task); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: c, task
func (_m *TaskRepository) Save(c context.Context, task domain.Task) (domain.Task, error) {
	ret := _m.Called(c, task)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) (domain.Task, error)); ok {
		return rf(c, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) domain.Task); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Task) error); ok {
		r1 = rf(c, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

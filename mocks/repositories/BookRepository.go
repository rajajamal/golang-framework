// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "github.com/rajajamal/golang-framework/models"
	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: model
func (_m *BookRepository) Save(model *models.Book) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Book) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBookRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookRepository creates a new instance of BookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookRepository(t mockConstructorTestingTNewBookRepository) *BookRepository {
	mock := &BookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

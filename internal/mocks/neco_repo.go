// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	models "github.com/egawata/mockery_test/models"
	mock "github.com/stretchr/testify/mock"
)

// NecoRepo is an autogenerated mock type for the NecoRepo type
type NecoRepo struct {
	mock.Mock
}

// GetNeco provides a mock function with given fields: id
func (_m *NecoRepo) GetNeco(id int) *models.Neco {
	ret := _m.Called(id)

	var r0 *models.Neco
	if rf, ok := ret.Get(0).(func(int) *models.Neco); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Neco)
		}
	}

	return r0
}

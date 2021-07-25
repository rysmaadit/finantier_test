// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	encryption_service_wrapper "github.com/rysmaadit/finantier_test/stock_service/external/encryption_service_wrapper"
	mock "github.com/stretchr/testify/mock"

	polygon "github.com/rysmaadit/finantier_test/stock_service/external/polygon"
)

// EncryptionServiceWrapperInterface is an autogenerated mock type for the EncryptionServiceWrapperInterface type
type EncryptionServiceWrapperInterface struct {
	mock.Mock
}

// Encrypt provides a mock function with given fields: payload
func (_m *EncryptionServiceWrapperInterface) Encrypt(payload *polygon.GetDailyTimeSeriesStockResponse) (*encryption_service_wrapper.EncryptedResponseContract, error) {
	ret := _m.Called(payload)

	var r0 *encryption_service_wrapper.EncryptedResponseContract
	if rf, ok := ret.Get(0).(func(*polygon.GetDailyTimeSeriesStockResponse) *encryption_service_wrapper.EncryptedResponseContract); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*encryption_service_wrapper.EncryptedResponseContract)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*polygon.GetDailyTimeSeriesStockResponse) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

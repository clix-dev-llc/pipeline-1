// Code generated by mockery v1.0.0. DO NOT EDIT.

package commonadapter

import (
	mock "github.com/stretchr/testify/mock"

	secret "github.com/banzaicloud/pipeline/src/secret"
)

// MockReadWriteOrganizationalSecretStore is an autogenerated mock type for the ReadWriteOrganizationalSecretStore type
type MockReadWriteOrganizationalSecretStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: organizationID, secretID
func (_m *MockReadWriteOrganizationalSecretStore) Delete(organizationID uint, secretID string) error {
	ret := _m.Called(organizationID, secretID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(organizationID, secretID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: organizationID, secretID
func (_m *MockReadWriteOrganizationalSecretStore) Get(organizationID uint, secretID string) (*secret.SecretItemResponse, error) {
	ret := _m.Called(organizationID, secretID)

	var r0 *secret.SecretItemResponse
	if rf, ok := ret.Get(0).(func(uint, string) *secret.SecretItemResponse); ok {
		r0 = rf(organizationID, secretID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secret.SecretItemResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(organizationID, secretID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: organizationID, name
func (_m *MockReadWriteOrganizationalSecretStore) GetByName(organizationID uint, name string) (*secret.SecretItemResponse, error) {
	ret := _m.Called(organizationID, name)

	var r0 *secret.SecretItemResponse
	if rf, ok := ret.Get(0).(func(uint, string) *secret.SecretItemResponse); ok {
		r0 = rf(organizationID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secret.SecretItemResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(organizationID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: organizationID, request
func (_m *MockReadWriteOrganizationalSecretStore) Store(organizationID uint, request *secret.CreateSecretRequest) (string, error) {
	ret := _m.Called(organizationID, request)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, *secret.CreateSecretRequest) string); ok {
		r0 = rf(organizationID, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *secret.CreateSecretRequest) error); ok {
		r1 = rf(organizationID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

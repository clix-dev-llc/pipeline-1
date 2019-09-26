// Code generated by mockery v1.0.0. DO NOT EDIT.

package token

import mock "github.com/stretchr/testify/mock"

// MockGenerator is an autogenerated mock type for the Generator type
type MockGenerator struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: sub, expiresAt, tokenType, value
func (_m *MockGenerator) GenerateToken(sub string, expiresAt int64, tokenType string, value string) (string, string, error) {
	ret := _m.Called(sub, expiresAt, tokenType, value)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int64, string, string) string); ok {
		r0 = rf(sub, expiresAt, tokenType, value)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, int64, string, string) string); ok {
		r1 = rf(sub, expiresAt, tokenType, value)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int64, string, string) error); ok {
		r2 = rf(sub, expiresAt, tokenType, value)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
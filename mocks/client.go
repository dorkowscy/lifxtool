// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	lifx "github.com/dorkowscy/lyslix/lifx"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Client) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetColor provides a mock function with given fields: color, fadeTime
func (_m *Client) SetColor(color lifx.HBSK, fadeTime time.Duration) error {
	ret := _m.Called(color, fadeTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(lifx.HBSK, time.Duration) error); ok {
		r0 = rf(color, fadeTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetColorZones provides a mock function with given fields: color, start, end, fadeTime
func (_m *Client) SetColorZones(color lifx.HBSK, start uint8, end uint8, fadeTime time.Duration) error {
	ret := _m.Called(color, start, end, fadeTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(lifx.HBSK, uint8, uint8, time.Duration) error); ok {
		r0 = rf(color, start, end, fadeTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

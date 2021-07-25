// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	jwt "github.com/dgrijalva/jwt-go"
	contract "github.com/rysmaadit/finantier_test/auth_service/contract"

	mock "github.com/stretchr/testify/mock"
)

// JWTClientInterface is an autogenerated mock type for the JWTClientInterface type
type JWTClientInterface struct {
	mock.Mock
}

// GenerateTokenStringWithClaims provides a mock function with given fields: jwtClaims, secret
func (_m *JWTClientInterface) GenerateTokenStringWithClaims(jwtClaims contract.JWTMapClaim, secret string) (string, error) {
	ret := _m.Called(jwtClaims, secret)

	var r0 string
	if rf, ok := ret.Get(0).(func(contract.JWTMapClaim, string) string); ok {
		r0 = rf(jwtClaims, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contract.JWTMapClaim, string) error); ok {
		r1 = rf(jwtClaims, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseTokenWithClaims provides a mock function with given fields: tokenString, claims, secret
func (_m *JWTClientInterface) ParseTokenWithClaims(tokenString string, claims jwt.MapClaims, secret string) error {
	ret := _m.Called(tokenString, claims, secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, jwt.MapClaims, string) error); ok {
		r0 = rf(tokenString, claims, secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

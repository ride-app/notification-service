// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ride-app/notification-service/internal/repositories/token (interfaces: TokenRepository)

// Package mock_token is a generated GoMock package.
package mock_token

import (
	context "context"
	reflect "reflect"

	logger "github.com/dragonfish/go/pkg/logger"
	gomock "github.com/golang/mock/gomock"
)

// MockTokenRepository is a mock of TokenRepository interface.
type MockTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRepositoryMockRecorder
}

// MockTokenRepositoryMockRecorder is the mock recorder for MockTokenRepository.
type MockTokenRepositoryMockRecorder struct {
	mock *MockTokenRepository
}

// NewMockTokenRepository creates a new mock instance.
func NewMockTokenRepository(ctrl *gomock.Controller) *MockTokenRepository {
	mock := &MockTokenRepository{ctrl: ctrl}
	mock.recorder = &MockTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenRepository) EXPECT() *MockTokenRepositoryMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockTokenRepository) GetToken(arg0 context.Context, arg1 logger.Logger, arg2 string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
// GetToken indicates an expected call of GetToken.
func (mr *MockTokenRepositoryMockRecorder) GetToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokenRepository)(nil).GetToken), arg0, arg1, arg2)
}

// UpdateToken mocks base method.
func (m *MockTokenRepository) UpdateToken(arg0 context.Context, arg1 logger.Logger, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockTokenRepositoryMockRecorder) UpdateToken(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockTokenRepository)(nil).UpdateToken), arg0, arg1, arg2, arg3)
}

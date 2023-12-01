// Code generated by MockGen. DO NOT EDIT.
// Source: link.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateShortLink mocks base method.
func (m *MockRepository) CreateShortLink(ctx context.Context, fullLink, shortLink string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortLink", ctx, fullLink, shortLink)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateShortLink indicates an expected call of CreateShortLink.
func (mr *MockRepositoryMockRecorder) CreateShortLink(ctx, fullLink, shortLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortLink", reflect.TypeOf((*MockRepository)(nil).CreateShortLink), ctx, fullLink, shortLink)
}

// GetFullLink mocks base method.
func (m *MockRepository) GetFullLink(ctx context.Context, shortLink string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullLink", ctx, shortLink)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullLink indicates an expected call of GetFullLink.
func (mr *MockRepositoryMockRecorder) GetFullLink(ctx, shortLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullLink", reflect.TypeOf((*MockRepository)(nil).GetFullLink), ctx, shortLink)
}

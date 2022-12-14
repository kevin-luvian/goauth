// Code generated by MockGen. DO NOT EDIT.
// Source: ./init.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	google "github.com/kevin-luvian/goauth/server/entity/google"
	user "github.com/kevin-luvian/goauth/server/entity/user"
	db "github.com/kevin-luvian/goauth/server/pkg/db"
)

// MockIAuthRepo is a mock of IAuthRepo interface.
type MockIAuthRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthRepoMockRecorder
}

// MockIAuthRepoMockRecorder is the mock recorder for MockIAuthRepo.
type MockIAuthRepoMockRecorder struct {
	mock *MockIAuthRepo
}

// NewMockIAuthRepo creates a new mock instance.
func NewMockIAuthRepo(ctrl *gomock.Controller) *MockIAuthRepo {
	mock := &MockIAuthRepo{ctrl: ctrl}
	mock.recorder = &MockIAuthRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthRepo) EXPECT() *MockIAuthRepoMockRecorder {
	return m.recorder
}

// GetGoogleUserInfo mocks base method.
func (m *MockIAuthRepo) GetGoogleUserInfo(ctx context.Context, code string) (google.GoogleUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoogleUserInfo", ctx, code)
	ret0, _ := ret[0].(google.GoogleUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoogleUserInfo indicates an expected call of GetGoogleUserInfo.
func (mr *MockIAuthRepoMockRecorder) GetGoogleUserInfo(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoogleUserInfo", reflect.TypeOf((*MockIAuthRepo)(nil).GetGoogleUserInfo), ctx, code)
}

// GoogleRedirectURL mocks base method.
func (m *MockIAuthRepo) GoogleRedirectURL(ctx context.Context, state string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoogleRedirectURL", ctx, state)
	ret0, _ := ret[0].(string)
	return ret0
}

// GoogleRedirectURL indicates an expected call of GoogleRedirectURL.
func (mr *MockIAuthRepoMockRecorder) GoogleRedirectURL(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoogleRedirectURL", reflect.TypeOf((*MockIAuthRepo)(nil).GoogleRedirectURL), ctx, state)
}

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepo) Create(ctx context.Context, usr user.User) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, usr)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepoMockRecorder) Create(ctx, usr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepo)(nil).Create), ctx, usr)
}

// Get mocks base method.
func (m *MockIUserRepo) Get(ctx context.Context, param db.GetDBParam) ([]user.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, param)
	ret0, _ := ret[0].([]user.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockIUserRepoMockRecorder) Get(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUserRepo)(nil).Get), ctx, param)
}

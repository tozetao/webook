// Code generated by MockGen. DO NOT EDIT.
// Source: ./interaction.go
//
// Generated by this command:
//
//	mockgen -source=./interaction.go -package=svcmocks -destination=./mocks/interaction.mock.go InteractionService
//

// Package svcmocks is a generated GoMock package.
package svcmocks

import (
	context "context"
	"learn_go/webook/interaction/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInteractionService is a mock of InteractionService interface.
type MockInteractionService struct {
	ctrl     *gomock.Controller
	recorder *MockInteractionServiceMockRecorder
}

// MockInteractionServiceMockRecorder is the mock recorder for MockInteractionService.
type MockInteractionServiceMockRecorder struct {
	mock *MockInteractionService
}

// NewMockInteractionService creates a new mock instance.
func NewMockInteractionService(ctrl *gomock.Controller) *MockInteractionService {
	mock := &MockInteractionService{ctrl: ctrl}
	mock.recorder = &MockInteractionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractionService) EXPECT() *MockInteractionServiceMockRecorder {
	return m.recorder
}

// CancelLike mocks base method.
func (m *MockInteractionService) CancelLike(ctx context.Context, uid int64, biz string, bizID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLike", ctx, uid, biz, bizID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelLike indicates an expected call of CancelLike.
func (mr *MockInteractionServiceMockRecorder) CancelLike(ctx, uid, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLike", reflect.TypeOf((*MockInteractionService)(nil).CancelLike), ctx, uid, biz, bizID)
}

// Collected mocks base method.
func (m *MockInteractionService) Collected(ctx context.Context, uid int64, biz string, bizID int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collected", ctx, uid, biz, bizID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collected indicates an expected call of Collected.
func (mr *MockInteractionServiceMockRecorder) Collected(ctx, uid, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collected", reflect.TypeOf((*MockInteractionService)(nil).Collected), ctx, uid, biz, bizID)
}

// Favorite mocks base method.
func (m *MockInteractionService) Favorite(ctx context.Context, uid, favoriteID int64, biz string, bizID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Favorite", ctx, uid, favoriteID, biz, bizID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Favorite indicates an expected call of Favorite.
func (mr *MockInteractionServiceMockRecorder) Favorite(ctx, uid, favoriteID, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Favorite", reflect.TypeOf((*MockInteractionService)(nil).Favorite), ctx, uid, favoriteID, biz, bizID)
}

// Get mocks base method.
func (m *MockInteractionService) Get(ctx context.Context, uid int64, biz string, bizID int64) (domain.Interaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, uid, biz, bizID)
	ret0, _ := ret[0].(domain.Interaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInteractionServiceMockRecorder) Get(ctx, uid, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInteractionService)(nil).Get), ctx, uid, biz, bizID)
}

// GetByIDs mocks base method.
func (m *MockInteractionService) GetByIDs(ctx context.Context, biz string, bizIDs []int64) (map[int64]domain.Interaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, biz, bizIDs)
	ret0, _ := ret[0].(map[int64]domain.Interaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockInteractionServiceMockRecorder) GetByIDs(ctx, biz, bizIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockInteractionService)(nil).GetByIDs), ctx, biz, bizIDs)
}

// Like mocks base method.
func (m *MockInteractionService) Like(ctx context.Context, uid int64, biz string, bizID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Like", ctx, uid, biz, bizID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Like indicates an expected call of Like.
func (mr *MockInteractionServiceMockRecorder) Like(ctx, uid, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockInteractionService)(nil).Like), ctx, uid, biz, bizID)
}

// Liked mocks base method.
func (m *MockInteractionService) Liked(ctx context.Context, uid int64, biz string, bizID int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Liked", ctx, uid, biz, bizID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Liked indicates an expected call of Liked.
func (mr *MockInteractionServiceMockRecorder) Liked(ctx, uid, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Liked", reflect.TypeOf((*MockInteractionService)(nil).Liked), ctx, uid, biz, bizID)
}

// View mocks base method.
func (m *MockInteractionService) View(ctx context.Context, biz string, bizID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "View", ctx, biz, bizID)
	ret0, _ := ret[0].(error)
	return ret0
}

// View indicates an expected call of View.
func (mr *MockInteractionServiceMockRecorder) View(ctx, biz, bizID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "View", reflect.TypeOf((*MockInteractionService)(nil).View), ctx, biz, bizID)
}

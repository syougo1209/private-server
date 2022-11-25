// Code generated by MockGen. DO NOT EDIT.
// Source: create_text_message.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/syougo1209/b-match-server/domain/model"
)

// MockCreateTextMessage is a mock of CreateTextMessage interface.
type MockCreateTextMessage struct {
	ctrl     *gomock.Controller
	recorder *MockCreateTextMessageMockRecorder
}

// MockCreateTextMessageMockRecorder is the mock recorder for MockCreateTextMessage.
type MockCreateTextMessageMockRecorder struct {
	mock *MockCreateTextMessage
}

// NewMockCreateTextMessage creates a new mock instance.
func NewMockCreateTextMessage(ctrl *gomock.Controller) *MockCreateTextMessage {
	mock := &MockCreateTextMessage{ctrl: ctrl}
	mock.recorder = &MockCreateTextMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateTextMessage) EXPECT() *MockCreateTextMessageMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockCreateTextMessage) Call(ctx context.Context, uid model.UserID, cid model.ConversationID, text string, now time.Time) (*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, uid, cid, text, now)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockCreateTextMessageMockRecorder) Call(ctx, uid, cid, text, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockCreateTextMessage)(nil).Call), ctx, uid, cid, text, now)
}

package darwin

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// mockEnv is a mock of Env interface
type mockEnv struct {
	ctrl     *gomock.Controller
	recorder *mockEnvMockRecorder
}

// mockEnvMockRecorder is the mock recorder for mockEnv
type mockEnvMockRecorder struct {
	mock *mockEnv
}

// NewmockEnv creates a new mock instance
func NewmockEnv(ctrl *gomock.Controller) *mockEnv {
	mock := &mockEnv{ctrl: ctrl}
	mock.recorder = &mockEnvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockEnv) EXPECT() *mockEnvMockRecorder {
	return m.recorder
}

// Fit mocks base method
func (m *mockEnv) Fit(ch Chromosome) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fit", ch)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Fit indicates an expected call of Fit
func (mr *mockEnvMockRecorder) Fit(ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fit", reflect.TypeOf((*mockEnv)(nil).Fit), ch)
}

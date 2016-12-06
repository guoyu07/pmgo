// Automatically generated by MockGen. DO NOT EDIT!
// Source: database.go

package pmgomock

import (
	"github.com/golang/mock/gomock"
	. "github.com/percona/pmgo"
)

// Mock of DatabaseManager interface
type MockDatabaseManager struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseManagerRecorder
}

// Recorder for MockDatabaseManager (not exported)
type _MockDatabaseManagerRecorder struct {
	mock *MockDatabaseManager
}

func NewMockDatabaseManager(ctrl *gomock.Controller) *MockDatabaseManager {
	mock := &MockDatabaseManager{ctrl: ctrl}
	mock.recorder = &_MockDatabaseManagerRecorder{mock}
	return mock
}

func (_m *MockDatabaseManager) EXPECT() *_MockDatabaseManagerRecorder {
	return _m.recorder
}

func (_m *MockDatabaseManager) C(name string) CollectionManager {
	ret := _m.ctrl.Call(_m, "C", name)
	ret0, _ := ret[0].(CollectionManager)
	return ret0
}

func (_mr *_MockDatabaseManagerRecorder) C(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "C", arg0)
}

func (_m *MockDatabaseManager) CollectionNames() ([]string, error) {
	ret := _m.ctrl.Call(_m, "CollectionNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseManagerRecorder) CollectionNames() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CollectionNames")
}

func (_m *MockDatabaseManager) DropDatabase() error {
	ret := _m.ctrl.Call(_m, "DropDatabase")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseManagerRecorder) DropDatabase() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DropDatabase")
}

func (_m *MockDatabaseManager) Login(user string, pass string) error {
	ret := _m.ctrl.Call(_m, "Login", user, pass)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseManagerRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Login", arg0, arg1)
}

func (_m *MockDatabaseManager) Run(cmd interface{}, result interface{}) error {
	ret := _m.ctrl.Call(_m, "Run", cmd, result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseManagerRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}

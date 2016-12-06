// Automatically generated by MockGen. DO NOT EDIT!
// Source: query.go

package pmgomock

import (
	"github.com/golang/mock/gomock"
	. "github.com/percona/pmgo"
	mgo_v2 "gopkg.in/mgo.v2"
)

// Mock of QueryManager interface
type MockQueryManager struct {
	ctrl     *gomock.Controller
	recorder *_MockQueryManagerRecorder
}

// Recorder for MockQueryManager (not exported)
type _MockQueryManagerRecorder struct {
	mock *MockQueryManager
}

func NewMockQueryManager(ctrl *gomock.Controller) *MockQueryManager {
	mock := &MockQueryManager{ctrl: ctrl}
	mock.recorder = &_MockQueryManagerRecorder{mock}
	return mock
}

func (_m *MockQueryManager) EXPECT() *_MockQueryManagerRecorder {
	return _m.recorder
}

func (_m *MockQueryManager) All(result interface{}) error {
	ret := _m.ctrl.Call(_m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) All(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "All", arg0)
}

func (_m *MockQueryManager) Count() (int, error) {
	ret := _m.ctrl.Call(_m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueryManagerRecorder) Count() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Count")
}

func (_m *MockQueryManager) Iter() *mgo_v2.Iter {
	ret := _m.ctrl.Call(_m, "Iter")
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Iter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iter")
}

func (_m *MockQueryManager) Limit(n int) QueryManager {
	ret := _m.ctrl.Call(_m, "Limit", n)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Limit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Limit", arg0)
}

func (_m *MockQueryManager) One(result interface{}) error {
	ret := _m.ctrl.Call(_m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) One(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "One", arg0)
}

func (_m *MockQueryManager) Sort(fields ...string) QueryManager {
	_s := []interface{}{}
	for _, _x := range fields {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Sort", _s...)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Sort(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sort", arg0...)
}

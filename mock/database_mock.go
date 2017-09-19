// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/impactasaurus/server/data (interfaces: Base)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	server "github.com/impactasaurus/server"
	auth "github.com/impactasaurus/server/auth"
)

// MockBase is a mock of Base interface
type MockBase struct {
	ctrl     *gomock.Controller
	recorder *MockBaseMockRecorder
}

// MockBaseMockRecorder is the mock recorder for MockBase
type MockBaseMockRecorder struct {
	mock *MockBase
}

// NewMockBase creates a new mock instance
func NewMockBase(ctrl *gomock.Controller) *MockBase {
	mock := &MockBase{ctrl: ctrl}
	mock.recorder = &MockBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBase) EXPECT() *MockBaseMockRecorder {
	return m.recorder
}

// DeleteCategory mocks base method
func (m *MockBase) DeleteCategory(arg0, arg1 string, arg2 auth.User) error {
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory
func (mr *MockBaseMockRecorder) DeleteCategory(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockBase)(nil).DeleteCategory), arg0, arg1, arg2)
}

// DeleteOutcomeSet mocks base method
func (m *MockBase) DeleteOutcomeSet(arg0 string, arg1 auth.User) error {
	ret := m.ctrl.Call(m, "DeleteOutcomeSet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOutcomeSet indicates an expected call of DeleteOutcomeSet
func (mr *MockBaseMockRecorder) DeleteOutcomeSet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOutcomeSet", reflect.TypeOf((*MockBase)(nil).DeleteOutcomeSet), arg0, arg1)
}

// DeleteQuestion mocks base method
func (m *MockBase) DeleteQuestion(arg0, arg1 string, arg2 auth.User) error {
	ret := m.ctrl.Call(m, "DeleteQuestion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuestion indicates an expected call of DeleteQuestion
func (mr *MockBaseMockRecorder) DeleteQuestion(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestion", reflect.TypeOf((*MockBase)(nil).DeleteQuestion), arg0, arg1, arg2)
}

// EditOutcomeSet mocks base method
func (m *MockBase) EditOutcomeSet(arg0, arg1, arg2 string, arg3 auth.User) (server.OutcomeSet, error) {
	ret := m.ctrl.Call(m, "EditOutcomeSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(server.OutcomeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditOutcomeSet indicates an expected call of EditOutcomeSet
func (mr *MockBaseMockRecorder) EditOutcomeSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditOutcomeSet", reflect.TypeOf((*MockBase)(nil).EditOutcomeSet), arg0, arg1, arg2, arg3)
}

// EditQuestion mocks base method
func (m *MockBase) EditQuestion(arg0, arg1, arg2, arg3 string, arg4 server.QuestionType, arg5 map[string]interface{}, arg6 auth.User) (server.Question, error) {
	ret := m.ctrl.Call(m, "EditQuestion", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(server.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditQuestion indicates an expected call of EditQuestion
func (mr *MockBaseMockRecorder) EditQuestion(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditQuestion", reflect.TypeOf((*MockBase)(nil).EditQuestion), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetCategory mocks base method
func (m *MockBase) GetCategory(arg0, arg1 string, arg2 auth.User) (server.Category, error) {
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1, arg2)
	ret0, _ := ret[0].(server.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory
func (mr *MockBaseMockRecorder) GetCategory(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockBase)(nil).GetCategory), arg0, arg1, arg2)
}

// GetMeeting mocks base method
func (m *MockBase) GetMeeting(arg0 string, arg1 auth.User) (server.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeeting", arg0, arg1)
	ret0, _ := ret[0].(server.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeeting indicates an expected call of GetMeeting
func (mr *MockBaseMockRecorder) GetMeeting(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeeting", reflect.TypeOf((*MockBase)(nil).GetMeeting), arg0, arg1)
}

// GetMeetingsForBeneficiary mocks base method
func (m *MockBase) GetMeetingsForBeneficiary(arg0 string, arg1 auth.User) ([]server.Meeting, error) {
	ret := m.ctrl.Call(m, "GetMeetingsForBeneficiary", arg0, arg1)
	ret0, _ := ret[0].([]server.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingsForBeneficiary indicates an expected call of GetMeetingsForBeneficiary
func (mr *MockBaseMockRecorder) GetMeetingsForBeneficiary(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingsForBeneficiary", reflect.TypeOf((*MockBase)(nil).GetMeetingsForBeneficiary), arg0, arg1)
}

// GetOSMeetingsForBeneficiary mocks base method
func (m *MockBase) GetOSMeetingsForBeneficiary(arg0, arg1 string, arg2 auth.User) ([]server.Meeting, error) {
	ret := m.ctrl.Call(m, "GetOSMeetingsForBeneficiary", arg0, arg1, arg2)
	ret0, _ := ret[0].([]server.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOSMeetingsForBeneficiary indicates an expected call of GetOSMeetingsForBeneficiary
func (mr *MockBaseMockRecorder) GetOSMeetingsForBeneficiary(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOSMeetingsForBeneficiary", reflect.TypeOf((*MockBase)(nil).GetOSMeetingsForBeneficiary), arg0, arg1, arg2)
}

// GetOSMeetingsInTimeRange mocks base method
func (m *MockBase) GetOSMeetingsInTimeRange(arg0, arg1 time.Time, arg2 string, arg3 auth.User) ([]server.Meeting, error) {
	ret := m.ctrl.Call(m, "GetOSMeetingsInTimeRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]server.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOSMeetingsInTimeRange indicates an expected call of GetOSMeetingsInTimeRange
func (mr *MockBaseMockRecorder) GetOSMeetingsInTimeRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOSMeetingsInTimeRange", reflect.TypeOf((*MockBase)(nil).GetOSMeetingsInTimeRange), arg0, arg1, arg2, arg3)
}

// GetOrganisation mocks base method
func (m *MockBase) GetOrganisation(arg0 string, arg1 auth.User) (server.Organisation, error) {
	ret := m.ctrl.Call(m, "GetOrganisation", arg0, arg1)
	ret0, _ := ret[0].(server.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganisation indicates an expected call of GetOrganisation
func (mr *MockBaseMockRecorder) GetOrganisation(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganisation", reflect.TypeOf((*MockBase)(nil).GetOrganisation), arg0, arg1)
}

// GetOutcomeSet mocks base method
func (m *MockBase) GetOutcomeSet(arg0 string, arg1 auth.User) (server.OutcomeSet, error) {
	ret := m.ctrl.Call(m, "GetOutcomeSet", arg0, arg1)
	ret0, _ := ret[0].(server.OutcomeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutcomeSet indicates an expected call of GetOutcomeSet
func (mr *MockBaseMockRecorder) GetOutcomeSet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutcomeSet", reflect.TypeOf((*MockBase)(nil).GetOutcomeSet), arg0, arg1)
}

// GetOutcomeSets mocks base method
func (m *MockBase) GetOutcomeSets(arg0 auth.User) ([]server.OutcomeSet, error) {
	ret := m.ctrl.Call(m, "GetOutcomeSets", arg0)
	ret0, _ := ret[0].([]server.OutcomeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutcomeSets indicates an expected call of GetOutcomeSets
func (mr *MockBaseMockRecorder) GetOutcomeSets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutcomeSets", reflect.TypeOf((*MockBase)(nil).GetOutcomeSets), arg0)
}

// GetQuestion mocks base method
func (m *MockBase) GetQuestion(arg0, arg1 string, arg2 auth.User) (server.Question, error) {
	ret := m.ctrl.Call(m, "GetQuestion", arg0, arg1, arg2)
	ret0, _ := ret[0].(server.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestion indicates an expected call of GetQuestion
func (mr *MockBaseMockRecorder) GetQuestion(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestion", reflect.TypeOf((*MockBase)(nil).GetQuestion), arg0, arg1, arg2)
}

// MoveQuestion mocks base method
func (m *MockBase) MoveQuestion(arg0, arg1 string, arg2 uint, arg3 auth.User) error {
	ret := m.ctrl.Call(m, "MoveQuestion", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveQuestion indicates an expected call of MoveQuestion
func (mr *MockBaseMockRecorder) MoveQuestion(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveQuestion", reflect.TypeOf((*MockBase)(nil).MoveQuestion), arg0, arg1, arg2, arg3)
}

// NewAnswer mocks base method
func (m *MockBase) NewAnswer(arg0 string, arg1 server.Answer, arg2 auth.User) (server.Meeting, error) {
	ret := m.ctrl.Call(m, "NewAnswer", arg0, arg1, arg2)
	ret0, _ := ret[0].(server.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAnswer indicates an expected call of NewAnswer
func (mr *MockBaseMockRecorder) NewAnswer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAnswer", reflect.TypeOf((*MockBase)(nil).NewAnswer), arg0, arg1, arg2)
}

// NewCategory mocks base method
func (m *MockBase) NewCategory(arg0, arg1, arg2 string, arg3 server.Aggregation, arg4 auth.User) (server.Category, error) {
	ret := m.ctrl.Call(m, "NewCategory", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(server.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCategory indicates an expected call of NewCategory
func (mr *MockBaseMockRecorder) NewCategory(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCategory", reflect.TypeOf((*MockBase)(nil).NewCategory), arg0, arg1, arg2, arg3, arg4)
}

// NewMeeting mocks base method
func (m *MockBase) NewMeeting(arg0, arg1 string, arg2 time.Time, arg3 auth.User) (server.Meeting, error) {
	ret := m.ctrl.Call(m, "NewMeeting", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(server.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMeeting indicates an expected call of NewMeeting
func (mr *MockBaseMockRecorder) NewMeeting(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMeeting", reflect.TypeOf((*MockBase)(nil).NewMeeting), arg0, arg1, arg2, arg3)
}

// NewOutcomeSet mocks base method
func (m *MockBase) NewOutcomeSet(arg0, arg1 string, arg2 auth.User) (server.OutcomeSet, error) {
	ret := m.ctrl.Call(m, "NewOutcomeSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(server.OutcomeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewOutcomeSet indicates an expected call of NewOutcomeSet
func (mr *MockBaseMockRecorder) NewOutcomeSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewOutcomeSet", reflect.TypeOf((*MockBase)(nil).NewOutcomeSet), arg0, arg1, arg2)
}

// NewQuestion mocks base method
func (m *MockBase) NewQuestion(arg0, arg1, arg2 string, arg3 server.QuestionType, arg4 map[string]interface{}, arg5 auth.User) (server.Question, error) {
	ret := m.ctrl.Call(m, "NewQuestion", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(server.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewQuestion indicates an expected call of NewQuestion
func (mr *MockBaseMockRecorder) NewQuestion(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewQuestion", reflect.TypeOf((*MockBase)(nil).NewQuestion), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveCategory mocks base method
func (m *MockBase) RemoveCategory(arg0, arg1 string, arg2 auth.User) (server.Question, error) {
	ret := m.ctrl.Call(m, "RemoveCategory", arg0, arg1, arg2)
	ret0, _ := ret[0].(server.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveCategory indicates an expected call of RemoveCategory
func (mr *MockBaseMockRecorder) RemoveCategory(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCategory", reflect.TypeOf((*MockBase)(nil).RemoveCategory), arg0, arg1, arg2)
}

// SetCategory mocks base method
func (m *MockBase) SetCategory(arg0, arg1, arg2 string, arg3 auth.User) (server.Question, error) {
	ret := m.ctrl.Call(m, "SetCategory", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(server.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCategory indicates an expected call of SetCategory
func (mr *MockBaseMockRecorder) SetCategory(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCategory", reflect.TypeOf((*MockBase)(nil).SetCategory), arg0, arg1, arg2, arg3)
}

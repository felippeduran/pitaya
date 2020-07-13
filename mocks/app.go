// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/topfreegames/pitaya (interfaces: Pitaya)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	cluster "github.com/topfreegames/pitaya/cluster"
	component "github.com/topfreegames/pitaya/component"
	config "github.com/topfreegames/pitaya/config"
	interfaces "github.com/topfreegames/pitaya/interfaces"
	metrics "github.com/topfreegames/pitaya/metrics"
	router "github.com/topfreegames/pitaya/router"
	session "github.com/topfreegames/pitaya/session"
	worker "github.com/topfreegames/pitaya/worker"
	reflect "reflect"
	time "time"
)

// MockPitaya is a mock of Pitaya interface
type MockPitaya struct {
	ctrl     *gomock.Controller
	recorder *MockPitayaMockRecorder
}

// MockPitayaMockRecorder is the mock recorder for MockPitaya
type MockPitayaMockRecorder struct {
	mock *MockPitaya
}

// NewMockPitaya creates a new mock instance
func NewMockPitaya(ctrl *gomock.Controller) *MockPitaya {
	mock := &MockPitaya{ctrl: ctrl}
	mock.recorder = &MockPitayaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPitaya) EXPECT() *MockPitayaMockRecorder {
	return m.recorder
}

// AddRoute mocks base method
func (m *MockPitaya) AddRoute(arg0 string, arg1 router.RoutingFunc) error {
	ret := m.ctrl.Call(m, "AddRoute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoute indicates an expected call of AddRoute
func (mr *MockPitayaMockRecorder) AddRoute(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockPitaya)(nil).AddRoute), arg0, arg1)
}

// Documentation mocks base method
func (m *MockPitaya) Documentation(arg0 bool) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "Documentation", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Documentation indicates an expected call of Documentation
func (mr *MockPitayaMockRecorder) Documentation(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Documentation", reflect.TypeOf((*MockPitaya)(nil).Documentation), arg0)
}

// GetConfig mocks base method
func (m *MockPitaya) GetConfig() *config.Config {
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockPitayaMockRecorder) GetConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockPitaya)(nil).GetConfig))
}

// GetDieChan mocks base method
func (m *MockPitaya) GetDieChan() chan bool {
	ret := m.ctrl.Call(m, "GetDieChan")
	ret0, _ := ret[0].(chan bool)
	return ret0
}

// GetDieChan indicates an expected call of GetDieChan
func (mr *MockPitayaMockRecorder) GetDieChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDieChan", reflect.TypeOf((*MockPitaya)(nil).GetDieChan))
}

// GetMetricsReporters mocks base method
func (m *MockPitaya) GetMetricsReporters() []metrics.Reporter {
	ret := m.ctrl.Call(m, "GetMetricsReporters")
	ret0, _ := ret[0].([]metrics.Reporter)
	return ret0
}

// GetMetricsReporters indicates an expected call of GetMetricsReporters
func (mr *MockPitayaMockRecorder) GetMetricsReporters() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsReporters", reflect.TypeOf((*MockPitaya)(nil).GetMetricsReporters))
}

// GetModule mocks base method
func (m *MockPitaya) GetModule(arg0 string) (interfaces.Module, error) {
	ret := m.ctrl.Call(m, "GetModule", arg0)
	ret0, _ := ret[0].(interfaces.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModule indicates an expected call of GetModule
func (mr *MockPitayaMockRecorder) GetModule(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModule", reflect.TypeOf((*MockPitaya)(nil).GetModule), arg0)
}

// GetServer mocks base method
func (m *MockPitaya) GetServer() *cluster.Server {
	ret := m.ctrl.Call(m, "GetServer")
	ret0, _ := ret[0].(*cluster.Server)
	return ret0
}

// GetServer indicates an expected call of GetServer
func (mr *MockPitayaMockRecorder) GetServer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockPitaya)(nil).GetServer))
}

// GetServerByID mocks base method
func (m *MockPitaya) GetServerByID(arg0 string) (*cluster.Server, error) {
	ret := m.ctrl.Call(m, "GetServerByID", arg0)
	ret0, _ := ret[0].(*cluster.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerByID indicates an expected call of GetServerByID
func (mr *MockPitayaMockRecorder) GetServerByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerByID", reflect.TypeOf((*MockPitaya)(nil).GetServerByID), arg0)
}

// GetServerID mocks base method
func (m *MockPitaya) GetServerID() string {
	ret := m.ctrl.Call(m, "GetServerID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServerID indicates an expected call of GetServerID
func (mr *MockPitayaMockRecorder) GetServerID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerID", reflect.TypeOf((*MockPitaya)(nil).GetServerID))
}

// GetServers mocks base method
func (m *MockPitaya) GetServers() []*cluster.Server {
	ret := m.ctrl.Call(m, "GetServers")
	ret0, _ := ret[0].([]*cluster.Server)
	return ret0
}

// GetServers indicates an expected call of GetServers
func (mr *MockPitayaMockRecorder) GetServers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServers", reflect.TypeOf((*MockPitaya)(nil).GetServers))
}

// GetServersByType mocks base method
func (m *MockPitaya) GetServersByType(arg0 string) (map[string]*cluster.Server, error) {
	ret := m.ctrl.Call(m, "GetServersByType", arg0)
	ret0, _ := ret[0].(map[string]*cluster.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServersByType indicates an expected call of GetServersByType
func (mr *MockPitayaMockRecorder) GetServersByType(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServersByType", reflect.TypeOf((*MockPitaya)(nil).GetServersByType), arg0)
}

// GetSessionFromCtx mocks base method
func (m *MockPitaya) GetSessionFromCtx(arg0 context.Context) session.Session {
	ret := m.ctrl.Call(m, "GetSessionFromCtx", arg0)
	ret0, _ := ret[0].(session.Session)
	return ret0
}

// GetSessionFromCtx indicates an expected call of GetSessionFromCtx
func (mr *MockPitayaMockRecorder) GetSessionFromCtx(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionFromCtx", reflect.TypeOf((*MockPitaya)(nil).GetSessionFromCtx), arg0)
}

// GroupAddMember mocks base method
func (m *MockPitaya) GroupAddMember(arg0 context.Context, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "GroupAddMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupAddMember indicates an expected call of GroupAddMember
func (mr *MockPitayaMockRecorder) GroupAddMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupAddMember", reflect.TypeOf((*MockPitaya)(nil).GroupAddMember), arg0, arg1, arg2)
}

// GroupBroadcast mocks base method
func (m *MockPitaya) GroupBroadcast(arg0 context.Context, arg1, arg2, arg3 string, arg4 interface{}) error {
	ret := m.ctrl.Call(m, "GroupBroadcast", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupBroadcast indicates an expected call of GroupBroadcast
func (mr *MockPitayaMockRecorder) GroupBroadcast(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupBroadcast", reflect.TypeOf((*MockPitaya)(nil).GroupBroadcast), arg0, arg1, arg2, arg3, arg4)
}

// GroupContainsMember mocks base method
func (m *MockPitaya) GroupContainsMember(arg0 context.Context, arg1, arg2 string) (bool, error) {
	ret := m.ctrl.Call(m, "GroupContainsMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupContainsMember indicates an expected call of GroupContainsMember
func (mr *MockPitayaMockRecorder) GroupContainsMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupContainsMember", reflect.TypeOf((*MockPitaya)(nil).GroupContainsMember), arg0, arg1, arg2)
}

// GroupCountMembers mocks base method
func (m *MockPitaya) GroupCountMembers(arg0 context.Context, arg1 string) (int, error) {
	ret := m.ctrl.Call(m, "GroupCountMembers", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupCountMembers indicates an expected call of GroupCountMembers
func (mr *MockPitayaMockRecorder) GroupCountMembers(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupCountMembers", reflect.TypeOf((*MockPitaya)(nil).GroupCountMembers), arg0, arg1)
}

// GroupCreate mocks base method
func (m *MockPitaya) GroupCreate(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "GroupCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupCreate indicates an expected call of GroupCreate
func (mr *MockPitayaMockRecorder) GroupCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupCreate", reflect.TypeOf((*MockPitaya)(nil).GroupCreate), arg0, arg1)
}

// GroupCreateWithTTL mocks base method
func (m *MockPitaya) GroupCreateWithTTL(arg0 context.Context, arg1 string, arg2 time.Duration) error {
	ret := m.ctrl.Call(m, "GroupCreateWithTTL", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupCreateWithTTL indicates an expected call of GroupCreateWithTTL
func (mr *MockPitayaMockRecorder) GroupCreateWithTTL(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupCreateWithTTL", reflect.TypeOf((*MockPitaya)(nil).GroupCreateWithTTL), arg0, arg1, arg2)
}

// GroupDelete mocks base method
func (m *MockPitaya) GroupDelete(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "GroupDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupDelete indicates an expected call of GroupDelete
func (mr *MockPitayaMockRecorder) GroupDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupDelete", reflect.TypeOf((*MockPitaya)(nil).GroupDelete), arg0, arg1)
}

// GroupMembers mocks base method
func (m *MockPitaya) GroupMembers(arg0 context.Context, arg1 string) ([]string, error) {
	ret := m.ctrl.Call(m, "GroupMembers", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupMembers indicates an expected call of GroupMembers
func (mr *MockPitayaMockRecorder) GroupMembers(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupMembers", reflect.TypeOf((*MockPitaya)(nil).GroupMembers), arg0, arg1)
}

// GroupRemoveAll mocks base method
func (m *MockPitaya) GroupRemoveAll(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "GroupRemoveAll", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupRemoveAll indicates an expected call of GroupRemoveAll
func (mr *MockPitayaMockRecorder) GroupRemoveAll(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupRemoveAll", reflect.TypeOf((*MockPitaya)(nil).GroupRemoveAll), arg0, arg1)
}

// GroupRemoveMember mocks base method
func (m *MockPitaya) GroupRemoveMember(arg0 context.Context, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "GroupRemoveMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupRemoveMember indicates an expected call of GroupRemoveMember
func (mr *MockPitayaMockRecorder) GroupRemoveMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupRemoveMember", reflect.TypeOf((*MockPitaya)(nil).GroupRemoveMember), arg0, arg1, arg2)
}

// GroupRenewTTL mocks base method
func (m *MockPitaya) GroupRenewTTL(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "GroupRenewTTL", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GroupRenewTTL indicates an expected call of GroupRenewTTL
func (mr *MockPitayaMockRecorder) GroupRenewTTL(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupRenewTTL", reflect.TypeOf((*MockPitaya)(nil).GroupRenewTTL), arg0, arg1)
}

// RPC mocks base method
func (m *MockPitaya) RPC(arg0 context.Context, arg1 string, arg2, arg3 proto.Message) error {
	ret := m.ctrl.Call(m, "RPC", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RPC indicates an expected call of RPC
func (mr *MockPitayaMockRecorder) RPC(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPC", reflect.TypeOf((*MockPitaya)(nil).RPC), arg0, arg1, arg2, arg3)
}

// RPCTo mocks base method
func (m *MockPitaya) RPCTo(arg0 context.Context, arg1, arg2 string, arg3, arg4 proto.Message) error {
	ret := m.ctrl.Call(m, "RPCTo", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// RPCTo indicates an expected call of RPCTo
func (mr *MockPitayaMockRecorder) RPCTo(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCTo", reflect.TypeOf((*MockPitaya)(nil).RPCTo), arg0, arg1, arg2, arg3, arg4)
}

// Register mocks base method
func (m *MockPitaya) Register(arg0 component.Component, arg1 ...component.Option) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Register", varargs...)
}

// Register indicates an expected call of Register
func (mr *MockPitayaMockRecorder) Register(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPitaya)(nil).Register), varargs...)
}

// RegisterModule mocks base method
func (m *MockPitaya) RegisterModule(arg0 interfaces.Module, arg1 string) error {
	ret := m.ctrl.Call(m, "RegisterModule", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterModule indicates an expected call of RegisterModule
func (mr *MockPitayaMockRecorder) RegisterModule(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterModule", reflect.TypeOf((*MockPitaya)(nil).RegisterModule), arg0, arg1)
}

// RegisterModuleAfter mocks base method
func (m *MockPitaya) RegisterModuleAfter(arg0 interfaces.Module, arg1 string) error {
	ret := m.ctrl.Call(m, "RegisterModuleAfter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterModuleAfter indicates an expected call of RegisterModuleAfter
func (mr *MockPitayaMockRecorder) RegisterModuleAfter(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterModuleAfter", reflect.TypeOf((*MockPitaya)(nil).RegisterModuleAfter), arg0, arg1)
}

// RegisterModuleBefore mocks base method
func (m *MockPitaya) RegisterModuleBefore(arg0 interfaces.Module, arg1 string) error {
	ret := m.ctrl.Call(m, "RegisterModuleBefore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterModuleBefore indicates an expected call of RegisterModuleBefore
func (mr *MockPitayaMockRecorder) RegisterModuleBefore(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterModuleBefore", reflect.TypeOf((*MockPitaya)(nil).RegisterModuleBefore), arg0, arg1)
}

// RegisterRPCJob mocks base method
func (m *MockPitaya) RegisterRPCJob(arg0 worker.RPCJob) error {
	ret := m.ctrl.Call(m, "RegisterRPCJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterRPCJob indicates an expected call of RegisterRPCJob
func (mr *MockPitayaMockRecorder) RegisterRPCJob(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRPCJob", reflect.TypeOf((*MockPitaya)(nil).RegisterRPCJob), arg0)
}

// RegisterRemote mocks base method
func (m *MockPitaya) RegisterRemote(arg0 component.Component, arg1 ...component.Option) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RegisterRemote", varargs...)
}

// RegisterRemote indicates an expected call of RegisterRemote
func (mr *MockPitayaMockRecorder) RegisterRemote(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRemote", reflect.TypeOf((*MockPitaya)(nil).RegisterRemote), varargs...)
}

// ReliableRPC mocks base method
func (m *MockPitaya) ReliableRPC(arg0 string, arg1 map[string]interface{}, arg2, arg3 proto.Message) (string, error) {
	ret := m.ctrl.Call(m, "ReliableRPC", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReliableRPC indicates an expected call of ReliableRPC
func (mr *MockPitayaMockRecorder) ReliableRPC(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReliableRPC", reflect.TypeOf((*MockPitaya)(nil).ReliableRPC), arg0, arg1, arg2, arg3)
}

// ReliableRPCWithOptions mocks base method
func (m *MockPitaya) ReliableRPCWithOptions(arg0 string, arg1 map[string]interface{}, arg2, arg3 proto.Message, arg4 *worker.EnqueueOpts) (string, error) {
	ret := m.ctrl.Call(m, "ReliableRPCWithOptions", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReliableRPCWithOptions indicates an expected call of ReliableRPCWithOptions
func (mr *MockPitayaMockRecorder) ReliableRPCWithOptions(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReliableRPCWithOptions", reflect.TypeOf((*MockPitaya)(nil).ReliableRPCWithOptions), arg0, arg1, arg2, arg3, arg4)
}

// SendKickToUsers mocks base method
func (m *MockPitaya) SendKickToUsers(arg0 []string, arg1 string) ([]string, error) {
	ret := m.ctrl.Call(m, "SendKickToUsers", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendKickToUsers indicates an expected call of SendKickToUsers
func (mr *MockPitayaMockRecorder) SendKickToUsers(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendKickToUsers", reflect.TypeOf((*MockPitaya)(nil).SendKickToUsers), arg0, arg1)
}

// SendPushToUsers mocks base method
func (m *MockPitaya) SendPushToUsers(arg0 string, arg1 interface{}, arg2 []string, arg3 string) ([]string, error) {
	ret := m.ctrl.Call(m, "SendPushToUsers", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendPushToUsers indicates an expected call of SendPushToUsers
func (mr *MockPitayaMockRecorder) SendPushToUsers(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPushToUsers", reflect.TypeOf((*MockPitaya)(nil).SendPushToUsers), arg0, arg1, arg2, arg3)
}

// SetDebug mocks base method
func (m *MockPitaya) SetDebug(arg0 bool) {
	m.ctrl.Call(m, "SetDebug", arg0)
}

// SetDebug indicates an expected call of SetDebug
func (mr *MockPitayaMockRecorder) SetDebug(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDebug", reflect.TypeOf((*MockPitaya)(nil).SetDebug), arg0)
}

// SetDictionary mocks base method
func (m *MockPitaya) SetDictionary(arg0 map[string]uint16) error {
	ret := m.ctrl.Call(m, "SetDictionary", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDictionary indicates an expected call of SetDictionary
func (mr *MockPitayaMockRecorder) SetDictionary(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDictionary", reflect.TypeOf((*MockPitaya)(nil).SetDictionary), arg0)
}

// SetHeartbeatTime mocks base method
func (m *MockPitaya) SetHeartbeatTime(arg0 time.Duration) {
	m.ctrl.Call(m, "SetHeartbeatTime", arg0)
}

// SetHeartbeatTime indicates an expected call of SetHeartbeatTime
func (mr *MockPitayaMockRecorder) SetHeartbeatTime(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeartbeatTime", reflect.TypeOf((*MockPitaya)(nil).SetHeartbeatTime), arg0)
}

// Shutdown mocks base method
func (m *MockPitaya) Shutdown() {
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockPitayaMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockPitaya)(nil).Shutdown))
}

// Start mocks base method
func (m *MockPitaya) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockPitayaMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPitaya)(nil).Start))
}

// StartWorker mocks base method
func (m *MockPitaya) StartWorker() {
	m.ctrl.Call(m, "StartWorker")
}

// StartWorker indicates an expected call of StartWorker
func (mr *MockPitayaMockRecorder) StartWorker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorker", reflect.TypeOf((*MockPitaya)(nil).StartWorker))
}

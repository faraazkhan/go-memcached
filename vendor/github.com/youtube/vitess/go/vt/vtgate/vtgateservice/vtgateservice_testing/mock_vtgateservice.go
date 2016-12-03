// Automatically generated by MockGen. DO NOT EDIT!
// Source: interface.go

package vtgateservice_testing

import (
	gomock "github.com/golang/mock/gomock"
	sqltypes "github.com/youtube/vitess/go/sqltypes"
	query "github.com/youtube/vitess/go/vt/proto/query"
	topodata "github.com/youtube/vitess/go/vt/proto/topodata"
	vtgate "github.com/youtube/vitess/go/vt/proto/vtgate"
	context "golang.org/x/net/context"
)

// Mock of VTGateService interface
type MockVTGateService struct {
	ctrl     *gomock.Controller
	recorder *_MockVTGateServiceRecorder
}

// Recorder for MockVTGateService (not exported)
type _MockVTGateServiceRecorder struct {
	mock *MockVTGateService
}

func NewMockVTGateService(ctrl *gomock.Controller) *MockVTGateService {
	mock := &MockVTGateService{ctrl: ctrl}
	mock.recorder = &_MockVTGateServiceRecorder{mock}
	return mock
}

func (_m *MockVTGateService) EXPECT() *_MockVTGateServiceRecorder {
	return _m.recorder
}

func (_m *MockVTGateService) Execute(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, tabletType topodata.TabletType, session *vtgate.Session, notInTransaction bool, options *query.ExecuteOptions) (*sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "Execute", ctx, sql, bindVariables, keyspace, tabletType, session, notInTransaction, options)
	ret0, _ := ret[0].(*sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) Execute(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Execute", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockVTGateService) ExecuteShards(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, shards []string, tabletType topodata.TabletType, session *vtgate.Session, notInTransaction bool, options *query.ExecuteOptions) (*sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "ExecuteShards", ctx, sql, bindVariables, keyspace, shards, tabletType, session, notInTransaction, options)
	ret0, _ := ret[0].(*sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) ExecuteShards(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecuteShards", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

func (_m *MockVTGateService) ExecuteKeyspaceIds(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, keyspaceIds [][]byte, tabletType topodata.TabletType, session *vtgate.Session, notInTransaction bool, options *query.ExecuteOptions) (*sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "ExecuteKeyspaceIds", ctx, sql, bindVariables, keyspace, keyspaceIds, tabletType, session, notInTransaction, options)
	ret0, _ := ret[0].(*sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) ExecuteKeyspaceIds(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecuteKeyspaceIds", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

func (_m *MockVTGateService) ExecuteKeyRanges(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, keyRanges []*topodata.KeyRange, tabletType topodata.TabletType, session *vtgate.Session, notInTransaction bool, options *query.ExecuteOptions) (*sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "ExecuteKeyRanges", ctx, sql, bindVariables, keyspace, keyRanges, tabletType, session, notInTransaction, options)
	ret0, _ := ret[0].(*sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) ExecuteKeyRanges(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecuteKeyRanges", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

func (_m *MockVTGateService) ExecuteEntityIds(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, entityColumnName string, entityKeyspaceIDs []*vtgate.ExecuteEntityIdsRequest_EntityId, tabletType topodata.TabletType, session *vtgate.Session, notInTransaction bool, options *query.ExecuteOptions) (*sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "ExecuteEntityIds", ctx, sql, bindVariables, keyspace, entityColumnName, entityKeyspaceIDs, tabletType, session, notInTransaction, options)
	ret0, _ := ret[0].(*sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) ExecuteEntityIds(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecuteEntityIds", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

func (_m *MockVTGateService) ExecuteBatchShards(ctx context.Context, queries []*vtgate.BoundShardQuery, tabletType topodata.TabletType, asTransaction bool, session *vtgate.Session, options *query.ExecuteOptions) ([]sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "ExecuteBatchShards", ctx, queries, tabletType, asTransaction, session, options)
	ret0, _ := ret[0].([]sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) ExecuteBatchShards(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecuteBatchShards", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockVTGateService) ExecuteBatchKeyspaceIds(ctx context.Context, queries []*vtgate.BoundKeyspaceIdQuery, tabletType topodata.TabletType, asTransaction bool, session *vtgate.Session, options *query.ExecuteOptions) ([]sqltypes.Result, error) {
	ret := _m.ctrl.Call(_m, "ExecuteBatchKeyspaceIds", ctx, queries, tabletType, asTransaction, session, options)
	ret0, _ := ret[0].([]sqltypes.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) ExecuteBatchKeyspaceIds(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecuteBatchKeyspaceIds", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockVTGateService) StreamExecute(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, tabletType topodata.TabletType, options *query.ExecuteOptions, sendReply func(*sqltypes.Result) error) error {
	ret := _m.ctrl.Call(_m, "StreamExecute", ctx, sql, bindVariables, keyspace, tabletType, options, sendReply)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) StreamExecute(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamExecute", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockVTGateService) StreamExecuteShards(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, shards []string, tabletType topodata.TabletType, options *query.ExecuteOptions, sendReply func(*sqltypes.Result) error) error {
	ret := _m.ctrl.Call(_m, "StreamExecuteShards", ctx, sql, bindVariables, keyspace, shards, tabletType, options, sendReply)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) StreamExecuteShards(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamExecuteShards", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockVTGateService) StreamExecuteKeyspaceIds(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, keyspaceIds [][]byte, tabletType topodata.TabletType, options *query.ExecuteOptions, sendReply func(*sqltypes.Result) error) error {
	ret := _m.ctrl.Call(_m, "StreamExecuteKeyspaceIds", ctx, sql, bindVariables, keyspace, keyspaceIds, tabletType, options, sendReply)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) StreamExecuteKeyspaceIds(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamExecuteKeyspaceIds", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockVTGateService) StreamExecuteKeyRanges(ctx context.Context, sql string, bindVariables map[string]interface{}, keyspace string, keyRanges []*topodata.KeyRange, tabletType topodata.TabletType, options *query.ExecuteOptions, sendReply func(*sqltypes.Result) error) error {
	ret := _m.ctrl.Call(_m, "StreamExecuteKeyRanges", ctx, sql, bindVariables, keyspace, keyRanges, tabletType, options, sendReply)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) StreamExecuteKeyRanges(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamExecuteKeyRanges", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockVTGateService) Begin(ctx context.Context) (*vtgate.Session, error) {
	ret := _m.ctrl.Call(_m, "Begin", ctx)
	ret0, _ := ret[0].(*vtgate.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) Begin(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Begin", arg0)
}

func (_m *MockVTGateService) Commit(ctx context.Context, session *vtgate.Session) error {
	ret := _m.ctrl.Call(_m, "Commit", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) Commit(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Commit", arg0, arg1)
}

func (_m *MockVTGateService) Rollback(ctx context.Context, session *vtgate.Session) error {
	ret := _m.ctrl.Call(_m, "Rollback", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) Rollback(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rollback", arg0, arg1)
}

func (_m *MockVTGateService) SplitQuery(ctx context.Context, keyspace string, sql string, bindVariables map[string]interface{}, splitColumns []string, splitCount int64, numRowsPerQueryPart int64, algorithm query.SplitQueryRequest_Algorithm) ([]*vtgate.SplitQueryResponse_Part, error) {
	ret := _m.ctrl.Call(_m, "SplitQuery", ctx, keyspace, sql, bindVariables, splitColumns, splitCount, numRowsPerQueryPart, algorithm)
	ret0, _ := ret[0].([]*vtgate.SplitQueryResponse_Part)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) SplitQuery(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SplitQuery", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockVTGateService) GetSrvKeyspace(ctx context.Context, keyspace string) (*topodata.SrvKeyspace, error) {
	ret := _m.ctrl.Call(_m, "GetSrvKeyspace", ctx, keyspace)
	ret0, _ := ret[0].(*topodata.SrvKeyspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVTGateServiceRecorder) GetSrvKeyspace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSrvKeyspace", arg0, arg1)
}

func (_m *MockVTGateService) UpdateStream(ctx context.Context, keyspace string, shard string, keyRange *topodata.KeyRange, tabletType topodata.TabletType, timestamp int64, event *query.EventToken, sendReply func(*query.StreamEvent, int64) error) error {
	ret := _m.ctrl.Call(_m, "UpdateStream", ctx, keyspace, shard, keyRange, tabletType, timestamp, event, sendReply)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVTGateServiceRecorder) UpdateStream(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateStream", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockVTGateService) HandlePanic(err *error) {
	_m.ctrl.Call(_m, "HandlePanic", err)
}

func (_mr *_MockVTGateServiceRecorder) HandlePanic(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandlePanic", arg0)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/backplane-api/pkg/client (interfaces: ClientWithResponsesInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	Openapi "github.com/openshift/backplane-api/pkg/client"
)

// MockClientWithResponsesInterface is a mock of ClientWithResponsesInterface interface.
type MockClientWithResponsesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientWithResponsesInterfaceMockRecorder
}

// MockClientWithResponsesInterfaceMockRecorder is the mock recorder for MockClientWithResponsesInterface.
type MockClientWithResponsesInterfaceMockRecorder struct {
	mock *MockClientWithResponsesInterface
}

// NewMockClientWithResponsesInterface creates a new mock instance.
func NewMockClientWithResponsesInterface(ctrl *gomock.Controller) *MockClientWithResponsesInterface {
	mock := &MockClientWithResponsesInterface{ctrl: ctrl}
	mock.recorder = &MockClientWithResponsesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientWithResponsesInterface) EXPECT() *MockClientWithResponsesInterfaceMockRecorder {
	return m.recorder
}

// CreateJobWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateJobWithBodyWithResponse(arg0 context.Context, arg1, arg2 string, arg3 io.Reader, arg4 ...Openapi.RequestEditorFn) (*Openapi.CreateJobResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateJobWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.CreateJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobWithBodyWithResponse indicates an expected call of CreateJobWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateJobWithBodyWithResponse(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateJobWithBodyWithResponse), varargs...)
}

// CreateJobWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateJobWithResponse(arg0 context.Context, arg1 string, arg2 Openapi.CreateJobJSONRequestBody, arg3 ...Openapi.RequestEditorFn) (*Openapi.CreateJobResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateJobWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.CreateJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobWithResponse indicates an expected call of CreateJobWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateJobWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateJobWithResponse), varargs...)
}

// CreateTestScriptRunWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateTestScriptRunWithBodyWithResponse(arg0 context.Context, arg1, arg2 string, arg3 io.Reader, arg4 ...Openapi.RequestEditorFn) (*Openapi.CreateTestScriptRunResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTestScriptRunWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.CreateTestScriptRunResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestScriptRunWithBodyWithResponse indicates an expected call of CreateTestScriptRunWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateTestScriptRunWithBodyWithResponse(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestScriptRunWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateTestScriptRunWithBodyWithResponse), varargs...)
}

// CreateTestScriptRunWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateTestScriptRunWithResponse(arg0 context.Context, arg1 string, arg2 Openapi.CreateTestScriptRunJSONRequestBody, arg3 ...Openapi.RequestEditorFn) (*Openapi.CreateTestScriptRunResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTestScriptRunWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.CreateTestScriptRunResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestScriptRunWithResponse indicates an expected call of CreateTestScriptRunWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateTestScriptRunWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestScriptRunWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateTestScriptRunWithResponse), varargs...)
}

// DeleteBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) DeleteBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.DeleteBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.DeleteBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackplaneClusterClusterIdWithResponse indicates an expected call of DeleteBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) DeleteBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).DeleteBackplaneClusterClusterIdWithResponse), varargs...)
}

// DeleteJobWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) DeleteJobWithResponse(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*Openapi.DeleteJobResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteJobWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.DeleteJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteJobWithResponse indicates an expected call of DeleteJobWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) DeleteJobWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJobWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).DeleteJobWithResponse), varargs...)
}

// GetAllJobsWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetAllJobsWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.GetAllJobsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllJobsWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetAllJobsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobsWithResponse indicates an expected call of GetAllJobsWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetAllJobsWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobsWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetAllJobsWithResponse), varargs...)
}

// GetBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.GetBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneClusterClusterIdWithResponse indicates an expected call of GetBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetBackplaneClusterClusterIdWithResponse), varargs...)
}

// GetCloudConsoleWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetCloudConsoleWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.GetCloudConsoleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCloudConsoleWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetCloudConsoleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudConsoleWithResponse indicates an expected call of GetCloudConsoleWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetCloudConsoleWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudConsoleWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetCloudConsoleWithResponse), varargs...)
}

// GetCloudCredentialsWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetCloudCredentialsWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.GetCloudCredentialsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCloudCredentialsWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetCloudCredentialsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudCredentialsWithResponse indicates an expected call of GetCloudCredentialsWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetCloudCredentialsWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudCredentialsWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetCloudCredentialsWithResponse), varargs...)
}

// GetJobLogsWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetJobLogsWithResponse(arg0 context.Context, arg1, arg2 string, arg3 *Openapi.GetJobLogsParams, arg4 ...Openapi.RequestEditorFn) (*Openapi.GetJobLogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobLogsWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetJobLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobLogsWithResponse indicates an expected call of GetJobLogsWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetJobLogsWithResponse(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobLogsWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetJobLogsWithResponse), varargs...)
}

// GetRunWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetRunWithResponse(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*Openapi.GetRunResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRunWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetRunResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunWithResponse indicates an expected call of GetRunWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetRunWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetRunWithResponse), varargs...)
}

// GetScriptsWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetScriptsWithResponse(arg0 context.Context, arg1 *Openapi.GetScriptsParams, arg2 ...Openapi.RequestEditorFn) (*Openapi.GetScriptsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScriptsWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetScriptsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScriptsWithResponse indicates an expected call of GetScriptsWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetScriptsWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScriptsWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetScriptsWithResponse), varargs...)
}

// GetTestScriptRunLogsWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetTestScriptRunLogsWithResponse(arg0 context.Context, arg1, arg2 string, arg3 *Openapi.GetTestScriptRunLogsParams, arg4 ...Openapi.RequestEditorFn) (*Openapi.GetTestScriptRunLogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTestScriptRunLogsWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetTestScriptRunLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestScriptRunLogsWithResponse indicates an expected call of GetTestScriptRunLogsWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetTestScriptRunLogsWithResponse(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestScriptRunLogsWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetTestScriptRunLogsWithResponse), varargs...)
}

// GetTestScriptRunWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetTestScriptRunWithResponse(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*Openapi.GetTestScriptRunResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTestScriptRunWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.GetTestScriptRunResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestScriptRunWithResponse indicates an expected call of GetTestScriptRunWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetTestScriptRunWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestScriptRunWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetTestScriptRunWithResponse), varargs...)
}

// HeadBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) HeadBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.HeadBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.HeadBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadBackplaneClusterClusterIdWithResponse indicates an expected call of HeadBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) HeadBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).HeadBackplaneClusterClusterIdWithResponse), varargs...)
}

// LoginClusterWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) LoginClusterWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.LoginClusterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginClusterWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.LoginClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginClusterWithResponse indicates an expected call of LoginClusterWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) LoginClusterWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginClusterWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).LoginClusterWithResponse), varargs...)
}

// OptionsBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) OptionsBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.OptionsBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OptionsBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.OptionsBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OptionsBackplaneClusterClusterIdWithResponse indicates an expected call of OptionsBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) OptionsBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OptionsBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).OptionsBackplaneClusterClusterIdWithResponse), varargs...)
}

// PatchBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PatchBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.PatchBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.PatchBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchBackplaneClusterClusterIdWithResponse indicates an expected call of PatchBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PatchBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PatchBackplaneClusterClusterIdWithResponse), varargs...)
}

// PostBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.PostBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.PostBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBackplaneClusterClusterIdWithResponse indicates an expected call of PostBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostBackplaneClusterClusterIdWithResponse), varargs...)
}

// PutBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PutBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.PutBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.PutBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBackplaneClusterClusterIdWithResponse indicates an expected call of PutBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PutBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PutBackplaneClusterClusterIdWithResponse), varargs...)
}

// TraceBackplaneClusterClusterIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) TraceBackplaneClusterClusterIdWithResponse(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*Openapi.TraceBackplaneClusterClusterIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TraceBackplaneClusterClusterIdWithResponse", varargs...)
	ret0, _ := ret[0].(*Openapi.TraceBackplaneClusterClusterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBackplaneClusterClusterIdWithResponse indicates an expected call of TraceBackplaneClusterClusterIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) TraceBackplaneClusterClusterIdWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBackplaneClusterClusterIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).TraceBackplaneClusterClusterIdWithResponse), varargs...)
}

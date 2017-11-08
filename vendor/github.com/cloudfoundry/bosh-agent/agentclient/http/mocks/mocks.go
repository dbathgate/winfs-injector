// Automatically generated by MockGen. DO NOT EDIT!
// Source: agent_client_factory.go

package mocks

import (
	agentclient "github.com/cloudfoundry/bosh-agent/agentclient"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AgentClientFactory interface
type MockAgentClientFactory struct {
	ctrl     *gomock.Controller
	recorder *_MockAgentClientFactoryRecorder
}

// Recorder for MockAgentClientFactory (not exported)
type _MockAgentClientFactoryRecorder struct {
	mock *MockAgentClientFactory
}

func NewMockAgentClientFactory(ctrl *gomock.Controller) *MockAgentClientFactory {
	mock := &MockAgentClientFactory{ctrl: ctrl}
	mock.recorder = &_MockAgentClientFactoryRecorder{mock}
	return mock
}

func (_m *MockAgentClientFactory) EXPECT() *_MockAgentClientFactoryRecorder {
	return _m.recorder
}

func (_m *MockAgentClientFactory) NewAgentClient(directorID string, mbusURL string, caCert string) (agentclient.AgentClient, error) {
	ret := _m.ctrl.Call(_m, "NewAgentClient", directorID, mbusURL, caCert)
	ret0, _ := ret[0].(agentclient.AgentClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAgentClientFactoryRecorder) NewAgentClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewAgentClient", arg0, arg1, arg2)
}

// This file was generated by counterfeiter
package configfakes

import (
	"sync"

	"github.com/dpb587/ssoca/client/config"
)

type FakeManager struct {
	GetSourceStub        func() string
	getSourceMutex       sync.RWMutex
	getSourceArgsForCall []struct{}
	getSourceReturns     struct {
		result1 string
	}
	GetEnvironmentsStub        func() (config.EnvironmentsState, error)
	getEnvironmentsMutex       sync.RWMutex
	getEnvironmentsArgsForCall []struct{}
	getEnvironmentsReturns     struct {
		result1 config.EnvironmentsState
		result2 error
	}
	GetEnvironmentStub        func(string) (config.EnvironmentState, error)
	getEnvironmentMutex       sync.RWMutex
	getEnvironmentArgsForCall []struct {
		arg1 string
	}
	getEnvironmentReturns struct {
		result1 config.EnvironmentState
		result2 error
	}
	SetEnvironmentStub        func(config.EnvironmentState) error
	setEnvironmentMutex       sync.RWMutex
	setEnvironmentArgsForCall []struct {
		arg1 config.EnvironmentState
	}
	setEnvironmentReturns struct {
		result1 error
	}
	UnsetEnvironmentStub        func(string) error
	unsetEnvironmentMutex       sync.RWMutex
	unsetEnvironmentArgsForCall []struct {
		arg1 string
	}
	unsetEnvironmentReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) GetSource() string {
	fake.getSourceMutex.Lock()
	fake.getSourceArgsForCall = append(fake.getSourceArgsForCall, struct{}{})
	fake.recordInvocation("GetSource", []interface{}{})
	fake.getSourceMutex.Unlock()
	if fake.GetSourceStub != nil {
		return fake.GetSourceStub()
	}
	return fake.getSourceReturns.result1
}

func (fake *FakeManager) GetSourceCallCount() int {
	fake.getSourceMutex.RLock()
	defer fake.getSourceMutex.RUnlock()
	return len(fake.getSourceArgsForCall)
}

func (fake *FakeManager) GetSourceReturns(result1 string) {
	fake.GetSourceStub = nil
	fake.getSourceReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeManager) GetEnvironments() (config.EnvironmentsState, error) {
	fake.getEnvironmentsMutex.Lock()
	fake.getEnvironmentsArgsForCall = append(fake.getEnvironmentsArgsForCall, struct{}{})
	fake.recordInvocation("GetEnvironments", []interface{}{})
	fake.getEnvironmentsMutex.Unlock()
	if fake.GetEnvironmentsStub != nil {
		return fake.GetEnvironmentsStub()
	}
	return fake.getEnvironmentsReturns.result1, fake.getEnvironmentsReturns.result2
}

func (fake *FakeManager) GetEnvironmentsCallCount() int {
	fake.getEnvironmentsMutex.RLock()
	defer fake.getEnvironmentsMutex.RUnlock()
	return len(fake.getEnvironmentsArgsForCall)
}

func (fake *FakeManager) GetEnvironmentsReturns(result1 config.EnvironmentsState, result2 error) {
	fake.GetEnvironmentsStub = nil
	fake.getEnvironmentsReturns = struct {
		result1 config.EnvironmentsState
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetEnvironment(arg1 string) (config.EnvironmentState, error) {
	fake.getEnvironmentMutex.Lock()
	fake.getEnvironmentArgsForCall = append(fake.getEnvironmentArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetEnvironment", []interface{}{arg1})
	fake.getEnvironmentMutex.Unlock()
	if fake.GetEnvironmentStub != nil {
		return fake.GetEnvironmentStub(arg1)
	}
	return fake.getEnvironmentReturns.result1, fake.getEnvironmentReturns.result2
}

func (fake *FakeManager) GetEnvironmentCallCount() int {
	fake.getEnvironmentMutex.RLock()
	defer fake.getEnvironmentMutex.RUnlock()
	return len(fake.getEnvironmentArgsForCall)
}

func (fake *FakeManager) GetEnvironmentArgsForCall(i int) string {
	fake.getEnvironmentMutex.RLock()
	defer fake.getEnvironmentMutex.RUnlock()
	return fake.getEnvironmentArgsForCall[i].arg1
}

func (fake *FakeManager) GetEnvironmentReturns(result1 config.EnvironmentState, result2 error) {
	fake.GetEnvironmentStub = nil
	fake.getEnvironmentReturns = struct {
		result1 config.EnvironmentState
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) SetEnvironment(arg1 config.EnvironmentState) error {
	fake.setEnvironmentMutex.Lock()
	fake.setEnvironmentArgsForCall = append(fake.setEnvironmentArgsForCall, struct {
		arg1 config.EnvironmentState
	}{arg1})
	fake.recordInvocation("SetEnvironment", []interface{}{arg1})
	fake.setEnvironmentMutex.Unlock()
	if fake.SetEnvironmentStub != nil {
		return fake.SetEnvironmentStub(arg1)
	}
	return fake.setEnvironmentReturns.result1
}

func (fake *FakeManager) SetEnvironmentCallCount() int {
	fake.setEnvironmentMutex.RLock()
	defer fake.setEnvironmentMutex.RUnlock()
	return len(fake.setEnvironmentArgsForCall)
}

func (fake *FakeManager) SetEnvironmentArgsForCall(i int) config.EnvironmentState {
	fake.setEnvironmentMutex.RLock()
	defer fake.setEnvironmentMutex.RUnlock()
	return fake.setEnvironmentArgsForCall[i].arg1
}

func (fake *FakeManager) SetEnvironmentReturns(result1 error) {
	fake.SetEnvironmentStub = nil
	fake.setEnvironmentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UnsetEnvironment(arg1 string) error {
	fake.unsetEnvironmentMutex.Lock()
	fake.unsetEnvironmentArgsForCall = append(fake.unsetEnvironmentArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnsetEnvironment", []interface{}{arg1})
	fake.unsetEnvironmentMutex.Unlock()
	if fake.UnsetEnvironmentStub != nil {
		return fake.UnsetEnvironmentStub(arg1)
	}
	return fake.unsetEnvironmentReturns.result1
}

func (fake *FakeManager) UnsetEnvironmentCallCount() int {
	fake.unsetEnvironmentMutex.RLock()
	defer fake.unsetEnvironmentMutex.RUnlock()
	return len(fake.unsetEnvironmentArgsForCall)
}

func (fake *FakeManager) UnsetEnvironmentArgsForCall(i int) string {
	fake.unsetEnvironmentMutex.RLock()
	defer fake.unsetEnvironmentMutex.RUnlock()
	return fake.unsetEnvironmentArgsForCall[i].arg1
}

func (fake *FakeManager) UnsetEnvironmentReturns(result1 error) {
	fake.UnsetEnvironmentStub = nil
	fake.unsetEnvironmentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSourceMutex.RLock()
	defer fake.getSourceMutex.RUnlock()
	fake.getEnvironmentsMutex.RLock()
	defer fake.getEnvironmentsMutex.RUnlock()
	fake.getEnvironmentMutex.RLock()
	defer fake.getEnvironmentMutex.RUnlock()
	fake.setEnvironmentMutex.RLock()
	defer fake.setEnvironmentMutex.RUnlock()
	fake.unsetEnvironmentMutex.RLock()
	defer fake.unsetEnvironmentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ config.Manager = new(FakeManager)

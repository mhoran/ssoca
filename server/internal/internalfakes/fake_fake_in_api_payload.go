// This file was generated by counterfeiter
package internalfakes

import (
	"sync"

	"github.com/dpb587/ssoca/server/internal"
)

type FakeFakeInApiPayload struct {
	RouteStub        func() string
	routeMutex       sync.RWMutex
	routeArgsForCall []struct{}
	routeReturns     struct {
		result1 string
	}
	ExecuteStub        func(map[string]interface{})
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 map[string]interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFakeInApiPayload) Route() string {
	fake.routeMutex.Lock()
	fake.routeArgsForCall = append(fake.routeArgsForCall, struct{}{})
	fake.recordInvocation("Route", []interface{}{})
	fake.routeMutex.Unlock()
	if fake.RouteStub != nil {
		return fake.RouteStub()
	}
	return fake.routeReturns.result1
}

func (fake *FakeFakeInApiPayload) RouteCallCount() int {
	fake.routeMutex.RLock()
	defer fake.routeMutex.RUnlock()
	return len(fake.routeArgsForCall)
}

func (fake *FakeFakeInApiPayload) RouteReturns(result1 string) {
	fake.RouteStub = nil
	fake.routeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFakeInApiPayload) Execute(arg1 map[string]interface{}) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 map[string]interface{}
	}{arg1})
	fake.recordInvocation("Execute", []interface{}{arg1})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		fake.ExecuteStub(arg1)
	}
}

func (fake *FakeFakeInApiPayload) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeFakeInApiPayload) ExecuteArgsForCall(i int) map[string]interface{} {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].arg1
}

func (fake *FakeFakeInApiPayload) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.routeMutex.RLock()
	defer fake.routeMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFakeInApiPayload) recordInvocation(key string, args []interface{}) {
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

var _ internal.FakeInApiPayload = new(FakeFakeInApiPayload)

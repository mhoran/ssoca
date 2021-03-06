// Code generated by counterfeiter. DO NOT EDIT.
package filterfakes

import (
	"sync"

	"github.com/dpb587/ssoca/auth/authz/filter"
)

type FakeFilter struct {
	CreateStub        func(interface{}) (filter.Requirement, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 interface{}
	}
	createReturns struct {
		result1 filter.Requirement
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 filter.Requirement
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilter) Create(arg1 interface{}) (filter.Requirement, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeFilter) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeFilter) CreateArgsForCall(i int) interface{} {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeFilter) CreateReturns(result1 filter.Requirement, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 filter.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) CreateReturnsOnCall(i int, result1 filter.Requirement, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 filter.Requirement
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 filter.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFilter) recordInvocation(key string, args []interface{}) {
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

var _ filter.Filter = new(FakeFilter)

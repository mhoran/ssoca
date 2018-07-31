// Code generated by counterfeiter. DO NOT EDIT.
package filterfakes

import (
	"sync"

	"github.com/dpb587/ssoca/auth/authz/filter"
)

type FakeManager struct {
	AddStub        func(string, filter.Filter)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
		arg2 filter.Filter
	}
	GetStub        func(string) (filter.Filter, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 filter.Filter
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 filter.Filter
		result2 error
	}
	FiltersStub        func() []string
	filtersMutex       sync.RWMutex
	filtersArgsForCall []struct{}
	filtersReturns     struct {
		result1 []string
	}
	filtersReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) Add(arg1 string, arg2 filter.Filter) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
		arg2 filter.Filter
	}{arg1, arg2})
	fake.recordInvocation("Add", []interface{}{arg1, arg2})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(arg1, arg2)
	}
}

func (fake *FakeManager) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeManager) AddArgsForCall(i int) (string, filter.Filter) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].arg1, fake.addArgsForCall[i].arg2
}

func (fake *FakeManager) Get(arg1 string) (filter.Filter, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeManager) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeManager) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1
}

func (fake *FakeManager) GetReturns(result1 filter.Filter, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 filter.Filter
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetReturnsOnCall(i int, result1 filter.Filter, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 filter.Filter
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 filter.Filter
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Filters() []string {
	fake.filtersMutex.Lock()
	ret, specificReturn := fake.filtersReturnsOnCall[len(fake.filtersArgsForCall)]
	fake.filtersArgsForCall = append(fake.filtersArgsForCall, struct{}{})
	fake.recordInvocation("Filters", []interface{}{})
	fake.filtersMutex.Unlock()
	if fake.FiltersStub != nil {
		return fake.FiltersStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.filtersReturns.result1
}

func (fake *FakeManager) FiltersCallCount() int {
	fake.filtersMutex.RLock()
	defer fake.filtersMutex.RUnlock()
	return len(fake.filtersArgsForCall)
}

func (fake *FakeManager) FiltersReturns(result1 []string) {
	fake.FiltersStub = nil
	fake.filtersReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeManager) FiltersReturnsOnCall(i int, result1 []string) {
	fake.FiltersStub = nil
	if fake.filtersReturnsOnCall == nil {
		fake.filtersReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.filtersReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.filtersMutex.RLock()
	defer fake.filtersMutex.RUnlock()
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

var _ filter.Manager = new(FakeManager)

// Code generated by counterfeiter. DO NOT EDIT.
package helperfakes

import (
	"sync"

	boshuaa "github.com/cloudfoundry/bosh-cli/uaa"
	"github.com/dpb587/ssoca/auth/authn/uaa/helper"
)

type FakeClientFactory struct {
	CreateClientStub        func(string, string, string, string) (boshuaa.UAA, error)
	createClientMutex       sync.RWMutex
	createClientArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	createClientReturns struct {
		result1 boshuaa.UAA
		result2 error
	}
	createClientReturnsOnCall map[int]struct {
		result1 boshuaa.UAA
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientFactory) CreateClient(arg1 string, arg2 string, arg3 string, arg4 string) (boshuaa.UAA, error) {
	fake.createClientMutex.Lock()
	ret, specificReturn := fake.createClientReturnsOnCall[len(fake.createClientArgsForCall)]
	fake.createClientArgsForCall = append(fake.createClientArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateClient", []interface{}{arg1, arg2, arg3, arg4})
	fake.createClientMutex.Unlock()
	if fake.CreateClientStub != nil {
		return fake.CreateClientStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createClientReturns.result1, fake.createClientReturns.result2
}

func (fake *FakeClientFactory) CreateClientCallCount() int {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return len(fake.createClientArgsForCall)
}

func (fake *FakeClientFactory) CreateClientArgsForCall(i int) (string, string, string, string) {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return fake.createClientArgsForCall[i].arg1, fake.createClientArgsForCall[i].arg2, fake.createClientArgsForCall[i].arg3, fake.createClientArgsForCall[i].arg4
}

func (fake *FakeClientFactory) CreateClientReturns(result1 boshuaa.UAA, result2 error) {
	fake.CreateClientStub = nil
	fake.createClientReturns = struct {
		result1 boshuaa.UAA
		result2 error
	}{result1, result2}
}

func (fake *FakeClientFactory) CreateClientReturnsOnCall(i int, result1 boshuaa.UAA, result2 error) {
	fake.CreateClientStub = nil
	if fake.createClientReturnsOnCall == nil {
		fake.createClientReturnsOnCall = make(map[int]struct {
			result1 boshuaa.UAA
			result2 error
		})
	}
	fake.createClientReturnsOnCall[i] = struct {
		result1 boshuaa.UAA
		result2 error
	}{result1, result2}
}

func (fake *FakeClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ helper.ClientFactory = new(FakeClientFactory)

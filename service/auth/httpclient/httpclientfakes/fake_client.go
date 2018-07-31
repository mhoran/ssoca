// Code generated by counterfeiter. DO NOT EDIT.
package httpclientfakes

import (
	"sync"

	"github.com/dpb587/ssoca/service/auth/api"
	"github.com/dpb587/ssoca/service/auth/httpclient"
)

type FakeClient struct {
	GetInfoStub        func() (api.InfoResponse, error)
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct{}
	getInfoReturns     struct {
		result1 api.InfoResponse
		result2 error
	}
	getInfoReturnsOnCall map[int]struct {
		result1 api.InfoResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) GetInfo() (api.InfoResponse, error) {
	fake.getInfoMutex.Lock()
	ret, specificReturn := fake.getInfoReturnsOnCall[len(fake.getInfoArgsForCall)]
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct{}{})
	fake.recordInvocation("GetInfo", []interface{}{})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInfoReturns.result1, fake.getInfoReturns.result2
}

func (fake *FakeClient) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeClient) GetInfoReturns(result1 api.InfoResponse, result2 error) {
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 api.InfoResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetInfoReturnsOnCall(i int, result1 api.InfoResponse, result2 error) {
	fake.GetInfoStub = nil
	if fake.getInfoReturnsOnCall == nil {
		fake.getInfoReturnsOnCall = make(map[int]struct {
			result1 api.InfoResponse
			result2 error
		})
	}
	fake.getInfoReturnsOnCall[i] = struct {
		result1 api.InfoResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ httpclient.Client = new(FakeClient)

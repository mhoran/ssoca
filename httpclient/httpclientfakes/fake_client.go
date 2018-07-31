// Code generated by counterfeiter. DO NOT EDIT.
package httpclientfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/dpb587/ssoca/httpclient"
)

type FakeClient struct {
	APIGetStub        func(string, interface{}) error
	aPIGetMutex       sync.RWMutex
	aPIGetArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	aPIGetReturns struct {
		result1 error
	}
	aPIGetReturnsOnCall map[int]struct {
		result1 error
	}
	APIPostStub        func(string, interface{}, interface{}) error
	aPIPostMutex       sync.RWMutex
	aPIPostArgsForCall []struct {
		arg1 string
		arg2 interface{}
		arg3 interface{}
	}
	aPIPostReturns struct {
		result1 error
	}
	aPIPostReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string) (*http.Response, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 *http.Response
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	PostStub        func(string, string, io.Reader) (*http.Response, error)
	postMutex       sync.RWMutex
	postArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 io.Reader
	}
	postReturns struct {
		result1 *http.Response
		result2 error
	}
	postReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) APIGet(arg1 string, arg2 interface{}) error {
	fake.aPIGetMutex.Lock()
	ret, specificReturn := fake.aPIGetReturnsOnCall[len(fake.aPIGetArgsForCall)]
	fake.aPIGetArgsForCall = append(fake.aPIGetArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("APIGet", []interface{}{arg1, arg2})
	fake.aPIGetMutex.Unlock()
	if fake.APIGetStub != nil {
		return fake.APIGetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.aPIGetReturns.result1
}

func (fake *FakeClient) APIGetCallCount() int {
	fake.aPIGetMutex.RLock()
	defer fake.aPIGetMutex.RUnlock()
	return len(fake.aPIGetArgsForCall)
}

func (fake *FakeClient) APIGetArgsForCall(i int) (string, interface{}) {
	fake.aPIGetMutex.RLock()
	defer fake.aPIGetMutex.RUnlock()
	return fake.aPIGetArgsForCall[i].arg1, fake.aPIGetArgsForCall[i].arg2
}

func (fake *FakeClient) APIGetReturns(result1 error) {
	fake.APIGetStub = nil
	fake.aPIGetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) APIGetReturnsOnCall(i int, result1 error) {
	fake.APIGetStub = nil
	if fake.aPIGetReturnsOnCall == nil {
		fake.aPIGetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.aPIGetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) APIPost(arg1 string, arg2 interface{}, arg3 interface{}) error {
	fake.aPIPostMutex.Lock()
	ret, specificReturn := fake.aPIPostReturnsOnCall[len(fake.aPIPostArgsForCall)]
	fake.aPIPostArgsForCall = append(fake.aPIPostArgsForCall, struct {
		arg1 string
		arg2 interface{}
		arg3 interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("APIPost", []interface{}{arg1, arg2, arg3})
	fake.aPIPostMutex.Unlock()
	if fake.APIPostStub != nil {
		return fake.APIPostStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.aPIPostReturns.result1
}

func (fake *FakeClient) APIPostCallCount() int {
	fake.aPIPostMutex.RLock()
	defer fake.aPIPostMutex.RUnlock()
	return len(fake.aPIPostArgsForCall)
}

func (fake *FakeClient) APIPostArgsForCall(i int) (string, interface{}, interface{}) {
	fake.aPIPostMutex.RLock()
	defer fake.aPIPostMutex.RUnlock()
	return fake.aPIPostArgsForCall[i].arg1, fake.aPIPostArgsForCall[i].arg2, fake.aPIPostArgsForCall[i].arg3
}

func (fake *FakeClient) APIPostReturns(result1 error) {
	fake.APIPostStub = nil
	fake.aPIPostReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) APIPostReturnsOnCall(i int, result1 error) {
	fake.APIPostStub = nil
	if fake.aPIPostReturnsOnCall == nil {
		fake.aPIPostReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.aPIPostReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Get(arg1 string) (*http.Response, error) {
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

func (fake *FakeClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeClient) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1
}

func (fake *FakeClient) GetReturns(result1 *http.Response, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Post(arg1 string, arg2 string, arg3 io.Reader) (*http.Response, error) {
	fake.postMutex.Lock()
	ret, specificReturn := fake.postReturnsOnCall[len(fake.postArgsForCall)]
	fake.postArgsForCall = append(fake.postArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 io.Reader
	}{arg1, arg2, arg3})
	fake.recordInvocation("Post", []interface{}{arg1, arg2, arg3})
	fake.postMutex.Unlock()
	if fake.PostStub != nil {
		return fake.PostStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.postReturns.result1, fake.postReturns.result2
}

func (fake *FakeClient) PostCallCount() int {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return len(fake.postArgsForCall)
}

func (fake *FakeClient) PostArgsForCall(i int) (string, string, io.Reader) {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return fake.postArgsForCall[i].arg1, fake.postArgsForCall[i].arg2, fake.postArgsForCall[i].arg3
}

func (fake *FakeClient) PostReturns(result1 *http.Response, result2 error) {
	fake.PostStub = nil
	fake.postReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PostReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.PostStub = nil
	if fake.postReturnsOnCall == nil {
		fake.postReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.postReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIGetMutex.RLock()
	defer fake.aPIGetMutex.RUnlock()
	fake.aPIPostMutex.RLock()
	defer fake.aPIPostMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
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

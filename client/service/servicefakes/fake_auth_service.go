// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"net/http"
	"sync"

	"github.com/dpb587/ssoca/client/service"
	"github.com/dpb587/ssoca/service/env/api"
)

type FakeAuthService struct {
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 string
	}
	versionReturnsOnCall map[int]struct {
		result1 string
	}
	AuthRequestStub        func(*http.Request) error
	authRequestMutex       sync.RWMutex
	authRequestArgsForCall []struct {
		arg1 *http.Request
	}
	authRequestReturns struct {
		result1 error
	}
	authRequestReturnsOnCall map[int]struct {
		result1 error
	}
	AuthLoginStub        func(api.InfoServiceResponse) (interface{}, error)
	authLoginMutex       sync.RWMutex
	authLoginArgsForCall []struct {
		arg1 api.InfoServiceResponse
	}
	authLoginReturns struct {
		result1 interface{}
		result2 error
	}
	authLoginReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	AuthLogoutStub        func() error
	authLogoutMutex       sync.RWMutex
	authLogoutArgsForCall []struct{}
	authLogoutReturns     struct {
		result1 error
	}
	authLogoutReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthService) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.typeReturns.result1
}

func (fake *FakeAuthService) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeAuthService) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthService) TypeReturnsOnCall(i int, result1 string) {
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthService) Version() string {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeAuthService) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeAuthService) VersionReturns(result1 string) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthService) VersionReturnsOnCall(i int, result1 string) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthService) AuthRequest(arg1 *http.Request) error {
	fake.authRequestMutex.Lock()
	ret, specificReturn := fake.authRequestReturnsOnCall[len(fake.authRequestArgsForCall)]
	fake.authRequestArgsForCall = append(fake.authRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("AuthRequest", []interface{}{arg1})
	fake.authRequestMutex.Unlock()
	if fake.AuthRequestStub != nil {
		return fake.AuthRequestStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.authRequestReturns.result1
}

func (fake *FakeAuthService) AuthRequestCallCount() int {
	fake.authRequestMutex.RLock()
	defer fake.authRequestMutex.RUnlock()
	return len(fake.authRequestArgsForCall)
}

func (fake *FakeAuthService) AuthRequestArgsForCall(i int) *http.Request {
	fake.authRequestMutex.RLock()
	defer fake.authRequestMutex.RUnlock()
	return fake.authRequestArgsForCall[i].arg1
}

func (fake *FakeAuthService) AuthRequestReturns(result1 error) {
	fake.AuthRequestStub = nil
	fake.authRequestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) AuthRequestReturnsOnCall(i int, result1 error) {
	fake.AuthRequestStub = nil
	if fake.authRequestReturnsOnCall == nil {
		fake.authRequestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.authRequestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) AuthLogin(arg1 api.InfoServiceResponse) (interface{}, error) {
	fake.authLoginMutex.Lock()
	ret, specificReturn := fake.authLoginReturnsOnCall[len(fake.authLoginArgsForCall)]
	fake.authLoginArgsForCall = append(fake.authLoginArgsForCall, struct {
		arg1 api.InfoServiceResponse
	}{arg1})
	fake.recordInvocation("AuthLogin", []interface{}{arg1})
	fake.authLoginMutex.Unlock()
	if fake.AuthLoginStub != nil {
		return fake.AuthLoginStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.authLoginReturns.result1, fake.authLoginReturns.result2
}

func (fake *FakeAuthService) AuthLoginCallCount() int {
	fake.authLoginMutex.RLock()
	defer fake.authLoginMutex.RUnlock()
	return len(fake.authLoginArgsForCall)
}

func (fake *FakeAuthService) AuthLoginArgsForCall(i int) api.InfoServiceResponse {
	fake.authLoginMutex.RLock()
	defer fake.authLoginMutex.RUnlock()
	return fake.authLoginArgsForCall[i].arg1
}

func (fake *FakeAuthService) AuthLoginReturns(result1 interface{}, result2 error) {
	fake.AuthLoginStub = nil
	fake.authLoginReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthService) AuthLoginReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.AuthLoginStub = nil
	if fake.authLoginReturnsOnCall == nil {
		fake.authLoginReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.authLoginReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthService) AuthLogout() error {
	fake.authLogoutMutex.Lock()
	ret, specificReturn := fake.authLogoutReturnsOnCall[len(fake.authLogoutArgsForCall)]
	fake.authLogoutArgsForCall = append(fake.authLogoutArgsForCall, struct{}{})
	fake.recordInvocation("AuthLogout", []interface{}{})
	fake.authLogoutMutex.Unlock()
	if fake.AuthLogoutStub != nil {
		return fake.AuthLogoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.authLogoutReturns.result1
}

func (fake *FakeAuthService) AuthLogoutCallCount() int {
	fake.authLogoutMutex.RLock()
	defer fake.authLogoutMutex.RUnlock()
	return len(fake.authLogoutArgsForCall)
}

func (fake *FakeAuthService) AuthLogoutReturns(result1 error) {
	fake.AuthLogoutStub = nil
	fake.authLogoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) AuthLogoutReturnsOnCall(i int, result1 error) {
	fake.AuthLogoutStub = nil
	if fake.authLogoutReturnsOnCall == nil {
		fake.authLogoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.authLogoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.authRequestMutex.RLock()
	defer fake.authRequestMutex.RUnlock()
	fake.authLoginMutex.RLock()
	defer fake.authLoginMutex.RUnlock()
	fake.authLogoutMutex.RLock()
	defer fake.authLogoutMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAuthService) recordInvocation(key string, args []interface{}) {
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

var _ service.AuthService = new(FakeAuthService)
// This file was generated by counterfeiter
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
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 string
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns     struct {
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
	AuthLoginStub        func(api.InfoServiceResponse) (interface{}, error)
	authLoginMutex       sync.RWMutex
	authLoginArgsForCall []struct {
		arg1 api.InfoServiceResponse
	}
	authLoginReturns struct {
		result1 interface{}
		result2 error
	}
	AuthLogoutStub        func() error
	authLogoutMutex       sync.RWMutex
	authLogoutArgsForCall []struct{}
	authLogoutReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthService) Type() string {
	fake.typeMutex.Lock()
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
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

func (fake *FakeAuthService) Version() string {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
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

func (fake *FakeAuthService) Description() string {
	fake.descriptionMutex.Lock()
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.recordInvocation("Description", []interface{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	}
	return fake.descriptionReturns.result1
}

func (fake *FakeAuthService) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeAuthService) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthService) AuthRequest(arg1 *http.Request) error {
	fake.authRequestMutex.Lock()
	fake.authRequestArgsForCall = append(fake.authRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("AuthRequest", []interface{}{arg1})
	fake.authRequestMutex.Unlock()
	if fake.AuthRequestStub != nil {
		return fake.AuthRequestStub(arg1)
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

func (fake *FakeAuthService) AuthLogin(arg1 api.InfoServiceResponse) (interface{}, error) {
	fake.authLoginMutex.Lock()
	fake.authLoginArgsForCall = append(fake.authLoginArgsForCall, struct {
		arg1 api.InfoServiceResponse
	}{arg1})
	fake.recordInvocation("AuthLogin", []interface{}{arg1})
	fake.authLoginMutex.Unlock()
	if fake.AuthLoginStub != nil {
		return fake.AuthLoginStub(arg1)
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

func (fake *FakeAuthService) AuthLogout() error {
	fake.authLogoutMutex.Lock()
	fake.authLogoutArgsForCall = append(fake.authLogoutArgsForCall, struct{}{})
	fake.recordInvocation("AuthLogout", []interface{}{})
	fake.authLogoutMutex.Unlock()
	if fake.AuthLogoutStub != nil {
		return fake.AuthLogoutStub()
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

func (fake *FakeAuthService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
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

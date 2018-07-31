// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"net/http"
	"sync"

	"github.com/dpb587/ssoca/auth"
	"github.com/dpb587/ssoca/server/service"
	"github.com/dpb587/ssoca/server/service/req"
)

type FakeAuthService struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
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
	MetadataStub        func() interface{}
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns     struct {
		result1 interface{}
	}
	metadataReturnsOnCall map[int]struct {
		result1 interface{}
	}
	GetRoutesStub        func() []req.RouteHandler
	getRoutesMutex       sync.RWMutex
	getRoutesArgsForCall []struct{}
	getRoutesReturns     struct {
		result1 []req.RouteHandler
	}
	getRoutesReturnsOnCall map[int]struct {
		result1 []req.RouteHandler
	}
	VerifyAuthorizationStub        func(http.Request, *auth.Token) error
	verifyAuthorizationMutex       sync.RWMutex
	verifyAuthorizationArgsForCall []struct {
		arg1 http.Request
		arg2 *auth.Token
	}
	verifyAuthorizationReturns struct {
		result1 error
	}
	verifyAuthorizationReturnsOnCall map[int]struct {
		result1 error
	}
	ParseRequestAuthStub        func(http.Request) (*auth.Token, error)
	parseRequestAuthMutex       sync.RWMutex
	parseRequestAuthArgsForCall []struct {
		arg1 http.Request
	}
	parseRequestAuthReturns struct {
		result1 *auth.Token
		result2 error
	}
	parseRequestAuthReturnsOnCall map[int]struct {
		result1 *auth.Token
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthService) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeAuthService) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeAuthService) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthService) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
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

func (fake *FakeAuthService) Metadata() interface{} {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.metadataReturns.result1
}

func (fake *FakeAuthService) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeAuthService) MetadataReturns(result1 interface{}) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeAuthService) MetadataReturnsOnCall(i int, result1 interface{}) {
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeAuthService) GetRoutes() []req.RouteHandler {
	fake.getRoutesMutex.Lock()
	ret, specificReturn := fake.getRoutesReturnsOnCall[len(fake.getRoutesArgsForCall)]
	fake.getRoutesArgsForCall = append(fake.getRoutesArgsForCall, struct{}{})
	fake.recordInvocation("GetRoutes", []interface{}{})
	fake.getRoutesMutex.Unlock()
	if fake.GetRoutesStub != nil {
		return fake.GetRoutesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getRoutesReturns.result1
}

func (fake *FakeAuthService) GetRoutesCallCount() int {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return len(fake.getRoutesArgsForCall)
}

func (fake *FakeAuthService) GetRoutesReturns(result1 []req.RouteHandler) {
	fake.GetRoutesStub = nil
	fake.getRoutesReturns = struct {
		result1 []req.RouteHandler
	}{result1}
}

func (fake *FakeAuthService) GetRoutesReturnsOnCall(i int, result1 []req.RouteHandler) {
	fake.GetRoutesStub = nil
	if fake.getRoutesReturnsOnCall == nil {
		fake.getRoutesReturnsOnCall = make(map[int]struct {
			result1 []req.RouteHandler
		})
	}
	fake.getRoutesReturnsOnCall[i] = struct {
		result1 []req.RouteHandler
	}{result1}
}

func (fake *FakeAuthService) VerifyAuthorization(arg1 http.Request, arg2 *auth.Token) error {
	fake.verifyAuthorizationMutex.Lock()
	ret, specificReturn := fake.verifyAuthorizationReturnsOnCall[len(fake.verifyAuthorizationArgsForCall)]
	fake.verifyAuthorizationArgsForCall = append(fake.verifyAuthorizationArgsForCall, struct {
		arg1 http.Request
		arg2 *auth.Token
	}{arg1, arg2})
	fake.recordInvocation("VerifyAuthorization", []interface{}{arg1, arg2})
	fake.verifyAuthorizationMutex.Unlock()
	if fake.VerifyAuthorizationStub != nil {
		return fake.VerifyAuthorizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyAuthorizationReturns.result1
}

func (fake *FakeAuthService) VerifyAuthorizationCallCount() int {
	fake.verifyAuthorizationMutex.RLock()
	defer fake.verifyAuthorizationMutex.RUnlock()
	return len(fake.verifyAuthorizationArgsForCall)
}

func (fake *FakeAuthService) VerifyAuthorizationArgsForCall(i int) (http.Request, *auth.Token) {
	fake.verifyAuthorizationMutex.RLock()
	defer fake.verifyAuthorizationMutex.RUnlock()
	return fake.verifyAuthorizationArgsForCall[i].arg1, fake.verifyAuthorizationArgsForCall[i].arg2
}

func (fake *FakeAuthService) VerifyAuthorizationReturns(result1 error) {
	fake.VerifyAuthorizationStub = nil
	fake.verifyAuthorizationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) VerifyAuthorizationReturnsOnCall(i int, result1 error) {
	fake.VerifyAuthorizationStub = nil
	if fake.verifyAuthorizationReturnsOnCall == nil {
		fake.verifyAuthorizationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyAuthorizationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) ParseRequestAuth(arg1 http.Request) (*auth.Token, error) {
	fake.parseRequestAuthMutex.Lock()
	ret, specificReturn := fake.parseRequestAuthReturnsOnCall[len(fake.parseRequestAuthArgsForCall)]
	fake.parseRequestAuthArgsForCall = append(fake.parseRequestAuthArgsForCall, struct {
		arg1 http.Request
	}{arg1})
	fake.recordInvocation("ParseRequestAuth", []interface{}{arg1})
	fake.parseRequestAuthMutex.Unlock()
	if fake.ParseRequestAuthStub != nil {
		return fake.ParseRequestAuthStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseRequestAuthReturns.result1, fake.parseRequestAuthReturns.result2
}

func (fake *FakeAuthService) ParseRequestAuthCallCount() int {
	fake.parseRequestAuthMutex.RLock()
	defer fake.parseRequestAuthMutex.RUnlock()
	return len(fake.parseRequestAuthArgsForCall)
}

func (fake *FakeAuthService) ParseRequestAuthArgsForCall(i int) http.Request {
	fake.parseRequestAuthMutex.RLock()
	defer fake.parseRequestAuthMutex.RUnlock()
	return fake.parseRequestAuthArgsForCall[i].arg1
}

func (fake *FakeAuthService) ParseRequestAuthReturns(result1 *auth.Token, result2 error) {
	fake.ParseRequestAuthStub = nil
	fake.parseRequestAuthReturns = struct {
		result1 *auth.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthService) ParseRequestAuthReturnsOnCall(i int, result1 *auth.Token, result2 error) {
	fake.ParseRequestAuthStub = nil
	if fake.parseRequestAuthReturnsOnCall == nil {
		fake.parseRequestAuthReturnsOnCall = make(map[int]struct {
			result1 *auth.Token
			result2 error
		})
	}
	fake.parseRequestAuthReturnsOnCall[i] = struct {
		result1 *auth.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	fake.verifyAuthorizationMutex.RLock()
	defer fake.verifyAuthorizationMutex.RUnlock()
	fake.parseRequestAuthMutex.RLock()
	defer fake.parseRequestAuthMutex.RUnlock()
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

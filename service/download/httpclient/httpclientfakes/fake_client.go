// Code generated by counterfeiter. DO NOT EDIT.
package httpclientfakes

import (
	"io"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/dpb587/ssoca/service/download/api"
	"github.com/dpb587/ssoca/service/download/httpclient"
)

type FakeClient struct {
	GetMetadataStub        func() (api.MetadataResponse, error)
	getMetadataMutex       sync.RWMutex
	getMetadataArgsForCall []struct{}
	getMetadataReturns     struct {
		result1 api.MetadataResponse
		result2 error
	}
	getMetadataReturnsOnCall map[int]struct {
		result1 api.MetadataResponse
		result2 error
	}
	GetListStub        func() (api.ListResponse, error)
	getListMutex       sync.RWMutex
	getListArgsForCall []struct{}
	getListReturns     struct {
		result1 api.ListResponse
		result2 error
	}
	getListReturnsOnCall map[int]struct {
		result1 api.ListResponse
		result2 error
	}
	DownloadStub        func(string, io.ReadWriteSeeker, *pb.ProgressBar) error
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		arg1 string
		arg2 io.ReadWriteSeeker
		arg3 *pb.ProgressBar
	}
	downloadReturns struct {
		result1 error
	}
	downloadReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) GetMetadata() (api.MetadataResponse, error) {
	fake.getMetadataMutex.Lock()
	ret, specificReturn := fake.getMetadataReturnsOnCall[len(fake.getMetadataArgsForCall)]
	fake.getMetadataArgsForCall = append(fake.getMetadataArgsForCall, struct{}{})
	fake.recordInvocation("GetMetadata", []interface{}{})
	fake.getMetadataMutex.Unlock()
	if fake.GetMetadataStub != nil {
		return fake.GetMetadataStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getMetadataReturns.result1, fake.getMetadataReturns.result2
}

func (fake *FakeClient) GetMetadataCallCount() int {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return len(fake.getMetadataArgsForCall)
}

func (fake *FakeClient) GetMetadataReturns(result1 api.MetadataResponse, result2 error) {
	fake.GetMetadataStub = nil
	fake.getMetadataReturns = struct {
		result1 api.MetadataResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetMetadataReturnsOnCall(i int, result1 api.MetadataResponse, result2 error) {
	fake.GetMetadataStub = nil
	if fake.getMetadataReturnsOnCall == nil {
		fake.getMetadataReturnsOnCall = make(map[int]struct {
			result1 api.MetadataResponse
			result2 error
		})
	}
	fake.getMetadataReturnsOnCall[i] = struct {
		result1 api.MetadataResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetList() (api.ListResponse, error) {
	fake.getListMutex.Lock()
	ret, specificReturn := fake.getListReturnsOnCall[len(fake.getListArgsForCall)]
	fake.getListArgsForCall = append(fake.getListArgsForCall, struct{}{})
	fake.recordInvocation("GetList", []interface{}{})
	fake.getListMutex.Unlock()
	if fake.GetListStub != nil {
		return fake.GetListStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getListReturns.result1, fake.getListReturns.result2
}

func (fake *FakeClient) GetListCallCount() int {
	fake.getListMutex.RLock()
	defer fake.getListMutex.RUnlock()
	return len(fake.getListArgsForCall)
}

func (fake *FakeClient) GetListReturns(result1 api.ListResponse, result2 error) {
	fake.GetListStub = nil
	fake.getListReturns = struct {
		result1 api.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetListReturnsOnCall(i int, result1 api.ListResponse, result2 error) {
	fake.GetListStub = nil
	if fake.getListReturnsOnCall == nil {
		fake.getListReturnsOnCall = make(map[int]struct {
			result1 api.ListResponse
			result2 error
		})
	}
	fake.getListReturnsOnCall[i] = struct {
		result1 api.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Download(arg1 string, arg2 io.ReadWriteSeeker, arg3 *pb.ProgressBar) error {
	fake.downloadMutex.Lock()
	ret, specificReturn := fake.downloadReturnsOnCall[len(fake.downloadArgsForCall)]
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		arg1 string
		arg2 io.ReadWriteSeeker
		arg3 *pb.ProgressBar
	}{arg1, arg2, arg3})
	fake.recordInvocation("Download", []interface{}{arg1, arg2, arg3})
	fake.downloadMutex.Unlock()
	if fake.DownloadStub != nil {
		return fake.DownloadStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.downloadReturns.result1
}

func (fake *FakeClient) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeClient) DownloadArgsForCall(i int) (string, io.ReadWriteSeeker, *pb.ProgressBar) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return fake.downloadArgsForCall[i].arg1, fake.downloadArgsForCall[i].arg2, fake.downloadArgsForCall[i].arg3
}

func (fake *FakeClient) DownloadReturns(result1 error) {
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DownloadReturnsOnCall(i int, result1 error) {
	fake.DownloadStub = nil
	if fake.downloadReturnsOnCall == nil {
		fake.downloadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	fake.getListMutex.RLock()
	defer fake.getListMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
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

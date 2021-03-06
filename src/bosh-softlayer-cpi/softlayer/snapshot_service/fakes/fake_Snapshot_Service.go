// This file was generated by counterfeiter
package fakes

import (
	snapshot "bosh-softlayer-cpi/softlayer/snapshot_service"
	"sync"
)

type FakeService struct {
	CreateStub        func(diskID int, description string) (int, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		diskID      int
		description string
	}
	createReturns struct {
		result1 int
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	DeleteStub        func(id int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		id int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeService) Create(diskID int, description string) (int, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		diskID      int
		description string
	}{diskID, description})
	fake.recordInvocation("Create", []interface{}{diskID, description})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(diskID, description)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeService) CreateArgsForCall(i int) (int, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].diskID, fake.createArgsForCall[i].description
}

func (fake *FakeService) CreateReturns(result1 int, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeService) CreateReturnsOnCall(i int, result1 int, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeService) Delete(id int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("Delete", []interface{}{id})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeService) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeService) DeleteArgsForCall(i int) int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].id
}

func (fake *FakeService) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeService) recordInvocation(key string, args []interface{}) {
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

var _ snapshot.Service = new(FakeService)

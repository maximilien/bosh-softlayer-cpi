// This file was generated by counterfeiter
package fakes

import (
	stemcell "bosh-softlayer-cpi/softlayer/stemcell_service"
	"sync"
)

type FakeService struct {
	FindStub        func(id int) (string, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		id int
	}
	findReturns struct {
		result1 string
		result2 error
	}
	findReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeService) Find(id int) (string, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("Find", []interface{}{id})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findReturns.result1, fake.findReturns.result2
}

func (fake *FakeService) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeService) FindArgsForCall(i int) int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].id
}

func (fake *FakeService) FindReturns(result1 string, result2 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeService) FindReturnsOnCall(i int, result1 string, result2 error) {
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
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

var _ stemcell.Service = new(FakeService)
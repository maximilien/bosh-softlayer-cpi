// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/client/vm"
	"github.com/go-openapi/runtime"
)

type FakeSoftLayerPoolClient struct {
	AddVMStub        func(params *vm.AddVMParams) (*vm.AddVMOK, error)
	addVMMutex       sync.RWMutex
	addVMArgsForCall []struct {
		params *vm.AddVMParams
	}
	addVMReturns struct {
		result1 *vm.AddVMOK
		result2 error
	}
	DeleteVMStub        func(params *vm.DeleteVMParams) (*vm.DeleteVMNoContent, error)
	deleteVMMutex       sync.RWMutex
	deleteVMArgsForCall []struct {
		params *vm.DeleteVMParams
	}
	deleteVMReturns struct {
		result1 *vm.DeleteVMNoContent
		result2 error
	}
	FindVmsByDeploymentStub        func(params *vm.FindVmsByDeploymentParams) (*vm.FindVmsByDeploymentOK, error)
	findVmsByDeploymentMutex       sync.RWMutex
	findVmsByDeploymentArgsForCall []struct {
		params *vm.FindVmsByDeploymentParams
	}
	findVmsByDeploymentReturns struct {
		result1 *vm.FindVmsByDeploymentOK
		result2 error
	}
	FindVmsByFiltersStub        func(params *vm.FindVmsByFiltersParams) (*vm.FindVmsByFiltersOK, error)
	findVmsByFiltersMutex       sync.RWMutex
	findVmsByFiltersArgsForCall []struct {
		params *vm.FindVmsByFiltersParams
	}
	findVmsByFiltersReturns struct {
		result1 *vm.FindVmsByFiltersOK
		result2 error
	}
	FindVmsByStatesStub        func(params *vm.FindVmsByStatesParams) (*vm.FindVmsByStatesOK, error)
	findVmsByStatesMutex       sync.RWMutex
	findVmsByStatesArgsForCall []struct {
		params *vm.FindVmsByStatesParams
	}
	findVmsByStatesReturns struct {
		result1 *vm.FindVmsByStatesOK
		result2 error
	}
	GetVMByCidStub        func(params *vm.GetVMByCidParams) (*vm.GetVMByCidOK, error)
	getVMByCidMutex       sync.RWMutex
	getVMByCidArgsForCall []struct {
		params *vm.GetVMByCidParams
	}
	getVMByCidReturns struct {
		result1 *vm.GetVMByCidOK
		result2 error
	}
	ListVMStub        func(params *vm.ListVMParams) (*vm.ListVMOK, error)
	listVMMutex       sync.RWMutex
	listVMArgsForCall []struct {
		params *vm.ListVMParams
	}
	listVMReturns struct {
		result1 *vm.ListVMOK
		result2 error
	}
	OrderVMByFilterStub        func(params *vm.OrderVMByFilterParams) (*vm.OrderVMByFilterOK, error)
	orderVMByFilterMutex       sync.RWMutex
	orderVMByFilterArgsForCall []struct {
		params *vm.OrderVMByFilterParams
	}
	orderVMByFilterReturns struct {
		result1 *vm.OrderVMByFilterOK
		result2 error
	}
	UpdateVMStub        func(params *vm.UpdateVMParams) (*vm.UpdateVMOK, error)
	updateVMMutex       sync.RWMutex
	updateVMArgsForCall []struct {
		params *vm.UpdateVMParams
	}
	updateVMReturns struct {
		result1 *vm.UpdateVMOK
		result2 error
	}
	UpdateVMWithStateStub        func(params *vm.UpdateVMWithStateParams) (*vm.UpdateVMWithStateOK, error)
	updateVMWithStateMutex       sync.RWMutex
	updateVMWithStateArgsForCall []struct {
		params *vm.UpdateVMWithStateParams
	}
	updateVMWithStateReturns struct {
		result1 *vm.UpdateVMWithStateOK
		result2 error
	}
	SetTransportStub        func(transport runtime.ClientTransport)
	setTransportMutex       sync.RWMutex
	setTransportArgsForCall []struct {
		transport runtime.ClientTransport
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSoftLayerPoolClient) AddVM(params *vm.AddVMParams) (*vm.AddVMOK, error) {
	fake.addVMMutex.Lock()
	fake.addVMArgsForCall = append(fake.addVMArgsForCall, struct {
		params *vm.AddVMParams
	}{params})
	fake.recordInvocation("AddVM", []interface{}{params})
	fake.addVMMutex.Unlock()
	if fake.AddVMStub != nil {
		return fake.AddVMStub(params)
	} else {
		return fake.addVMReturns.result1, fake.addVMReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) AddVMCallCount() int {
	fake.addVMMutex.RLock()
	defer fake.addVMMutex.RUnlock()
	return len(fake.addVMArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) AddVMArgsForCall(i int) *vm.AddVMParams {
	fake.addVMMutex.RLock()
	defer fake.addVMMutex.RUnlock()
	return fake.addVMArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) AddVMReturns(result1 *vm.AddVMOK, result2 error) {
	fake.AddVMStub = nil
	fake.addVMReturns = struct {
		result1 *vm.AddVMOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) DeleteVM(params *vm.DeleteVMParams) (*vm.DeleteVMNoContent, error) {
	fake.deleteVMMutex.Lock()
	fake.deleteVMArgsForCall = append(fake.deleteVMArgsForCall, struct {
		params *vm.DeleteVMParams
	}{params})
	fake.recordInvocation("DeleteVM", []interface{}{params})
	fake.deleteVMMutex.Unlock()
	if fake.DeleteVMStub != nil {
		return fake.DeleteVMStub(params)
	} else {
		return fake.deleteVMReturns.result1, fake.deleteVMReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) DeleteVMCallCount() int {
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	return len(fake.deleteVMArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) DeleteVMArgsForCall(i int) *vm.DeleteVMParams {
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	return fake.deleteVMArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) DeleteVMReturns(result1 *vm.DeleteVMNoContent, result2 error) {
	fake.DeleteVMStub = nil
	fake.deleteVMReturns = struct {
		result1 *vm.DeleteVMNoContent
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) FindVmsByDeployment(params *vm.FindVmsByDeploymentParams) (*vm.FindVmsByDeploymentOK, error) {
	fake.findVmsByDeploymentMutex.Lock()
	fake.findVmsByDeploymentArgsForCall = append(fake.findVmsByDeploymentArgsForCall, struct {
		params *vm.FindVmsByDeploymentParams
	}{params})
	fake.recordInvocation("FindVmsByDeployment", []interface{}{params})
	fake.findVmsByDeploymentMutex.Unlock()
	if fake.FindVmsByDeploymentStub != nil {
		return fake.FindVmsByDeploymentStub(params)
	} else {
		return fake.findVmsByDeploymentReturns.result1, fake.findVmsByDeploymentReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) FindVmsByDeploymentCallCount() int {
	fake.findVmsByDeploymentMutex.RLock()
	defer fake.findVmsByDeploymentMutex.RUnlock()
	return len(fake.findVmsByDeploymentArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) FindVmsByDeploymentArgsForCall(i int) *vm.FindVmsByDeploymentParams {
	fake.findVmsByDeploymentMutex.RLock()
	defer fake.findVmsByDeploymentMutex.RUnlock()
	return fake.findVmsByDeploymentArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) FindVmsByDeploymentReturns(result1 *vm.FindVmsByDeploymentOK, result2 error) {
	fake.FindVmsByDeploymentStub = nil
	fake.findVmsByDeploymentReturns = struct {
		result1 *vm.FindVmsByDeploymentOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) FindVmsByFilters(params *vm.FindVmsByFiltersParams) (*vm.FindVmsByFiltersOK, error) {
	fake.findVmsByFiltersMutex.Lock()
	fake.findVmsByFiltersArgsForCall = append(fake.findVmsByFiltersArgsForCall, struct {
		params *vm.FindVmsByFiltersParams
	}{params})
	fake.recordInvocation("FindVmsByFilters", []interface{}{params})
	fake.findVmsByFiltersMutex.Unlock()
	if fake.FindVmsByFiltersStub != nil {
		return fake.FindVmsByFiltersStub(params)
	} else {
		return fake.findVmsByFiltersReturns.result1, fake.findVmsByFiltersReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) FindVmsByFiltersCallCount() int {
	fake.findVmsByFiltersMutex.RLock()
	defer fake.findVmsByFiltersMutex.RUnlock()
	return len(fake.findVmsByFiltersArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) FindVmsByFiltersArgsForCall(i int) *vm.FindVmsByFiltersParams {
	fake.findVmsByFiltersMutex.RLock()
	defer fake.findVmsByFiltersMutex.RUnlock()
	return fake.findVmsByFiltersArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) FindVmsByFiltersReturns(result1 *vm.FindVmsByFiltersOK, result2 error) {
	fake.FindVmsByFiltersStub = nil
	fake.findVmsByFiltersReturns = struct {
		result1 *vm.FindVmsByFiltersOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) FindVmsByStates(params *vm.FindVmsByStatesParams) (*vm.FindVmsByStatesOK, error) {
	fake.findVmsByStatesMutex.Lock()
	fake.findVmsByStatesArgsForCall = append(fake.findVmsByStatesArgsForCall, struct {
		params *vm.FindVmsByStatesParams
	}{params})
	fake.recordInvocation("FindVmsByStates", []interface{}{params})
	fake.findVmsByStatesMutex.Unlock()
	if fake.FindVmsByStatesStub != nil {
		return fake.FindVmsByStatesStub(params)
	} else {
		return fake.findVmsByStatesReturns.result1, fake.findVmsByStatesReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) FindVmsByStatesCallCount() int {
	fake.findVmsByStatesMutex.RLock()
	defer fake.findVmsByStatesMutex.RUnlock()
	return len(fake.findVmsByStatesArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) FindVmsByStatesArgsForCall(i int) *vm.FindVmsByStatesParams {
	fake.findVmsByStatesMutex.RLock()
	defer fake.findVmsByStatesMutex.RUnlock()
	return fake.findVmsByStatesArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) FindVmsByStatesReturns(result1 *vm.FindVmsByStatesOK, result2 error) {
	fake.FindVmsByStatesStub = nil
	fake.findVmsByStatesReturns = struct {
		result1 *vm.FindVmsByStatesOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) GetVMByCid(params *vm.GetVMByCidParams) (*vm.GetVMByCidOK, error) {
	fake.getVMByCidMutex.Lock()
	fake.getVMByCidArgsForCall = append(fake.getVMByCidArgsForCall, struct {
		params *vm.GetVMByCidParams
	}{params})
	fake.recordInvocation("GetVMByCid", []interface{}{params})
	fake.getVMByCidMutex.Unlock()
	if fake.GetVMByCidStub != nil {
		return fake.GetVMByCidStub(params)
	} else {
		return fake.getVMByCidReturns.result1, fake.getVMByCidReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) GetVMByCidCallCount() int {
	fake.getVMByCidMutex.RLock()
	defer fake.getVMByCidMutex.RUnlock()
	return len(fake.getVMByCidArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) GetVMByCidArgsForCall(i int) *vm.GetVMByCidParams {
	fake.getVMByCidMutex.RLock()
	defer fake.getVMByCidMutex.RUnlock()
	return fake.getVMByCidArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) GetVMByCidReturns(result1 *vm.GetVMByCidOK, result2 error) {
	fake.GetVMByCidStub = nil
	fake.getVMByCidReturns = struct {
		result1 *vm.GetVMByCidOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) ListVM(params *vm.ListVMParams) (*vm.ListVMOK, error) {
	fake.listVMMutex.Lock()
	fake.listVMArgsForCall = append(fake.listVMArgsForCall, struct {
		params *vm.ListVMParams
	}{params})
	fake.recordInvocation("ListVM", []interface{}{params})
	fake.listVMMutex.Unlock()
	if fake.ListVMStub != nil {
		return fake.ListVMStub(params)
	} else {
		return fake.listVMReturns.result1, fake.listVMReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) ListVMCallCount() int {
	fake.listVMMutex.RLock()
	defer fake.listVMMutex.RUnlock()
	return len(fake.listVMArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) ListVMArgsForCall(i int) *vm.ListVMParams {
	fake.listVMMutex.RLock()
	defer fake.listVMMutex.RUnlock()
	return fake.listVMArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) ListVMReturns(result1 *vm.ListVMOK, result2 error) {
	fake.ListVMStub = nil
	fake.listVMReturns = struct {
		result1 *vm.ListVMOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) OrderVMByFilter(params *vm.OrderVMByFilterParams) (*vm.OrderVMByFilterOK, error) {
	fake.orderVMByFilterMutex.Lock()
	fake.orderVMByFilterArgsForCall = append(fake.orderVMByFilterArgsForCall, struct {
		params *vm.OrderVMByFilterParams
	}{params})
	fake.recordInvocation("OrderVMByFilter", []interface{}{params})
	fake.orderVMByFilterMutex.Unlock()
	if fake.OrderVMByFilterStub != nil {
		return fake.OrderVMByFilterStub(params)
	} else {
		return fake.orderVMByFilterReturns.result1, fake.orderVMByFilterReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) OrderVMByFilterCallCount() int {
	fake.orderVMByFilterMutex.RLock()
	defer fake.orderVMByFilterMutex.RUnlock()
	return len(fake.orderVMByFilterArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) OrderVMByFilterArgsForCall(i int) *vm.OrderVMByFilterParams {
	fake.orderVMByFilterMutex.RLock()
	defer fake.orderVMByFilterMutex.RUnlock()
	return fake.orderVMByFilterArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) OrderVMByFilterReturns(result1 *vm.OrderVMByFilterOK, result2 error) {
	fake.OrderVMByFilterStub = nil
	fake.orderVMByFilterReturns = struct {
		result1 *vm.OrderVMByFilterOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) UpdateVM(params *vm.UpdateVMParams) (*vm.UpdateVMOK, error) {
	fake.updateVMMutex.Lock()
	fake.updateVMArgsForCall = append(fake.updateVMArgsForCall, struct {
		params *vm.UpdateVMParams
	}{params})
	fake.recordInvocation("UpdateVM", []interface{}{params})
	fake.updateVMMutex.Unlock()
	if fake.UpdateVMStub != nil {
		return fake.UpdateVMStub(params)
	} else {
		return fake.updateVMReturns.result1, fake.updateVMReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) UpdateVMCallCount() int {
	fake.updateVMMutex.RLock()
	defer fake.updateVMMutex.RUnlock()
	return len(fake.updateVMArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) UpdateVMArgsForCall(i int) *vm.UpdateVMParams {
	fake.updateVMMutex.RLock()
	defer fake.updateVMMutex.RUnlock()
	return fake.updateVMArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) UpdateVMReturns(result1 *vm.UpdateVMOK, result2 error) {
	fake.UpdateVMStub = nil
	fake.updateVMReturns = struct {
		result1 *vm.UpdateVMOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) UpdateVMWithState(params *vm.UpdateVMWithStateParams) (*vm.UpdateVMWithStateOK, error) {
	fake.updateVMWithStateMutex.Lock()
	fake.updateVMWithStateArgsForCall = append(fake.updateVMWithStateArgsForCall, struct {
		params *vm.UpdateVMWithStateParams
	}{params})
	fake.recordInvocation("UpdateVMWithState", []interface{}{params})
	fake.updateVMWithStateMutex.Unlock()
	if fake.UpdateVMWithStateStub != nil {
		return fake.UpdateVMWithStateStub(params)
	} else {
		return fake.updateVMWithStateReturns.result1, fake.updateVMWithStateReturns.result2
	}
}

func (fake *FakeSoftLayerPoolClient) UpdateVMWithStateCallCount() int {
	fake.updateVMWithStateMutex.RLock()
	defer fake.updateVMWithStateMutex.RUnlock()
	return len(fake.updateVMWithStateArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) UpdateVMWithStateArgsForCall(i int) *vm.UpdateVMWithStateParams {
	fake.updateVMWithStateMutex.RLock()
	defer fake.updateVMWithStateMutex.RUnlock()
	return fake.updateVMWithStateArgsForCall[i].params
}

func (fake *FakeSoftLayerPoolClient) UpdateVMWithStateReturns(result1 *vm.UpdateVMWithStateOK, result2 error) {
	fake.UpdateVMWithStateStub = nil
	fake.updateVMWithStateReturns = struct {
		result1 *vm.UpdateVMWithStateOK
		result2 error
	}{result1, result2}
}

func (fake *FakeSoftLayerPoolClient) SetTransport(transport runtime.ClientTransport) {
	fake.setTransportMutex.Lock()
	fake.setTransportArgsForCall = append(fake.setTransportArgsForCall, struct {
		transport runtime.ClientTransport
	}{transport})
	fake.recordInvocation("SetTransport", []interface{}{transport})
	fake.setTransportMutex.Unlock()
	if fake.SetTransportStub != nil {
		fake.SetTransportStub(transport)
	}
}

func (fake *FakeSoftLayerPoolClient) SetTransportCallCount() int {
	fake.setTransportMutex.RLock()
	defer fake.setTransportMutex.RUnlock()
	return len(fake.setTransportArgsForCall)
}

func (fake *FakeSoftLayerPoolClient) SetTransportArgsForCall(i int) runtime.ClientTransport {
	fake.setTransportMutex.RLock()
	defer fake.setTransportMutex.RUnlock()
	return fake.setTransportArgsForCall[i].transport
}

func (fake *FakeSoftLayerPoolClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addVMMutex.RLock()
	defer fake.addVMMutex.RUnlock()
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	fake.findVmsByDeploymentMutex.RLock()
	defer fake.findVmsByDeploymentMutex.RUnlock()
	fake.findVmsByFiltersMutex.RLock()
	defer fake.findVmsByFiltersMutex.RUnlock()
	fake.findVmsByStatesMutex.RLock()
	defer fake.findVmsByStatesMutex.RUnlock()
	fake.getVMByCidMutex.RLock()
	defer fake.getVMByCidMutex.RUnlock()
	fake.listVMMutex.RLock()
	defer fake.listVMMutex.RUnlock()
	fake.orderVMByFilterMutex.RLock()
	defer fake.orderVMByFilterMutex.RUnlock()
	fake.updateVMMutex.RLock()
	defer fake.updateVMMutex.RUnlock()
	fake.updateVMWithStateMutex.RLock()
	defer fake.updateVMWithStateMutex.RUnlock()
	fake.setTransportMutex.RLock()
	defer fake.setTransportMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSoftLayerPoolClient) recordInvocation(key string, args []interface{}) {
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

var _ vm.SoftLayerPoolClient = new(FakeSoftLayerPoolClient)

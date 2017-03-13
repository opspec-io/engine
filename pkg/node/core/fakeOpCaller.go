// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/pkg/model"
)

type fakeOpCaller struct {
	CallStub        func(inputScope map[string]*model.Data, opId string, pkgRef string, rootOpId string) (err error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		inputScope map[string]*model.Data
		opId       string
		pkgRef     string
		rootOpId   string
	}
	callReturns struct {
		result1 error
	}
	callReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeOpCaller) Call(inputScope map[string]*model.Data, opId string, pkgRef string, rootOpId string) (err error) {
	fake.callMutex.Lock()
	ret, specificReturn := fake.callReturnsOnCall[len(fake.callArgsForCall)]
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		inputScope map[string]*model.Data
		opId       string
		pkgRef     string
		rootOpId   string
	}{inputScope, opId, pkgRef, rootOpId})
	fake.recordInvocation("Call", []interface{}{inputScope, opId, pkgRef, rootOpId})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(inputScope, opId, pkgRef, rootOpId)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.callReturns.result1
}

func (fake *fakeOpCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeOpCaller) CallArgsForCall(i int) (map[string]*model.Data, string, string, string) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].inputScope, fake.callArgsForCall[i].opId, fake.callArgsForCall[i].pkgRef, fake.callArgsForCall[i].rootOpId
}

func (fake *fakeOpCaller) CallReturns(result1 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 error
	}{result1}
}

func (fake *fakeOpCaller) CallReturnsOnCall(i int, result1 error) {
	fake.CallStub = nil
	if fake.callReturnsOnCall == nil {
		fake.callReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.callReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *fakeOpCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeOpCaller) recordInvocation(key string, args []interface{}) {
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

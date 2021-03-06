// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/internal/iruntime"
)

type FakeIRuntime struct {
	GOOSStub        func() string
	gOOSMutex       sync.RWMutex
	gOOSArgsForCall []struct {
	}
	gOOSReturns struct {
		result1 string
	}
	gOOSReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIRuntime) GOOS() string {
	fake.gOOSMutex.Lock()
	ret, specificReturn := fake.gOOSReturnsOnCall[len(fake.gOOSArgsForCall)]
	fake.gOOSArgsForCall = append(fake.gOOSArgsForCall, struct {
	}{})
	fake.recordInvocation("GOOS", []interface{}{})
	fake.gOOSMutex.Unlock()
	if fake.GOOSStub != nil {
		return fake.GOOSStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gOOSReturns
	return fakeReturns.result1
}

func (fake *FakeIRuntime) GOOSCallCount() int {
	fake.gOOSMutex.RLock()
	defer fake.gOOSMutex.RUnlock()
	return len(fake.gOOSArgsForCall)
}

func (fake *FakeIRuntime) GOOSCalls(stub func() string) {
	fake.gOOSMutex.Lock()
	defer fake.gOOSMutex.Unlock()
	fake.GOOSStub = stub
}

func (fake *FakeIRuntime) GOOSReturns(result1 string) {
	fake.gOOSMutex.Lock()
	defer fake.gOOSMutex.Unlock()
	fake.GOOSStub = nil
	fake.gOOSReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIRuntime) GOOSReturnsOnCall(i int, result1 string) {
	fake.gOOSMutex.Lock()
	defer fake.gOOSMutex.Unlock()
	fake.GOOSStub = nil
	if fake.gOOSReturnsOnCall == nil {
		fake.gOOSReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.gOOSReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeIRuntime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gOOSMutex.RLock()
	defer fake.gOOSMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIRuntime) recordInvocation(key string, args []interface{}) {
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

var _ iruntime.IRuntime = new(FakeIRuntime)

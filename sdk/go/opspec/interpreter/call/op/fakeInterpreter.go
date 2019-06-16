// Code generated by counterfeiter. DO NOT EDIT.
package op

import (
	"sync"

	"github.com/opctl/opctl/sdk/go/model"
)

type FakeInterpreter struct {
	InterpretStub        func(scope map[string]*model.Value, scgOpCall *model.SCGOpCall, opID string, parentOpHandle model.DataHandle, rootOpID string) (*model.DCGOpCall, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		scope          map[string]*model.Value
		scgOpCall      *model.SCGOpCall
		opID           string
		parentOpHandle model.DataHandle
		rootOpID       string
	}
	interpretReturns struct {
		result1 *model.DCGOpCall
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 *model.DCGOpCall
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(scope map[string]*model.Value, scgOpCall *model.SCGOpCall, opID string, parentOpHandle model.DataHandle, rootOpID string) (*model.DCGOpCall, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		scope          map[string]*model.Value
		scgOpCall      *model.SCGOpCall
		opID           string
		parentOpHandle model.DataHandle
		rootOpID       string
	}{scope, scgOpCall, opID, parentOpHandle, rootOpID})
	fake.recordInvocation("Interpret", []interface{}{scope, scgOpCall, opID, parentOpHandle, rootOpID})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(scope, scgOpCall, opID, parentOpHandle, rootOpID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *FakeInterpreter) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *FakeInterpreter) InterpretArgsForCall(i int) (map[string]*model.Value, *model.SCGOpCall, string, model.DataHandle, string) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].scope, fake.interpretArgsForCall[i].scgOpCall, fake.interpretArgsForCall[i].opID, fake.interpretArgsForCall[i].parentOpHandle, fake.interpretArgsForCall[i].rootOpID
}

func (fake *FakeInterpreter) InterpretReturns(result1 *model.DCGOpCall, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 *model.DCGOpCall
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 *model.DCGOpCall, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 *model.DCGOpCall
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 *model.DCGOpCall
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterpreter) recordInvocation(key string, args []interface{}) {
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

var _ Interpreter = new(FakeInterpreter)

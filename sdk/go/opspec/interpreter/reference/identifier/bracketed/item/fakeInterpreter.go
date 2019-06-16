// Code generated by counterfeiter. DO NOT EDIT.
package item

import (
	"sync"

	"github.com/opctl/opctl/sdk/go/model"
)

type FakeInterpreter struct {
	InterpretStub        func(identifier string, value model.Value) (*model.Value, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		identifier string
		value      model.Value
	}
	interpretReturns struct {
		result1 *model.Value
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(identifier string, value model.Value) (*model.Value, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		identifier string
		value      model.Value
	}{identifier, value})
	fake.recordInvocation("Interpret", []interface{}{identifier, value})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(identifier, value)
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

func (fake *FakeInterpreter) InterpretArgsForCall(i int) (string, model.Value) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].identifier, fake.interpretArgsForCall[i].value
}

func (fake *FakeInterpreter) InterpretReturns(result1 *model.Value, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 *model.Value
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

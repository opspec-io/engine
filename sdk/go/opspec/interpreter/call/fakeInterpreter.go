// Code generated by counterfeiter. DO NOT EDIT.
package call

import (
	"sync"

	"github.com/opctl/opctl/sdk/go/model"
)

type FakeInterpreter struct {
	InterpretStub        func(map[string]*model.Value, *model.SCG, string, model.DataHandle, *string, string) (*model.DCG, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		arg1 map[string]*model.Value
		arg2 *model.SCG
		arg3 string
		arg4 model.DataHandle
		arg5 *string
		arg6 string
	}
	interpretReturns struct {
		result1 *model.DCG
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 *model.DCG
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(arg1 map[string]*model.Value, arg2 *model.SCG, arg3 string, arg4 model.DataHandle, arg5 *string, arg6 string) (*model.DCG, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		arg1 map[string]*model.Value
		arg2 *model.SCG
		arg3 string
		arg4 model.DataHandle
		arg5 *string
		arg6 string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Interpret", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.interpretReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterpreter) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *FakeInterpreter) InterpretCalls(stub func(map[string]*model.Value, *model.SCG, string, model.DataHandle, *string, string) (*model.DCG, error)) {
	fake.interpretMutex.Lock()
	defer fake.interpretMutex.Unlock()
	fake.InterpretStub = stub
}

func (fake *FakeInterpreter) InterpretArgsForCall(i int) (map[string]*model.Value, *model.SCG, string, model.DataHandle, *string, string) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	argsForCall := fake.interpretArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeInterpreter) InterpretReturns(result1 *model.DCG, result2 error) {
	fake.interpretMutex.Lock()
	defer fake.interpretMutex.Unlock()
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 *model.DCG
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 *model.DCG, result2 error) {
	fake.interpretMutex.Lock()
	defer fake.interpretMutex.Unlock()
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 *model.DCG
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 *model.DCG
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

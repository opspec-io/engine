// Code generated by counterfeiter. DO NOT EDIT.
package bracketed

import (
	"sync"

	"github.com/opctl/opctl/sdk/go/model"
)

type fakeCoerceToArrayOrObjecter struct {
	CoerceToArrayOrObjectStub        func(data *model.Value) (*model.Value, error)
	coerceToArrayOrObjectMutex       sync.RWMutex
	coerceToArrayOrObjectArgsForCall []struct {
		data *model.Value
	}
	coerceToArrayOrObjectReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToArrayOrObjectReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCoerceToArrayOrObjecter) CoerceToArrayOrObject(data *model.Value) (*model.Value, error) {
	fake.coerceToArrayOrObjectMutex.Lock()
	ret, specificReturn := fake.coerceToArrayOrObjectReturnsOnCall[len(fake.coerceToArrayOrObjectArgsForCall)]
	fake.coerceToArrayOrObjectArgsForCall = append(fake.coerceToArrayOrObjectArgsForCall, struct {
		data *model.Value
	}{data})
	fake.recordInvocation("CoerceToArrayOrObject", []interface{}{data})
	fake.coerceToArrayOrObjectMutex.Unlock()
	if fake.CoerceToArrayOrObjectStub != nil {
		return fake.CoerceToArrayOrObjectStub(data)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToArrayOrObjectReturns.result1, fake.coerceToArrayOrObjectReturns.result2
}

func (fake *fakeCoerceToArrayOrObjecter) CoerceToArrayOrObjectCallCount() int {
	fake.coerceToArrayOrObjectMutex.RLock()
	defer fake.coerceToArrayOrObjectMutex.RUnlock()
	return len(fake.coerceToArrayOrObjectArgsForCall)
}

func (fake *fakeCoerceToArrayOrObjecter) CoerceToArrayOrObjectArgsForCall(i int) *model.Value {
	fake.coerceToArrayOrObjectMutex.RLock()
	defer fake.coerceToArrayOrObjectMutex.RUnlock()
	return fake.coerceToArrayOrObjectArgsForCall[i].data
}

func (fake *fakeCoerceToArrayOrObjecter) CoerceToArrayOrObjectReturns(result1 *model.Value, result2 error) {
	fake.CoerceToArrayOrObjectStub = nil
	fake.coerceToArrayOrObjectReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoerceToArrayOrObjecter) CoerceToArrayOrObjectReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToArrayOrObjectStub = nil
	if fake.coerceToArrayOrObjectReturnsOnCall == nil {
		fake.coerceToArrayOrObjectReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToArrayOrObjectReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoerceToArrayOrObjecter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.coerceToArrayOrObjectMutex.RLock()
	defer fake.coerceToArrayOrObjectMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeCoerceToArrayOrObjecter) recordInvocation(key string, args []interface{}) {
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

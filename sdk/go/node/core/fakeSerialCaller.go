// Code generated by counterfeiter. DO NOT EDIT.
package core

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdk/go/model"
)

type fakeSerialCaller struct {
	CallStub        func(context.Context, string, map[string]*model.Value, string, model.DataHandle, []*model.SCG)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]*model.Value
		arg4 string
		arg5 model.DataHandle
		arg6 []*model.SCG
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeSerialCaller) Call(arg1 context.Context, arg2 string, arg3 map[string]*model.Value, arg4 string, arg5 model.DataHandle, arg6 []*model.SCG) {
	var arg6Copy []*model.SCG
	if arg6 != nil {
		arg6Copy = make([]*model.SCG, len(arg6))
		copy(arg6Copy, arg6)
	}
	fake.callMutex.Lock()
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]*model.Value
		arg4 string
		arg5 model.DataHandle
		arg6 []*model.SCG
	}{arg1, arg2, arg3, arg4, arg5, arg6Copy})
	fake.recordInvocation("Call", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6Copy})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		fake.CallStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func (fake *fakeSerialCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeSerialCaller) CallCalls(stub func(context.Context, string, map[string]*model.Value, string, model.DataHandle, []*model.SCG)) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = stub
}

func (fake *fakeSerialCaller) CallArgsForCall(i int) (context.Context, string, map[string]*model.Value, string, model.DataHandle, []*model.SCG) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	argsForCall := fake.callArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *fakeSerialCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeSerialCaller) recordInvocation(key string, args []interface{}) {
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

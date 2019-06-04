// Code generated by counterfeiter. DO NOT EDIT.
package core

import (
	"context"
	"sync"
)

type Fake struct {
	EventsStub        func(context.Context)
	eventsMutex       sync.RWMutex
	eventsArgsForCall []struct {
		arg1 context.Context
	}
	LsStub        func(context.Context, string)
	lsMutex       sync.RWMutex
	lsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	NodeCreateStub        func(NodeCreateOpts)
	nodeCreateMutex       sync.RWMutex
	nodeCreateArgsForCall []struct {
		arg1 NodeCreateOpts
	}
	NodeKillStub        func()
	nodeKillMutex       sync.RWMutex
	nodeKillArgsForCall []struct {
	}
	OpCreateStub        func(string, string, string)
	opCreateMutex       sync.RWMutex
	opCreateArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	OpInstallStub        func(context.Context, string, string, string, string)
	opInstallMutex       sync.RWMutex
	opInstallArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	OpKillStub        func(context.Context, string)
	opKillMutex       sync.RWMutex
	opKillArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	OpValidateStub        func(context.Context, string)
	opValidateMutex       sync.RWMutex
	opValidateArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	RunStub        func(context.Context, string, *RunOpts)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 *RunOpts
	}
	SelfUpdateStub        func(string)
	selfUpdateMutex       sync.RWMutex
	selfUpdateArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Events(arg1 context.Context) {
	fake.eventsMutex.Lock()
	fake.eventsArgsForCall = append(fake.eventsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Events", []interface{}{arg1})
	fake.eventsMutex.Unlock()
	if fake.EventsStub != nil {
		fake.EventsStub(arg1)
	}
}

func (fake *Fake) EventsCallCount() int {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	return len(fake.eventsArgsForCall)
}

func (fake *Fake) EventsCalls(stub func(context.Context)) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = stub
}

func (fake *Fake) EventsArgsForCall(i int) context.Context {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	argsForCall := fake.eventsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Fake) Ls(arg1 context.Context, arg2 string) {
	fake.lsMutex.Lock()
	fake.lsArgsForCall = append(fake.lsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Ls", []interface{}{arg1, arg2})
	fake.lsMutex.Unlock()
	if fake.LsStub != nil {
		fake.LsStub(arg1, arg2)
	}
}

func (fake *Fake) LsCallCount() int {
	fake.lsMutex.RLock()
	defer fake.lsMutex.RUnlock()
	return len(fake.lsArgsForCall)
}

func (fake *Fake) LsCalls(stub func(context.Context, string)) {
	fake.lsMutex.Lock()
	defer fake.lsMutex.Unlock()
	fake.LsStub = stub
}

func (fake *Fake) LsArgsForCall(i int) (context.Context, string) {
	fake.lsMutex.RLock()
	defer fake.lsMutex.RUnlock()
	argsForCall := fake.lsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Fake) NodeCreate(arg1 NodeCreateOpts) {
	fake.nodeCreateMutex.Lock()
	fake.nodeCreateArgsForCall = append(fake.nodeCreateArgsForCall, struct {
		arg1 NodeCreateOpts
	}{arg1})
	fake.recordInvocation("NodeCreate", []interface{}{arg1})
	fake.nodeCreateMutex.Unlock()
	if fake.NodeCreateStub != nil {
		fake.NodeCreateStub(arg1)
	}
}

func (fake *Fake) NodeCreateCallCount() int {
	fake.nodeCreateMutex.RLock()
	defer fake.nodeCreateMutex.RUnlock()
	return len(fake.nodeCreateArgsForCall)
}

func (fake *Fake) NodeCreateCalls(stub func(NodeCreateOpts)) {
	fake.nodeCreateMutex.Lock()
	defer fake.nodeCreateMutex.Unlock()
	fake.NodeCreateStub = stub
}

func (fake *Fake) NodeCreateArgsForCall(i int) NodeCreateOpts {
	fake.nodeCreateMutex.RLock()
	defer fake.nodeCreateMutex.RUnlock()
	argsForCall := fake.nodeCreateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Fake) NodeKill() {
	fake.nodeKillMutex.Lock()
	fake.nodeKillArgsForCall = append(fake.nodeKillArgsForCall, struct {
	}{})
	fake.recordInvocation("NodeKill", []interface{}{})
	fake.nodeKillMutex.Unlock()
	if fake.NodeKillStub != nil {
		fake.NodeKillStub()
	}
}

func (fake *Fake) NodeKillCallCount() int {
	fake.nodeKillMutex.RLock()
	defer fake.nodeKillMutex.RUnlock()
	return len(fake.nodeKillArgsForCall)
}

func (fake *Fake) NodeKillCalls(stub func()) {
	fake.nodeKillMutex.Lock()
	defer fake.nodeKillMutex.Unlock()
	fake.NodeKillStub = stub
}

func (fake *Fake) OpCreate(arg1 string, arg2 string, arg3 string) {
	fake.opCreateMutex.Lock()
	fake.opCreateArgsForCall = append(fake.opCreateArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("OpCreate", []interface{}{arg1, arg2, arg3})
	fake.opCreateMutex.Unlock()
	if fake.OpCreateStub != nil {
		fake.OpCreateStub(arg1, arg2, arg3)
	}
}

func (fake *Fake) OpCreateCallCount() int {
	fake.opCreateMutex.RLock()
	defer fake.opCreateMutex.RUnlock()
	return len(fake.opCreateArgsForCall)
}

func (fake *Fake) OpCreateCalls(stub func(string, string, string)) {
	fake.opCreateMutex.Lock()
	defer fake.opCreateMutex.Unlock()
	fake.OpCreateStub = stub
}

func (fake *Fake) OpCreateArgsForCall(i int) (string, string, string) {
	fake.opCreateMutex.RLock()
	defer fake.opCreateMutex.RUnlock()
	argsForCall := fake.opCreateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Fake) OpInstall(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string) {
	fake.opInstallMutex.Lock()
	fake.opInstallArgsForCall = append(fake.opInstallArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("OpInstall", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.opInstallMutex.Unlock()
	if fake.OpInstallStub != nil {
		fake.OpInstallStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *Fake) OpInstallCallCount() int {
	fake.opInstallMutex.RLock()
	defer fake.opInstallMutex.RUnlock()
	return len(fake.opInstallArgsForCall)
}

func (fake *Fake) OpInstallCalls(stub func(context.Context, string, string, string, string)) {
	fake.opInstallMutex.Lock()
	defer fake.opInstallMutex.Unlock()
	fake.OpInstallStub = stub
}

func (fake *Fake) OpInstallArgsForCall(i int) (context.Context, string, string, string, string) {
	fake.opInstallMutex.RLock()
	defer fake.opInstallMutex.RUnlock()
	argsForCall := fake.opInstallArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *Fake) OpKill(arg1 context.Context, arg2 string) {
	fake.opKillMutex.Lock()
	fake.opKillArgsForCall = append(fake.opKillArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("OpKill", []interface{}{arg1, arg2})
	fake.opKillMutex.Unlock()
	if fake.OpKillStub != nil {
		fake.OpKillStub(arg1, arg2)
	}
}

func (fake *Fake) OpKillCallCount() int {
	fake.opKillMutex.RLock()
	defer fake.opKillMutex.RUnlock()
	return len(fake.opKillArgsForCall)
}

func (fake *Fake) OpKillCalls(stub func(context.Context, string)) {
	fake.opKillMutex.Lock()
	defer fake.opKillMutex.Unlock()
	fake.OpKillStub = stub
}

func (fake *Fake) OpKillArgsForCall(i int) (context.Context, string) {
	fake.opKillMutex.RLock()
	defer fake.opKillMutex.RUnlock()
	argsForCall := fake.opKillArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Fake) OpValidate(arg1 context.Context, arg2 string) {
	fake.opValidateMutex.Lock()
	fake.opValidateArgsForCall = append(fake.opValidateArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("OpValidate", []interface{}{arg1, arg2})
	fake.opValidateMutex.Unlock()
	if fake.OpValidateStub != nil {
		fake.OpValidateStub(arg1, arg2)
	}
}

func (fake *Fake) OpValidateCallCount() int {
	fake.opValidateMutex.RLock()
	defer fake.opValidateMutex.RUnlock()
	return len(fake.opValidateArgsForCall)
}

func (fake *Fake) OpValidateCalls(stub func(context.Context, string)) {
	fake.opValidateMutex.Lock()
	defer fake.opValidateMutex.Unlock()
	fake.OpValidateStub = stub
}

func (fake *Fake) OpValidateArgsForCall(i int) (context.Context, string) {
	fake.opValidateMutex.RLock()
	defer fake.opValidateMutex.RUnlock()
	argsForCall := fake.opValidateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Fake) Run(arg1 context.Context, arg2 string, arg3 *RunOpts) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 *RunOpts
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub(arg1, arg2, arg3)
	}
}

func (fake *Fake) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *Fake) RunCalls(stub func(context.Context, string, *RunOpts)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *Fake) RunArgsForCall(i int) (context.Context, string, *RunOpts) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Fake) SelfUpdate(arg1 string) {
	fake.selfUpdateMutex.Lock()
	fake.selfUpdateArgsForCall = append(fake.selfUpdateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SelfUpdate", []interface{}{arg1})
	fake.selfUpdateMutex.Unlock()
	if fake.SelfUpdateStub != nil {
		fake.SelfUpdateStub(arg1)
	}
}

func (fake *Fake) SelfUpdateCallCount() int {
	fake.selfUpdateMutex.RLock()
	defer fake.selfUpdateMutex.RUnlock()
	return len(fake.selfUpdateArgsForCall)
}

func (fake *Fake) SelfUpdateCalls(stub func(string)) {
	fake.selfUpdateMutex.Lock()
	defer fake.selfUpdateMutex.Unlock()
	fake.SelfUpdateStub = stub
}

func (fake *Fake) SelfUpdateArgsForCall(i int) string {
	fake.selfUpdateMutex.RLock()
	defer fake.selfUpdateMutex.RUnlock()
	argsForCall := fake.selfUpdateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	fake.lsMutex.RLock()
	defer fake.lsMutex.RUnlock()
	fake.nodeCreateMutex.RLock()
	defer fake.nodeCreateMutex.RUnlock()
	fake.nodeKillMutex.RLock()
	defer fake.nodeKillMutex.RUnlock()
	fake.opCreateMutex.RLock()
	defer fake.opCreateMutex.RUnlock()
	fake.opInstallMutex.RLock()
	defer fake.opInstallMutex.RUnlock()
	fake.opKillMutex.RLock()
	defer fake.opKillMutex.RUnlock()
	fake.opValidateMutex.RLock()
	defer fake.opValidateMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.selfUpdateMutex.RLock()
	defer fake.selfUpdateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ Core = new(Fake)

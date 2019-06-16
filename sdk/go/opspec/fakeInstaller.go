// Code generated by counterfeiter. DO NOT EDIT.
package op

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdk/go/model"
)

type FakeInstaller struct {
	InstallStub        func(ctx context.Context, path string, opHandle model.DataHandle) error
	installMutex       sync.RWMutex
	installArgsForCall []struct {
		ctx      context.Context
		path     string
		opHandle model.DataHandle
	}
	installReturns struct {
		result1 error
	}
	installReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstaller) Install(ctx context.Context, path string, opHandle model.DataHandle) error {
	fake.installMutex.Lock()
	ret, specificReturn := fake.installReturnsOnCall[len(fake.installArgsForCall)]
	fake.installArgsForCall = append(fake.installArgsForCall, struct {
		ctx      context.Context
		path     string
		opHandle model.DataHandle
	}{ctx, path, opHandle})
	fake.recordInvocation("Install", []interface{}{ctx, path, opHandle})
	fake.installMutex.Unlock()
	if fake.InstallStub != nil {
		return fake.InstallStub(ctx, path, opHandle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.installReturns.result1
}

func (fake *FakeInstaller) InstallCallCount() int {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return len(fake.installArgsForCall)
}

func (fake *FakeInstaller) InstallArgsForCall(i int) (context.Context, string, model.DataHandle) {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return fake.installArgsForCall[i].ctx, fake.installArgsForCall[i].path, fake.installArgsForCall[i].opHandle
}

func (fake *FakeInstaller) InstallReturns(result1 error) {
	fake.InstallStub = nil
	fake.installReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstaller) InstallReturnsOnCall(i int, result1 error) {
	fake.InstallStub = nil
	if fake.installReturnsOnCall == nil {
		fake.installReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.installReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstaller) recordInvocation(key string, args []interface{}) {
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

var _ Installer = new(FakeInstaller)

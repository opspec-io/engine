// This file was generated by counterfeiter
package core

import (
	"sync"
)

type FakeCore struct {
	CreateCollectionStub        func(description string, name string)
	createCollectionMutex       sync.RWMutex
	createCollectionArgsForCall []struct {
		description string
		name        string
	}
	CreateOpStub        func(collection string, description string, name string)
	createOpMutex       sync.RWMutex
	createOpArgsForCall []struct {
		collection  string
		description string
		name        string
	}
	KillOpStub        func(opId string)
	killOpMutex       sync.RWMutex
	killOpArgsForCall []struct {
		opId string
	}
	ListOpsInCollectionStub        func(collection string)
	listOpsInCollectionMutex       sync.RWMutex
	listOpsInCollectionArgsForCall []struct {
		collection string
	}
	RunOpStub        func(args []string, collection string, name string)
	runOpMutex       sync.RWMutex
	runOpArgsForCall []struct {
		args       []string
		collection string
		name       string
	}
	SetCollectionDescriptionStub        func(description string)
	setCollectionDescriptionMutex       sync.RWMutex
	setCollectionDescriptionArgsForCall []struct {
		description string
	}
	SetOpDescriptionStub        func(collection string, description string, name string)
	setOpDescriptionMutex       sync.RWMutex
	setOpDescriptionArgsForCall []struct {
		collection  string
		description string
		name        string
	}
	StreamEventsStub        func()
	streamEventsMutex       sync.RWMutex
	streamEventsArgsForCall []struct{}
	SelfUpdateStub          func(channel string)
	selfUpdateMutex         sync.RWMutex
	selfUpdateArgsForCall   []struct {
		channel string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCore) CreateCollection(description string, name string) {
	fake.createCollectionMutex.Lock()
	fake.createCollectionArgsForCall = append(fake.createCollectionArgsForCall, struct {
		description string
		name        string
	}{description, name})
	fake.recordInvocation("CreateCollection", []interface{}{description, name})
	fake.createCollectionMutex.Unlock()
	if fake.CreateCollectionStub != nil {
		fake.CreateCollectionStub(description, name)
	}
}

func (fake *FakeCore) CreateCollectionCallCount() int {
	fake.createCollectionMutex.RLock()
	defer fake.createCollectionMutex.RUnlock()
	return len(fake.createCollectionArgsForCall)
}

func (fake *FakeCore) CreateCollectionArgsForCall(i int) (string, string) {
	fake.createCollectionMutex.RLock()
	defer fake.createCollectionMutex.RUnlock()
	return fake.createCollectionArgsForCall[i].description, fake.createCollectionArgsForCall[i].name
}

func (fake *FakeCore) CreateOp(collection string, description string, name string) {
	fake.createOpMutex.Lock()
	fake.createOpArgsForCall = append(fake.createOpArgsForCall, struct {
		collection  string
		description string
		name        string
	}{collection, description, name})
	fake.recordInvocation("CreateOp", []interface{}{collection, description, name})
	fake.createOpMutex.Unlock()
	if fake.CreateOpStub != nil {
		fake.CreateOpStub(collection, description, name)
	}
}

func (fake *FakeCore) CreateOpCallCount() int {
	fake.createOpMutex.RLock()
	defer fake.createOpMutex.RUnlock()
	return len(fake.createOpArgsForCall)
}

func (fake *FakeCore) CreateOpArgsForCall(i int) (string, string, string) {
	fake.createOpMutex.RLock()
	defer fake.createOpMutex.RUnlock()
	return fake.createOpArgsForCall[i].collection, fake.createOpArgsForCall[i].description, fake.createOpArgsForCall[i].name
}

func (fake *FakeCore) KillOp(opId string) {
	fake.killOpMutex.Lock()
	fake.killOpArgsForCall = append(fake.killOpArgsForCall, struct {
		opId string
	}{opId})
	fake.recordInvocation("KillOp", []interface{}{opId})
	fake.killOpMutex.Unlock()
	if fake.KillOpStub != nil {
		fake.KillOpStub(opId)
	}
}

func (fake *FakeCore) KillOpCallCount() int {
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	return len(fake.killOpArgsForCall)
}

func (fake *FakeCore) KillOpArgsForCall(i int) string {
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	return fake.killOpArgsForCall[i].opId
}

func (fake *FakeCore) ListOpsInCollection(collection string) {
	fake.listOpsInCollectionMutex.Lock()
	fake.listOpsInCollectionArgsForCall = append(fake.listOpsInCollectionArgsForCall, struct {
		collection string
	}{collection})
	fake.recordInvocation("ListOpsInCollection", []interface{}{collection})
	fake.listOpsInCollectionMutex.Unlock()
	if fake.ListOpsInCollectionStub != nil {
		fake.ListOpsInCollectionStub(collection)
	}
}

func (fake *FakeCore) ListOpsInCollectionCallCount() int {
	fake.listOpsInCollectionMutex.RLock()
	defer fake.listOpsInCollectionMutex.RUnlock()
	return len(fake.listOpsInCollectionArgsForCall)
}

func (fake *FakeCore) ListOpsInCollectionArgsForCall(i int) string {
	fake.listOpsInCollectionMutex.RLock()
	defer fake.listOpsInCollectionMutex.RUnlock()
	return fake.listOpsInCollectionArgsForCall[i].collection
}

func (fake *FakeCore) RunOp(args []string, collection string, name string) {
	var argsCopy []string
	if args != nil {
		argsCopy = make([]string, len(args))
		copy(argsCopy, args)
	}
	fake.runOpMutex.Lock()
	fake.runOpArgsForCall = append(fake.runOpArgsForCall, struct {
		args       []string
		collection string
		name       string
	}{argsCopy, collection, name})
	fake.recordInvocation("RunOp", []interface{}{argsCopy, collection, name})
	fake.runOpMutex.Unlock()
	if fake.RunOpStub != nil {
		fake.RunOpStub(args, collection, name)
	}
}

func (fake *FakeCore) RunOpCallCount() int {
	fake.runOpMutex.RLock()
	defer fake.runOpMutex.RUnlock()
	return len(fake.runOpArgsForCall)
}

func (fake *FakeCore) RunOpArgsForCall(i int) ([]string, string, string) {
	fake.runOpMutex.RLock()
	defer fake.runOpMutex.RUnlock()
	return fake.runOpArgsForCall[i].args, fake.runOpArgsForCall[i].collection, fake.runOpArgsForCall[i].name
}

func (fake *FakeCore) SetCollectionDescription(description string) {
	fake.setCollectionDescriptionMutex.Lock()
	fake.setCollectionDescriptionArgsForCall = append(fake.setCollectionDescriptionArgsForCall, struct {
		description string
	}{description})
	fake.recordInvocation("SetCollectionDescription", []interface{}{description})
	fake.setCollectionDescriptionMutex.Unlock()
	if fake.SetCollectionDescriptionStub != nil {
		fake.SetCollectionDescriptionStub(description)
	}
}

func (fake *FakeCore) SetCollectionDescriptionCallCount() int {
	fake.setCollectionDescriptionMutex.RLock()
	defer fake.setCollectionDescriptionMutex.RUnlock()
	return len(fake.setCollectionDescriptionArgsForCall)
}

func (fake *FakeCore) SetCollectionDescriptionArgsForCall(i int) string {
	fake.setCollectionDescriptionMutex.RLock()
	defer fake.setCollectionDescriptionMutex.RUnlock()
	return fake.setCollectionDescriptionArgsForCall[i].description
}

func (fake *FakeCore) SetOpDescription(collection string, description string, name string) {
	fake.setOpDescriptionMutex.Lock()
	fake.setOpDescriptionArgsForCall = append(fake.setOpDescriptionArgsForCall, struct {
		collection  string
		description string
		name        string
	}{collection, description, name})
	fake.recordInvocation("SetOpDescription", []interface{}{collection, description, name})
	fake.setOpDescriptionMutex.Unlock()
	if fake.SetOpDescriptionStub != nil {
		fake.SetOpDescriptionStub(collection, description, name)
	}
}

func (fake *FakeCore) SetOpDescriptionCallCount() int {
	fake.setOpDescriptionMutex.RLock()
	defer fake.setOpDescriptionMutex.RUnlock()
	return len(fake.setOpDescriptionArgsForCall)
}

func (fake *FakeCore) SetOpDescriptionArgsForCall(i int) (string, string, string) {
	fake.setOpDescriptionMutex.RLock()
	defer fake.setOpDescriptionMutex.RUnlock()
	return fake.setOpDescriptionArgsForCall[i].collection, fake.setOpDescriptionArgsForCall[i].description, fake.setOpDescriptionArgsForCall[i].name
}

func (fake *FakeCore) StreamEvents() {
	fake.streamEventsMutex.Lock()
	fake.streamEventsArgsForCall = append(fake.streamEventsArgsForCall, struct{}{})
	fake.recordInvocation("StreamEvents", []interface{}{})
	fake.streamEventsMutex.Unlock()
	if fake.StreamEventsStub != nil {
		fake.StreamEventsStub()
	}
}

func (fake *FakeCore) StreamEventsCallCount() int {
	fake.streamEventsMutex.RLock()
	defer fake.streamEventsMutex.RUnlock()
	return len(fake.streamEventsArgsForCall)
}

func (fake *FakeCore) SelfUpdate(channel string) {
	fake.selfUpdateMutex.Lock()
	fake.selfUpdateArgsForCall = append(fake.selfUpdateArgsForCall, struct {
		channel string
	}{channel})
	fake.recordInvocation("SelfUpdate", []interface{}{channel})
	fake.selfUpdateMutex.Unlock()
	if fake.SelfUpdateStub != nil {
		fake.SelfUpdateStub(channel)
	}
}

func (fake *FakeCore) SelfUpdateCallCount() int {
	fake.selfUpdateMutex.RLock()
	defer fake.selfUpdateMutex.RUnlock()
	return len(fake.selfUpdateArgsForCall)
}

func (fake *FakeCore) SelfUpdateArgsForCall(i int) string {
	fake.selfUpdateMutex.RLock()
	defer fake.selfUpdateMutex.RUnlock()
	return fake.selfUpdateArgsForCall[i].channel
}

func (fake *FakeCore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCollectionMutex.RLock()
	defer fake.createCollectionMutex.RUnlock()
	fake.createOpMutex.RLock()
	defer fake.createOpMutex.RUnlock()
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	fake.listOpsInCollectionMutex.RLock()
	defer fake.listOpsInCollectionMutex.RUnlock()
	fake.runOpMutex.RLock()
	defer fake.runOpMutex.RUnlock()
	fake.setCollectionDescriptionMutex.RLock()
	defer fake.setCollectionDescriptionMutex.RUnlock()
	fake.setOpDescriptionMutex.RLock()
	defer fake.setOpDescriptionMutex.RUnlock()
	fake.streamEventsMutex.RLock()
	defer fake.streamEventsMutex.RUnlock()
	fake.selfUpdateMutex.RLock()
	defer fake.selfUpdateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCore) recordInvocation(key string, args []interface{}) {
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

var _ = new(FakeCore)
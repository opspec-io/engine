// This file was generated by counterfeiter
package core

import (
  "sync"

  "github.com/opctl/engine/core/models"
)

type FakeApi struct {
  GetEventStreamStub        func(eventChannel chan models.Event) (err error)
  getEventStreamMutex       sync.RWMutex
  getEventStreamArgsForCall []struct {
    eventChannel chan models.Event
  }
  getEventStreamReturns     struct {
                              result1 error
                            }
  KillOpRunStub             func(req models.KillOpRunReq) (correlationId string, err error)
  killOpRunMutex            sync.RWMutex
  killOpRunArgsForCall      []struct {
    req models.KillOpRunReq
  }
  killOpRunReturns          struct {
                              result1 string
                              result2 error
                            }
  RunOpStub                 func(req models.RunOpReq) (opRunId string, correlationId string, err error)
  runOpMutex                sync.RWMutex
  runOpArgsForCall          []struct {
    req models.RunOpReq
  }
  runOpReturns              struct {
                              result1 string
                              result2 string
                              result3 error
                            }
  invocations               map[string][][]interface{}
  invocationsMutex          sync.RWMutex
}

func (fake *FakeApi) GetEventStream(eventChannel chan models.Event) (err error) {
  fake.getEventStreamMutex.Lock()
  fake.getEventStreamArgsForCall = append(fake.getEventStreamArgsForCall, struct {
    eventChannel chan models.Event
  }{eventChannel})
  fake.recordInvocation("GetEventStream", []interface{}{eventChannel})
  fake.getEventStreamMutex.Unlock()
  if fake.GetEventStreamStub != nil {
    return fake.GetEventStreamStub(eventChannel)
  } else {
    return fake.getEventStreamReturns.result1
  }
}

func (fake *FakeApi) GetEventStreamCallCount() int {
  fake.getEventStreamMutex.RLock()
  defer fake.getEventStreamMutex.RUnlock()
  return len(fake.getEventStreamArgsForCall)
}

func (fake *FakeApi) GetEventStreamArgsForCall(i int) chan models.Event {
  fake.getEventStreamMutex.RLock()
  defer fake.getEventStreamMutex.RUnlock()
  return fake.getEventStreamArgsForCall[i].eventChannel
}

func (fake *FakeApi) GetEventStreamReturns(result1 error) {
  fake.GetEventStreamStub = nil
  fake.getEventStreamReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) KillOpRun(req models.KillOpRunReq) (correlationId string, err error) {
  fake.killOpRunMutex.Lock()
  fake.killOpRunArgsForCall = append(fake.killOpRunArgsForCall, struct {
    req models.KillOpRunReq
  }{req})
  fake.recordInvocation("KillOpRun", []interface{}{req})
  fake.killOpRunMutex.Unlock()
  if fake.KillOpRunStub != nil {
    return fake.KillOpRunStub(req)
  } else {
    return fake.killOpRunReturns.result1, fake.killOpRunReturns.result2
  }
}

func (fake *FakeApi) KillOpRunCallCount() int {
  fake.killOpRunMutex.RLock()
  defer fake.killOpRunMutex.RUnlock()
  return len(fake.killOpRunArgsForCall)
}

func (fake *FakeApi) KillOpRunArgsForCall(i int) models.KillOpRunReq {
  fake.killOpRunMutex.RLock()
  defer fake.killOpRunMutex.RUnlock()
  return fake.killOpRunArgsForCall[i].req
}

func (fake *FakeApi) KillOpRunReturns(result1 string, result2 error) {
  fake.KillOpRunStub = nil
  fake.killOpRunReturns = struct {
    result1 string
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) RunOp(req models.RunOpReq) (opRunId string, correlationId string, err error) {
  fake.runOpMutex.Lock()
  fake.runOpArgsForCall = append(fake.runOpArgsForCall, struct {
    req models.RunOpReq
  }{req})
  fake.recordInvocation("RunOp", []interface{}{req})
  fake.runOpMutex.Unlock()
  if fake.RunOpStub != nil {
    return fake.RunOpStub(req)
  } else {
    return fake.runOpReturns.result1, fake.runOpReturns.result2, fake.runOpReturns.result3
  }
}

func (fake *FakeApi) RunOpCallCount() int {
  fake.runOpMutex.RLock()
  defer fake.runOpMutex.RUnlock()
  return len(fake.runOpArgsForCall)
}

func (fake *FakeApi) RunOpArgsForCall(i int) models.RunOpReq {
  fake.runOpMutex.RLock()
  defer fake.runOpMutex.RUnlock()
  return fake.runOpArgsForCall[i].req
}

func (fake *FakeApi) RunOpReturns(result1 string, result2 string, result3 error) {
  fake.RunOpStub = nil
  fake.runOpReturns = struct {
    result1 string
    result2 string
    result3 error
  }{result1, result2, result3}
}

func (fake *FakeApi) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.getEventStreamMutex.RLock()
  defer fake.getEventStreamMutex.RUnlock()
  fake.killOpRunMutex.RLock()
  defer fake.killOpRunMutex.RUnlock()
  fake.runOpMutex.RLock()
  defer fake.runOpMutex.RUnlock()
  return fake.invocations
}

func (fake *FakeApi) recordInvocation(key string, args []interface{}) {
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

var _ Api = new(FakeApi)

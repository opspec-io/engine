// This file was generated by counterfeiter
package core

import (
  "sync"

  "github.com/dev-op-spec/engine/core/models"
)

type FakeApi struct {
  AddOpStub                     func(req models.AddOpReq) (err error)
  addOpMutex                    sync.RWMutex
  addOpArgsForCall              []struct {
    req models.AddOpReq
  }
  addOpReturns                  struct {
                                  result1 error
                                }
  AddSubOpStub                  func(req models.AddSubOpReq) (err error)
  addSubOpMutex                 sync.RWMutex
  addSubOpArgsForCall           []struct {
    req models.AddSubOpReq
  }
  addSubOpReturns               struct {
                                  result1 error
                                }
  GetEventStreamStub            func(eventChannel chan models.Event) (err error)
  getEventStreamMutex           sync.RWMutex
  getEventStreamArgsForCall     []struct {
    eventChannel chan models.Event
  }
  getEventStreamReturns         struct {
                                  result1 error
                                }
  GetLogForOpRunStub            func(opRunId string, logChannel chan *models.LogEntry) (err error)
  getLogForOpRunMutex           sync.RWMutex
  getLogForOpRunArgsForCall     []struct {
    opRunId    string
    logChannel chan *models.LogEntry
  }
  getLogForOpRunReturns         struct {
                                  result1 error
                                }
  ListOpsStub                   func(projectUrl *models.Url) (ops []models.OpDetailedView, err error)
  listOpsMutex                  sync.RWMutex
  listOpsArgsForCall            []struct {
    projectUrl *models.Url
  }
  listOpsReturns                struct {
                                  result1 []models.OpDetailedView
                                  result2 error
                                }
  RunOpStub                     func(req models.RunOpReq) (opRunId string, err error)
  runOpMutex                    sync.RWMutex
  runOpArgsForCall              []struct {
    req models.RunOpReq
  }
  runOpReturns                  struct {
                                  result1 string
                                  result2 error
                                }
  SetDescriptionOfOpStub        func(req models.SetDescriptionOfOpReq) (err error)
  setDescriptionOfOpMutex       sync.RWMutex
  setDescriptionOfOpArgsForCall []struct {
    req models.SetDescriptionOfOpReq
  }
  setDescriptionOfOpReturns     struct {
                                  result1 error
                                }
}

func (fake *FakeApi) AddOp(req models.AddOpReq) (err error) {
  fake.addOpMutex.Lock()
  fake.addOpArgsForCall = append(fake.addOpArgsForCall, struct {
    req models.AddOpReq
  }{req})
  fake.addOpMutex.Unlock()
  if fake.AddOpStub != nil {
    return fake.AddOpStub(req)
  } else {
    return fake.addOpReturns.result1
  }
}

func (fake *FakeApi) AddOpCallCount() int {
  fake.addOpMutex.RLock()
  defer fake.addOpMutex.RUnlock()
  return len(fake.addOpArgsForCall)
}

func (fake *FakeApi) AddOpArgsForCall(i int) models.AddOpReq {
  fake.addOpMutex.RLock()
  defer fake.addOpMutex.RUnlock()
  return fake.addOpArgsForCall[i].req
}

func (fake *FakeApi) AddOpReturns(result1 error) {
  fake.AddOpStub = nil
  fake.addOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) AddSubOp(req models.AddSubOpReq) (err error) {
  fake.addSubOpMutex.Lock()
  fake.addSubOpArgsForCall = append(fake.addSubOpArgsForCall, struct {
    req models.AddSubOpReq
  }{req})
  fake.addSubOpMutex.Unlock()
  if fake.AddSubOpStub != nil {
    return fake.AddSubOpStub(req)
  } else {
    return fake.addSubOpReturns.result1
  }
}

func (fake *FakeApi) AddSubOpCallCount() int {
  fake.addSubOpMutex.RLock()
  defer fake.addSubOpMutex.RUnlock()
  return len(fake.addSubOpArgsForCall)
}

func (fake *FakeApi) AddSubOpArgsForCall(i int) models.AddSubOpReq {
  fake.addSubOpMutex.RLock()
  defer fake.addSubOpMutex.RUnlock()
  return fake.addSubOpArgsForCall[i].req
}

func (fake *FakeApi) AddSubOpReturns(result1 error) {
  fake.AddSubOpStub = nil
  fake.addSubOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) GetEventStream(eventChannel chan models.Event) (err error) {
  fake.getEventStreamMutex.Lock()
  fake.getEventStreamArgsForCall = append(fake.getEventStreamArgsForCall, struct {
    eventChannel chan models.Event
  }{eventChannel})
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

func (fake *FakeApi) GetLogForOpRun(opRunId string, logChannel chan *models.LogEntry) (err error) {
  fake.getLogForOpRunMutex.Lock()
  fake.getLogForOpRunArgsForCall = append(fake.getLogForOpRunArgsForCall, struct {
    opRunId    string
    logChannel chan *models.LogEntry
  }{opRunId, logChannel})
  fake.getLogForOpRunMutex.Unlock()
  if fake.GetLogForOpRunStub != nil {
    return fake.GetLogForOpRunStub(opRunId, logChannel)
  } else {
    return fake.getLogForOpRunReturns.result1
  }
}

func (fake *FakeApi) GetLogForOpRunCallCount() int {
  fake.getLogForOpRunMutex.RLock()
  defer fake.getLogForOpRunMutex.RUnlock()
  return len(fake.getLogForOpRunArgsForCall)
}

func (fake *FakeApi) GetLogForOpRunArgsForCall(i int) (string, chan *models.LogEntry) {
  fake.getLogForOpRunMutex.RLock()
  defer fake.getLogForOpRunMutex.RUnlock()
  return fake.getLogForOpRunArgsForCall[i].opRunId, fake.getLogForOpRunArgsForCall[i].logChannel
}

func (fake *FakeApi) GetLogForOpRunReturns(result1 error) {
  fake.GetLogForOpRunStub = nil
  fake.getLogForOpRunReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) ListOps(projectUrl *models.Url) (ops []models.OpDetailedView, err error) {
  fake.listOpsMutex.Lock()
  fake.listOpsArgsForCall = append(fake.listOpsArgsForCall, struct {
    projectUrl *models.Url
  }{projectUrl})
  fake.listOpsMutex.Unlock()
  if fake.ListOpsStub != nil {
    return fake.ListOpsStub(projectUrl)
  } else {
    return fake.listOpsReturns.result1, fake.listOpsReturns.result2
  }
}

func (fake *FakeApi) ListOpsCallCount() int {
  fake.listOpsMutex.RLock()
  defer fake.listOpsMutex.RUnlock()
  return len(fake.listOpsArgsForCall)
}

func (fake *FakeApi) ListOpsArgsForCall(i int) *models.Url {
  fake.listOpsMutex.RLock()
  defer fake.listOpsMutex.RUnlock()
  return fake.listOpsArgsForCall[i].projectUrl
}

func (fake *FakeApi) ListOpsReturns(result1 []models.OpDetailedView, result2 error) {
  fake.ListOpsStub = nil
  fake.listOpsReturns = struct {
    result1 []models.OpDetailedView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) RunOp(req models.RunOpReq) (opRunId string, err error) {
  fake.runOpMutex.Lock()
  fake.runOpArgsForCall = append(fake.runOpArgsForCall, struct {
    req models.RunOpReq
  }{req})
  fake.runOpMutex.Unlock()
  if fake.RunOpStub != nil {
    return fake.RunOpStub(req)
  } else {
    return fake.runOpReturns.result1, fake.runOpReturns.result2
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

func (fake *FakeApi) RunOpReturns(result1 string, result2 error) {
  fake.RunOpStub = nil
  fake.runOpReturns = struct {
    result1 string
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) SetDescriptionOfOp(req models.SetDescriptionOfOpReq) (err error) {
  fake.setDescriptionOfOpMutex.Lock()
  fake.setDescriptionOfOpArgsForCall = append(fake.setDescriptionOfOpArgsForCall, struct {
    req models.SetDescriptionOfOpReq
  }{req})
  fake.setDescriptionOfOpMutex.Unlock()
  if fake.SetDescriptionOfOpStub != nil {
    return fake.SetDescriptionOfOpStub(req)
  } else {
    return fake.setDescriptionOfOpReturns.result1
  }
}

func (fake *FakeApi) SetDescriptionOfOpCallCount() int {
  fake.setDescriptionOfOpMutex.RLock()
  defer fake.setDescriptionOfOpMutex.RUnlock()
  return len(fake.setDescriptionOfOpArgsForCall)
}

func (fake *FakeApi) SetDescriptionOfOpArgsForCall(i int) models.SetDescriptionOfOpReq {
  fake.setDescriptionOfOpMutex.RLock()
  defer fake.setDescriptionOfOpMutex.RUnlock()
  return fake.setDescriptionOfOpArgsForCall[i].req
}

func (fake *FakeApi) SetDescriptionOfOpReturns(result1 error) {
  fake.SetDescriptionOfOpStub = nil
  fake.setDescriptionOfOpReturns = struct {
    result1 error
  }{result1}
}

var _ Api = new(FakeApi)

package core

//go:generate counterfeiter -o ./fakeOpRunner.go --fake-name fakeOpRunner ./ opRunner

import (
  "github.com/opctl/engine/core/models"
  "github.com/opctl/engine/core/ports"
  "github.com/opctl/engine/core/logging"
  "github.com/opspec-io/sdk-golang"
  "path/filepath"
  "path"
  "time"
  "sync"
)

type opRunner interface {
  Run(
  correlationId string,
  opArgs map[string]string,
  opNamespace string,
  opUrl string,
  parentOpRunId string,
  ) (
  opRunId string,
  err error,
  )

  Kill(
  correlationId string,
  opRunId string,
  ) (
  err error,
  )
}

func newOpRunner(
containerEngine ports.ContainerEngine,
eventStream eventStream,
logger logging.Logger,
opspecSdk opspec.Sdk,
uniqueStringFactory uniqueStringFactory,
) opRunner {

  return &_opRunner{
    containerEngine: containerEngine,
    eventStream:eventStream,
    logger:logger,
    opspecSdk:opspecSdk,
    uniqueStringFactory:uniqueStringFactory,
    unfinishedOpRunsByIdMap:make(map[string]models.OpRunStartedEvent),
  }

}

type _opRunner struct {
  containerEngine                ports.ContainerEngine
  eventStream                    eventStream
  logger                         logging.Logger
  opspecSdk                      opspec.Sdk
  uniqueStringFactory            uniqueStringFactory

  unfinishedOpRunsByIdMapRWMutex sync.RWMutex
  unfinishedOpRunsByIdMap        map[string]models.OpRunStartedEvent
}

func (this _opRunner) Run(
correlationId string,
opArgs map[string]string,
opNamespace string,
opBundleUrl string,
parentOpRunId string,
) (
opRunId string,
err error,
) {

  _opFile, err := this.opspecSdk.GetOp(
    opBundleUrl,
  )
  if (nil != err) {
    return
  }

  opRunId, err = this.uniqueStringFactory.Construct()
  if (nil != err) {
    return
  }

  var rootOpRunId string
  if ("" == parentOpRunId) {
    // handle root op run

    rootOpRunId = opRunId

  } else {
    // handle sub op run

    this.unfinishedOpRunsByIdMapRWMutex.RLock()
    rootOpRunId = this.unfinishedOpRunsByIdMap[parentOpRunId].RootOpRunId()
    this.unfinishedOpRunsByIdMapRWMutex.RUnlock()

  }

  opRunStartedEvent := models.NewOpRunStartedEvent(
    correlationId,
    opBundleUrl,
    opRunId,
    parentOpRunId,
    rootOpRunId,
    time.Now().UTC(),
  )

  this.eventStream.Publish(opRunStartedEvent)

  this.unfinishedOpRunsByIdMapRWMutex.Lock()
  this.unfinishedOpRunsByIdMap[opRunId] = opRunStartedEvent
  this.unfinishedOpRunsByIdMapRWMutex.Unlock()

  go func() {

    var opRunExitCode int
    defer func() {

      this.unfinishedOpRunsByIdMapRWMutex.RLock()
      _, isUnfinishedOpRun := this.unfinishedOpRunsByIdMap[opRunId]
      this.unfinishedOpRunsByIdMapRWMutex.RUnlock()

      if (isUnfinishedOpRun) {

        this.unfinishedOpRunsByIdMapRWMutex.Lock()
        delete(this.unfinishedOpRunsByIdMap, opRunId)
        this.unfinishedOpRunsByIdMapRWMutex.Unlock()

        this.eventStream.Publish(
          models.NewOpRunFinishedEvent(
            correlationId,
            opRunExitCode,
            opRunId,
            rootOpRunId,
            time.Now().UTC(),
          ),
        )

      }

    }()

    if (len(_opFile.SubOps) == 0) {
      // run op

      opRunExitCode, err = this.containerEngine.RunOp(
        correlationId,
        opArgs,
        opBundleUrl,
        _opFile.Name,
        opNamespace,
        this.logger,
      )
      if (nil != err) {
        this.logger(
          models.NewLogEntryEmittedEvent(
            correlationId,
            time.Now().UTC(),
            err.Error(),
            logging.StdErrStream,
          ),
        )
      }

    } else {
      // run sub ops

      for _, subOp := range _opFile.SubOps {

        subOpUrl := path.Join(
          filepath.Dir(opBundleUrl),
          subOp.Url,
        )

        eventChannel := make(chan models.Event)
        this.eventStream.RegisterSubscriber(eventChannel)
        defer this.eventStream.UnregisterSubscriber(eventChannel)

        var subOpRunId string
        subOpRunId, err = this.Run(
          correlationId,
          opArgs,
          opNamespace,
          subOpUrl,
          opRunId,
        )
        if (nil != err) {
          this.logger(
            models.NewLogEntryEmittedEvent(
              correlationId,
              time.Now().UTC(),
              err.Error(),
              logging.StdErrStream,
            ),
          )
        }

        eventLoop:for {
          event := <-eventChannel

          switch event := event.(type) {
          case models.OpRunFinishedEvent:
            if (event.OpRunId() == subOpRunId) {

              this.eventStream.UnregisterSubscriber(eventChannel)

              if (event.OpRunExitCode() != 0) {
                // if non-zero exit code return immediately

                opRunExitCode = event.OpRunExitCode()
                return

              } else {
                break eventLoop
              }
            }
          default: // no op
          }

        }

      }

    }

  }()

  return

}

func (this _opRunner) Kill(
correlationId string,
opRunId string,
) (
err error,
) {

  // get rootOpRunId
  this.unfinishedOpRunsByIdMapRWMutex.RLock()
  opRun := this.unfinishedOpRunsByIdMap[opRunId]
  this.unfinishedOpRunsByIdMapRWMutex.RUnlock()

  opCollection, err := this.opspecSdk.GetCollection(
    filepath.Dir(opRun.OpRunOpUrl()),
  )
  if (nil != err) {
    return
  }

  this.unfinishedOpRunsByIdMapRWMutex.Lock()
  for _, unfinishedOpRun := range this.unfinishedOpRunsByIdMap {

    // remove all ops with this root
    if (unfinishedOpRun.RootOpRunId() == opRun.RootOpRunId()) {

      delete(this.unfinishedOpRunsByIdMap, unfinishedOpRun.OpRunId())

      this.eventStream.Publish(
        models.NewOpRunKilledEvent(
          correlationId,
          unfinishedOpRun.OpRunId(),
          opRun.RootOpRunId(),
          time.Now(),
        ),
      )

      go func() {

        // @TODO: handle failure scenario; currently this can leak docker-compose resources
        err := this.containerEngine.KillOpRun(
          correlationId,
          unfinishedOpRun.OpRunOpUrl(),
          opCollection.Name,
          this.logger,
        )
        if (nil != err) {
          this.logger(
            models.NewLogEntryEmittedEvent(
              correlationId,
              time.Now().UTC(),
              err.Error(),
              logging.StdErrStream,
            ),
          )
        }

        this.eventStream.Publish(
          models.NewOpRunFinishedEvent(
            correlationId,
            138,
            opRunId,
            opRun.RootOpRunId(),
            time.Now().UTC(),
          ),
        )

      }()
    }

  }
  this.unfinishedOpRunsByIdMapRWMutex.Unlock()

  return
}
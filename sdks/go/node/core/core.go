// Package core defines the core interface for an opspec node
package core

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/opctl/opctl/sdks/go/internal/uniquestring"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/core/containerruntime"
	"github.com/opctl/opctl/sdks/go/pubsub"
)

// New returns a new LocalCore initialized with the given options
func New(
	containerRuntime containerruntime.ContainerRuntime,
	dataDirPath string,
) core {
	eventDbPath := path.Join(dataDirPath, "dcg", "events")
	err := os.MkdirAll(eventDbPath, 0700)
	if nil != err {
		panic(err)
	}

	// per badger README.MD#FAQ "maximizes throughput"
	runtime.GOMAXPROCS(128)

	db, err := badger.Open(
		badger.DefaultOptions(
			eventDbPath,
		).WithLogger(nil),
	)
	if err != nil {
		panic(err)
	}

	pubSub := pubsub.New(db)

	uniqueStringFactory := uniquestring.NewUniqueStringFactory()

	stateStore := newStateStore(
		db,
		pubSub,
	)

	caller := newCaller(
		newContainerCaller(
			containerRuntime,
			pubSub,
			stateStore,
		),
		dataDirPath,
		stateStore,
		pubSub,
	)

	go func() {
		// process events in background
		callKiller := newCallKiller(
			stateStore,
			containerRuntime,
			pubSub,
		)

		ctx := context.Background()

		since := time.Now().UTC()
		eventChannel, _ := pubSub.Subscribe(
			ctx,
			model.EventFilter{
				Since: &since,
			},
		)

		for event := range eventChannel {
			switch {
			case nil != event.CallKillRequested:
				req := event.CallKillRequested.Request
				callKiller.Kill(
					ctx,
					req.OpID,
					req.RootCallID,
				)
			}
		}
	}()

	return core{
		caller:           caller,
		containerRuntime: containerRuntime,
		dataCachePath:    filepath.Join(dataDirPath, "ops"),
		opCaller: newOpCaller(
			stateStore,
			caller,
			dataDirPath,
		),
		pubSub:              pubSub,
		stateStore:          stateStore,
		uniqueStringFactory: uniqueStringFactory,
	}
}

// core is an OpNode that supports running ops directly on the host
type core struct {
	caller              caller
	containerRuntime    containerruntime.ContainerRuntime
	dataCachePath       string
	opCaller            opCaller
	pubSub              pubsub.PubSub
	stateStore          stateStore
	uniqueStringFactory uniquestring.UniqueStringFactory
}

//counterfeiter:generate -o fakes/core.go . Core

// Core is an OpNode that supports running ops directly on the current machine
type Core interface {
	// AddAuth records authentication within the core
	AddAuth(
		ctx context.Context,
		req model.AddAuthReq,
	) error

	GetEventStream(
		ctx context.Context,
		req *model.GetEventStreamReq,
	) (
		<-chan model.Event,
		error,
	)

	// KillOp kills a running op
	KillOp(
		ctx context.Context,
		req model.KillOpReq,
	) (
		err error,
	)

	// StartOp starts an op and returns the root call ID
	StartOp(
		ctx context.Context,
		req model.StartOpReq,
	) (
		rootCallID string,
		err error,
	)

	// GetData gets data
	//
	// expected errs:
	//  - ErrDataProviderAuthentication on authentication failure
	//  - ErrDataProviderAuthorization on authorization failure
	//  - ErrDataRefResolution on resolution failure
	GetData(
		ctx context.Context,
		req model.GetDataReq,
	) (
		model.ReadSeekCloser,
		error,
	)

	// ListDescendants lists file system entries
	//
	// expected errs:
	//  - ErrDataProviderAuthentication on authentication failure
	//  - ErrDataProviderAuthorization on authorization failure
	//  - ErrDataRefResolution on resolution failure
	ListDescendants(
		ctx context.Context,
		req model.ListDescendantsReq,
	) (
		[]*model.DirEntry,
		error,
	)

	// Resolve attempts to resolve data via local filesystem or git
	// nil pullCreds will be ignored
	//
	// expected errs:
	//  - ErrDataProviderAuthentication on authentication failure
	//  - ErrDataProviderAuthorization on authorization failure
	//  - ErrDataRefResolution on resolution failure
	ResolveData(
		ctx context.Context,
		dataRef string,
		pullCreds *model.Creds,
	) (
		model.DataHandle,
		error,
	)
}

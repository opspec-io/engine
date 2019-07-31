// Package containerruntime defines an interface abstracting container runtime interactions.
// A fake implementation is included to allow faking said interactions.
package containerruntime

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fake.go --fake-name Fake ./ ContainerRuntime

import (
	"context"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/util/pubsub"
	"io"
)

// ContainerRuntime defines the interface container runtimes must implement to be supported by
type ContainerRuntime interface {
	DeleteContainerIfExists(
		containerID string,
	) error

	// RunContainer creates, starts, and waits on a container. ExitCode &/Or an error will be returned
	RunContainer(
		ctx context.Context,
		req *model.DCGContainerCall,
		eventPublisher pubsub.EventPublisher,
		stdout io.WriteCloser,
		stderr io.WriteCloser,
	) (*int64, error)
}

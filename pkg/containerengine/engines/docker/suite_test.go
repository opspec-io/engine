package docker

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"
)

func Test(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "engine/pkg/containerengine/engines/docker")
}
package v2

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/worker"
)

type resource struct {
	container      worker.Container
	info           ResourceInfo
	resourceConfig db.ResourceConfig
}

type ResourceInfo struct {
	Artifacts Artifacts
}

type Artifacts struct {
	APIVersion string `json:"api_version"`
	Check      string `json:"check"`
	Get        string `json:"get"`
	Put        string `json:"put"`
}

func NewResource(container worker.Container, info ResourceInfo, resourceConfig db.ResourceConfig) *resource {
	return &resource{
		container:      container,
		info:           info,
		resourceConfig: resourceConfig,
	}
}

func (r *resource) Container() worker.Container {
	return r.container
}

// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/worker"
)

type FakeResourceFactory struct {
	NewResourceStub        func(ctx context.Context, logger lager.Logger, owner db.ContainerOwner, metadata db.ContainerMetadata, containerSpec worker.ContainerSpec, resourceTypes creds.VersionedResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate) (resource.Resource, error)
	newResourceMutex       sync.RWMutex
	newResourceArgsForCall []struct {
		ctx                   context.Context
		logger                lager.Logger
		owner                 db.ContainerOwner
		metadata              db.ContainerMetadata
		containerSpec         worker.ContainerSpec
		resourceTypes         creds.VersionedResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
	}
	newResourceReturns struct {
		result1 resource.Resource
		result2 error
	}
	newResourceReturnsOnCall map[int]struct {
		result1 resource.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) NewResource(ctx context.Context, logger lager.Logger, owner db.ContainerOwner, metadata db.ContainerMetadata, containerSpec worker.ContainerSpec, resourceTypes creds.VersionedResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate) (resource.Resource, error) {
	fake.newResourceMutex.Lock()
	ret, specificReturn := fake.newResourceReturnsOnCall[len(fake.newResourceArgsForCall)]
	fake.newResourceArgsForCall = append(fake.newResourceArgsForCall, struct {
		ctx                   context.Context
		logger                lager.Logger
		owner                 db.ContainerOwner
		metadata              db.ContainerMetadata
		containerSpec         worker.ContainerSpec
		resourceTypes         creds.VersionedResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
	}{ctx, logger, owner, metadata, containerSpec, resourceTypes, imageFetchingDelegate})
	fake.recordInvocation("NewResource", []interface{}{ctx, logger, owner, metadata, containerSpec, resourceTypes, imageFetchingDelegate})
	fake.newResourceMutex.Unlock()
	if fake.NewResourceStub != nil {
		return fake.NewResourceStub(ctx, logger, owner, metadata, containerSpec, resourceTypes, imageFetchingDelegate)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newResourceReturns.result1, fake.newResourceReturns.result2
}

func (fake *FakeResourceFactory) NewResourceCallCount() int {
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	return len(fake.newResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewResourceArgsForCall(i int) (context.Context, lager.Logger, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes, worker.ImageFetchingDelegate) {
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	return fake.newResourceArgsForCall[i].ctx, fake.newResourceArgsForCall[i].logger, fake.newResourceArgsForCall[i].owner, fake.newResourceArgsForCall[i].metadata, fake.newResourceArgsForCall[i].containerSpec, fake.newResourceArgsForCall[i].resourceTypes, fake.newResourceArgsForCall[i].imageFetchingDelegate
}

func (fake *FakeResourceFactory) NewResourceReturns(result1 resource.Resource, result2 error) {
	fake.NewResourceStub = nil
	fake.newResourceReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) NewResourceReturnsOnCall(i int, result1 resource.Resource, result2 error) {
	fake.NewResourceStub = nil
	if fake.newResourceReturnsOnCall == nil {
		fake.newResourceReturnsOnCall = make(map[int]struct {
			result1 resource.Resource
			result2 error
		})
	}
	fake.newResourceReturnsOnCall[i] = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceFactory = new(FakeResourceFactory)
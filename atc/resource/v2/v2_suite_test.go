package v2_test

import (
	"testing"

	res "github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/resource/v2"
	"github.com/concourse/concourse/atc/worker/workerfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	workerClient  *workerfakes.FakeClient
	fakeContainer *workerfakes.FakeContainer

	resourceInfo res.ResourceInfo
	resource     res.Resource
)

var _ = BeforeEach(func() {
	workerClient = new(workerfakes.FakeClient)
	fakeContainer = new(workerfakes.FakeContainer)

	resourceInfo = res.ResourceInfo{
		Artifacts: res.Artifacts{
			APIVersion: "2.0",
			Check:      "artifact check",
			Get:        "artifact get",
			Put:        "artifact put",
		},
	}

	resource = v2.NewResource(fakeContainer, resourceInfo)
})

func TestResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resource V2 Suite")
}

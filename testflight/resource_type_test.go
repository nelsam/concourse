package testflight_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Configuring a resource type in a pipeline config", func() {
	Context("with custom resource types", func() {
		BeforeEach(func() {
			setAndUnpausePipeline("fixtures/resource-types.yml")
		})

		It("can use custom resource types for 'get', 'put', and task 'image_resource's", func() {
			watch := fly("trigger-job", "-j", inPipeline("resource-getter"), "-w")
			Expect(watch).To(gbytes.Say("fetched version: hello-from-custom-type"))

			watch = fly("trigger-job", "-j", inPipeline("resource-putter"), "-w")
			Expect(watch).To(gbytes.Say("pushing version: some-pushed-version"))

			watch = fly("trigger-job", "-j", inPipeline("resource-image-resourcer"), "-w")
			Expect(watch).To(gbytes.Say("MIRRORED_VERSION=image-version"))
		})

		It("can check for resources using a custom type", func() {
			checkResource := fly("check-resource", "-r", inPipeline("my-resource"))
			Expect(checkResource).To(gbytes.Say("checked 'my-resource'"))
		})
	})

	Context("with custom resource types that have params", func() {
		BeforeEach(func() {
			setAndUnpausePipeline("fixtures/resource-types-with-params.yml")
		})

		It("can use a custom resource with parameters", func() {
			watch := fly("trigger-job", "-j", inPipeline("resource-test"), "-w")
			Expect(watch).To(gbytes.Say("mock"))
		})
	})

	Context("when resource type named as base resource type", func() {
		BeforeEach(func() {
			setAndUnpausePipeline("fixtures/resource-type-named-as-base-type.yml")
		})

		It("can use custom resource type named as base resource type", func() {
			watch := fly("trigger-job", "-j", inPipeline("resource-getter"), "-w")
			Expect(watch).To(gbytes.Say("mirror-mirror"))
		})
	})
})

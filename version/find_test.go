package version_test

import (
	version "github.com/elentok/gever/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindInVersionFile", func() {

	It("Parses the contents of a semver file", func() {
		v, err := version.FindInVersionFile("test-fixtures/semver.txt")
		Expect(err).To(BeNil())
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("3.2.1-hotfix3"))
	})

})

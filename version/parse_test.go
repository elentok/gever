package version_test

import (
	"errors"

	version "github.com/elentok/gever/version"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {

	It("Parses 1 as 1.0.0", func() {
		v, _ := version.Parse("1")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.0.0"))
	})

	It("Parses 1.2 as 1.2.0", func() {
		v, _ := version.Parse("1.2")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.2.0"))
	})

	It("Parses 1.2.3", func() {
		v, _ := version.Parse("1.2.3")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.2.3"))
	})

	It("Parses 1.2.3-rc1", func() {
		v, _ := version.Parse("1.2.3-rc1")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.2.3-rc1"))
	})

	It("Parses 1.2.3-rc1\n", func() {
		v, _ := version.Parse("1.2.3-rc1\n")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.2.3-rc1"))
	})

	It("Parses 1-rc1 as 1.0.0-rc1", func() {
		v, _ := version.Parse("1-rc1")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.0.0-rc1"))
	})

	It("Parses 1.2-rc1 as 1.2.0-rc1", func() {
		v, _ := version.Parse("1.2-rc1")
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("1.2.0-rc1"))
	})

	It("Returns an error for 'abc'", func() {
		v, err := version.Parse("abc")
		Expect(v).To(BeNil())
		Expect(err).To(Equal(errors.New("Invalid version 'abc'")))
	})
})

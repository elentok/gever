package gever_test

import (
	. "github.com/elentok/gever"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version creation", func() {

	var (
		v *Version
	)

	BeforeEach(func() {
		v = NewVersion(1, 2, 3, "")
	})

	Describe("#CreateRC", func() {
		Context("when current version is a release", func() {
			It("creates rc1", func() {
				v.CreateRC()
				Expect(v.ToString()).To(Equal("1.2.3-rc1"))
			})
		})

		Context("when current version is a release candidate", func() {
			It("increments the rc count", func() {
				v.SetRC(3)
				v.CreateRC()
				Expect(v.ToString()).To(Equal("1.2.3-rc4"))
			})
		})

		Context("when current version is a hotfix", func() {
			It("creates rc1", func() {
				v.SetHotfix(5)
				v.CreateRC()
				Expect(v.ToString()).To(Equal("1.2.3-rc1"))
			})
		})
	})

	Describe("#CreateHotfix", func() {

		Context("when current version is a release", func() {
			It("creates hotfix1", func() {
				v.CreateHotfix()
				Expect(v.ToString()).To(Equal("1.2.3-hotfix1"))
			})
		})

		Context("when current version is a release candidate", func() {
			It("creates hotfix1", func() {
				v.SetRC(1)
				v.CreateHotfix()
				Expect(v.ToString()).To(Equal("1.2.3-hotfix1"))
			})
		})

		Context("when current version is a hotfix", func() {
			It("increments the hotfix", func() {
				v.SetHotfix(5)
				v.CreateHotfix()
				Expect(v.ToString()).To(Equal("1.2.3-hotfix6"))
			})
		})
	})

	Describe("#CreatePatch", func() {

		It("increments the patch version", func() {
			v.CreatePatch()
			Expect(v.Patch()).To(Equal(4))
		})

		It("resets the special string", func() {
			v.SetSpecial("-bla")
			v.CreatePatch()
			Expect(v.Special()).To(Equal(""))
		})
	})

	Describe("#CreateMinor", func() {

		It("increments the minor version", func() {
			v.CreateMinor()
			Expect(v.Minor()).To(Equal(3))
		})

		It("resets the patch version", func() {
			v.CreateMinor()
			Expect(v.Patch()).To(Equal(0))
		})

		It("resets the special string", func() {
			v.SetSpecial("-bla")
			v.CreateMinor()
			Expect(v.Special()).To(Equal(""))
		})
	})

	Describe("#CreateMajor", func() {

		It("increments the major version", func() {
			v.CreateMajor()
			Expect(v.Major()).To(Equal(2))
		})

		It("resets the minor version", func() {
			v.CreateMajor()
			Expect(v.Minor()).To(Equal(0))
		})

		It("resets the patch version", func() {
			v.CreateMajor()
			Expect(v.Patch()).To(Equal(0))
		})

		It("resets the special string", func() {
			v.SetSpecial("-bla")
			v.CreateMajor()
			Expect(v.Special()).To(Equal(""))
		})
	})

})

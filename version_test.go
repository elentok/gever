package gever_test

import (
	. "github.com/elentok/gever"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version", func() {

	Describe(".NewVersion", func() {
		It("Creates a new version", func() {
			v := NewVersion(1, 2, 3, "-rc1")
			Expect(v.Major()).To(Equal(1))
			Expect(v.Minor()).To(Equal(2))
			Expect(v.Patch()).To(Equal(3))
			Expect(v.Special()).To(Equal("-rc1"))
		})
	})

	Describe("#Hotfix", func() {
		Context("When it's not a hotfix", func() {
			It("Returns 0", func() {
				v := NewVersion(1, 2, 3, "")
				Expect(v.Hotfix()).To(Equal(0))
			})
		})

		Context("When it's a hotfix", func() {
			It("Returns the hotfix index", func() {
				v := NewVersion(1, 2, 3, "-hotfix12")
				Expect(v.Hotfix()).To(Equal(12))
			})
		})
	})

	Describe("#RC", func() {
		Context("When it's not a release candidate", func() {
			It("Returns 0", func() {
				v := NewVersion(1, 2, 3, "")
				Expect(v.RC()).To(Equal(0))
			})
		})

		Context("When it's a release candidate", func() {
			It("Returns the RC index", func() {
				v := NewVersion(1, 2, 3, "-rc11")
				Expect(v.RC()).To(Equal(11))
			})
		})
	})

	Describe("#SetMajor", func() {
		It("Sets the major version", func() {
			v := NewVersion(1, 2, 3, "")
			v.SetMajor(4)
			Expect(v.Major()).To(Equal(4))
		})
	})

	Describe("#SetMinor", func() {
		It("Sets the minor version", func() {
			v := NewVersion(1, 2, 3, "")
			v.SetMinor(4)
			Expect(v.Minor()).To(Equal(4))
		})
	})

	Describe("#SetPatch", func() {
		It("Sets the patch version", func() {
			v := NewVersion(1, 2, 3, "")
			v.SetPatch(4)
			Expect(v.Patch()).To(Equal(4))
		})
	})

	Describe("#SetSpecial", func() {
		It("Sets the special version", func() {
			v := NewVersion(1, 2, 3, "")
			v.SetSpecial("-rc1")
			Expect(v.Special()).To(Equal("-rc1"))
		})
	})

	Describe("#ToString", func() {
		It("Joins the version parts", func() {
			v := NewVersion(1, 2, 3, "-rc1")
			Expect(v.ToString()).To(Equal("1.2.3-rc1"))
		})
	})

	Describe("#SetHotfix", func() {
		It("Sets the special as a hotfix", func() {
			v := NewVersion(1, 2, 3, "")
			v.SetHotfix(4)
			Expect(v.ToString()).To(Equal("1.2.3-hotfix4"))
		})
	})

	Describe("#SetRC", func() {
		It("Sets the special as a release candidate", func() {
			v := NewVersion(1, 2, 3, "")
			v.SetRC(4)
			Expect(v.ToString()).To(Equal("1.2.3-rc4"))
		})
	})
})

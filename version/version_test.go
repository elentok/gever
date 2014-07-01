package version_test

import (
	version "github.com/elentok/gever/version"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version", func() {

	var v version.Version

	BeforeEach(func() {
		v = version.New(1, 2, 3, "")
	})

	Describe(".version.New", func() {
		It("Creates a new version", func() {
			v := version.New(1, 2, 3, "-rc1")
			Expect(v.Major()).To(Equal(1))
			Expect(v.Minor()).To(Equal(2))
			Expect(v.Patch()).To(Equal(3))
			Expect(v.Special()).To(Equal("-rc1"))
		})
	})

	Describe("#Hotfix", func() {
		Context("When it's not a hotfix", func() {
			It("Returns 0", func() {
				v := version.New(1, 2, 3, "")
				Expect(v.Hotfix()).To(Equal(0))
			})
		})

		Context("When it's a hotfix", func() {
			It("Returns the hotfix index", func() {
				v := version.New(1, 2, 3, "-hotfix12")
				Expect(v.Hotfix()).To(Equal(12))
			})
		})
	})

	Describe("#RC", func() {
		Context("When it's not a release candidate", func() {
			It("Returns 0", func() {
				v := version.New(1, 2, 3, "")
				Expect(v.RC()).To(Equal(0))
			})
		})

		Context("When it's a release candidate", func() {
			It("Returns the RC index", func() {
				v := version.New(1, 2, 3, "-rc11")
				Expect(v.RC()).To(Equal(11))
			})
		})
	})

	Describe("#SetMajor", func() {
		It("Sets the major version", func() {
			v := version.New(1, 2, 3, "")
			v.SetMajor(4)
			Expect(v.Major()).To(Equal(4))
		})
	})

	Describe("#SetMinor", func() {
		It("Sets the minor version", func() {
			v := version.New(1, 2, 3, "")
			v.SetMinor(4)
			Expect(v.Minor()).To(Equal(4))
		})
	})

	Describe("#SetPatch", func() {
		It("Sets the patch version", func() {
			v := version.New(1, 2, 3, "")
			v.SetPatch(4)
			Expect(v.Patch()).To(Equal(4))
		})
	})

	Describe("#SetSpecial", func() {
		It("Sets the special version", func() {
			v := version.New(1, 2, 3, "")
			v.SetSpecial("-rc1")
			Expect(v.Special()).To(Equal("-rc1"))
		})
	})

	Describe("#ToString", func() {
		It("Joins the version parts", func() {
			v := version.New(1, 2, 3, "-rc1")
			Expect(v.ToString()).To(Equal("1.2.3-rc1"))
		})
	})

	Describe("#SetHotfix", func() {
		It("Sets the special as a hotfix", func() {
			v := version.New(1, 2, 3, "")
			v.SetHotfix(4)
			Expect(v.ToString()).To(Equal("1.2.3-hotfix4"))
		})
	})

	Describe("#SetRC", func() {
		It("Sets the special as a release candidate", func() {
			v := version.New(1, 2, 3, "")
			v.SetRC(4)
			Expect(v.ToString()).To(Equal("1.2.3-rc4"))
		})
	})

	Describe("#Set", func() {
		It("Sets all of the values", func() {
			v := version.New(0, 0, 0, "")
			v.Set(3, 2, 1, "-rc4")
			Expect(v.ToString()).To(Equal("3.2.1-rc4"))
		})
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

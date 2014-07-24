package version_test

import (
	"io/ioutil"
	"path/filepath"

	"github.com/elentok/gever/git"
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

var _ = Describe("FindInYaml", func() {

	It("Parses the contents of a yaml file", func() {
		v, err := version.FindInYaml("test-fixtures/semver.yml")
		Expect(err).To(BeNil())
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("3.4.5+hotfix3"))
	})

})

var _ = Describe("FindInPackageJSON", func() {

	It("Parses the contents of a package.json file", func() {
		v, err := version.FindInPackageJSON("test-fixtures/package.json")
		Expect(err).To(BeNil())
		Expect(v).NotTo(BeNil())
		Expect(v.ToString()).To(Equal("5.4.3"))
	})

})

var _ = Describe("FindInGitTag", func() {
	var root string

	BeforeEach(func() {
		var err error
		root, err = ioutil.TempDir("", "gever")
		Expect(err).To(BeNil())
		repo := git.NewRepo(root)
		err = repo.Init()
		Expect(err).To(BeNil())
		ioutil.WriteFile(
			filepath.Join(root, "file.txt"),
			[]byte("Testing"),
			0644,
		)

		repo.AddAll()
		repo.Commit("initial commit")
		repo.Tag("v4.5.6", "Bump to v4.5.6")
	})

	It("Finds the last tagged version", func() {
		version, err := version.FindInGitTag(root)
		Expect(err).To(BeNil())
		Expect(version.ToString()).To(Equal("4.5.6"))
	})

})

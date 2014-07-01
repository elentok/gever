package version

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	Major = 0
	Minor
	Patch
)

type Version interface {
	Major() int
	Minor() int
	Patch() int
	Special() string
	Hotfix() int
	RC() int

	ToString() string

	Set(major, minor, patch int, special string)
	SetMajor(int)
	SetMinor(int)
	SetPatch(int)
	SetSpecial(string)
	SetHotfix(int)
	SetRC(int)

	CreateRC()
	CreateHotfix()
	CreatePatch()
	CreateMinor()
	CreateMajor()
}

type version struct {
	major, minor, patch int
	special             string
}

func New(major, minor, patch int, special string) Version {
	return &version{
		major:   major,
		minor:   minor,
		patch:   patch,
		special: special,
	}
}

func (v *version) Major() int {
	return v.major
}

func (v *version) SetMajor(value int) {
	v.major = value
}

func (v *version) Minor() int {
	return v.minor
}

func (v *version) SetMinor(value int) {
	v.minor = value
}

func (v *version) Patch() int {
	return v.patch
}

func (v *version) SetPatch(value int) {
	v.patch = value
}

func (v *version) Special() string {
	return v.special
}

func (v *version) SetSpecial(value string) {
	v.special = value
}

func (v *version) ToString() string {
	return fmt.Sprintf("%d.%d.%d%s", v.major, v.minor, v.patch, v.special)
}

var (
	hotfixRegexp = regexp.MustCompile("^.*hotfix(?P<index>[0-9]+)$")
	rcRegexp     = regexp.MustCompile("^.*rc(?P<index>[0-9]+)$")
)

func (v *version) Hotfix() int {
	return v.findIndex(hotfixRegexp)
}

func (v *version) SetHotfix(index int) {
	v.SetSpecial(fmt.Sprintf("-hotfix%d", index))
}

func (v *version) RC() int {
	return v.findIndex(rcRegexp)
}

func (v *version) SetRC(index int) {
	v.SetSpecial(fmt.Sprintf("-rc%d", index))
}

func (v *version) findIndex(re *regexp.Regexp) int {
	special := v.Special()
	if re.MatchString(special) {
		indexString := re.ReplaceAllString(special, "${index}")
		index, err := strconv.Atoi(indexString)
		if err != nil {
			panic(err)
		}
		return index
	}
	return 0
}

func (v *version) CreateRC() {
	rc := v.RC() + 1
	v.SetRC(rc)
}

func (v *version) CreateHotfix() {
	hotfix := v.Hotfix() + 1
	v.SetHotfix(hotfix)
}

func (v *version) CreatePatch() {
	patch := v.Patch() + 1
	v.SetPatch(patch)
	v.SetSpecial("")
}

func (v *version) CreateMinor() {
	minor := v.Minor() + 1
	v.SetMinor(minor)
	v.SetPatch(0)
	v.SetSpecial("")
}

func (v *version) CreateMajor() {
	major := v.Major() + 1
	v.SetMajor(major)
	v.SetMinor(0)
	v.SetPatch(0)
	v.SetSpecial("")
}

func (v *version) Set(major, minor, patch int, special string) {
	v.SetMajor(major)
	v.SetMinor(minor)
	v.SetPatch(patch)
	v.SetSpecial(special)
}

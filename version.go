package gever

import (
	"fmt"
	"regexp"
	"strconv"
)

type Version struct {
	major, minor, patch int
	special             string
}

func NewVersion(major, minor, patch int, special string) *Version {
	return &Version{
		major:   major,
		minor:   minor,
		patch:   patch,
		special: special,
	}
}

func (v *Version) Major() int {
	return v.major
}

func (v *Version) SetMajor(value int) {
	v.major = value
}

func (v *Version) Minor() int {
	return v.minor
}

func (v *Version) SetMinor(value int) {
	v.minor = value
}

func (v *Version) Patch() int {
	return v.patch
}

func (v *Version) SetPatch(value int) {
	v.patch = value
}

func (v *Version) Special() string {
	return v.special
}

func (v *Version) SetSpecial(value string) {
	v.special = value
}

func (v *Version) ToString() string {
	return fmt.Sprintf("%d.%d.%d%s", v.major, v.minor, v.patch, v.special)
}

var (
	hotfixRegexp = regexp.MustCompile("^.*hotfix(?P<index>[0-9]+)$")
	rcRegexp     = regexp.MustCompile("^.*rc(?P<index>[0-9]+)$")
)

func (v *Version) Hotfix() int {
	return v.findIndex(hotfixRegexp)
}

func (v *Version) SetHotfix(index int) {
	v.SetSpecial(fmt.Sprintf("-hotfix%d", index))
}

func (v *Version) RC() int {
	return v.findIndex(rcRegexp)
}

func (v *Version) SetRC(index int) {
	v.SetSpecial(fmt.Sprintf("-rc%d", index))
}

func (v *Version) findIndex(re *regexp.Regexp) int {
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

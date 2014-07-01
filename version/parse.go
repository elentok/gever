package version

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var versionRegexp = regexp.MustCompile(
	"^(?P<major>[0-9]+)(\\.(?P<minor>[0-9]+)(\\.(?P<patch>[0-9]+))?)?(?P<special>.*)\n?$")

func Parse(s string) (Version, error) {
	if versionRegexp.MatchString(s) {

		majorString := versionRegexp.ReplaceAllString(s, "${major}")
		major, _ := strconv.Atoi(majorString)

		minorString := versionRegexp.ReplaceAllString(s, "${minor}")
		minor, _ := strconv.Atoi(minorString)

		patchString := versionRegexp.ReplaceAllString(s, "${patch}")
		patch, _ := strconv.Atoi(patchString)

		special := versionRegexp.ReplaceAllString(s, "${special}")

		return New(major, minor, patch, special), nil
	}

	return nil, errors.New(fmt.Sprintf("Invalid version '%s'", s))
}

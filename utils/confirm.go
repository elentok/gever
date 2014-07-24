package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var yesRegexp = regexp.MustCompile("^[yY]")

func Confirm(question string, defaultValue bool) bool {
	var defaults = "[Y/n]"
	if !defaultValue {
		defaults = "[y/N]"
	}

	fmt.Printf("%s %s? ", question, defaults)
	var response string
	fmt.Scanln(&response)
	response = strings.TrimRight(response, "\r\n")

	if len(response) == 0 {
		return defaultValue
	}

	return yesRegexp.MatchString(response)
}

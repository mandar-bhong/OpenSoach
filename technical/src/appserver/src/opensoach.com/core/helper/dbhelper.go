package helper

import (
	"regexp"
)

func DBQueryParamValidate(queryInput string, emptyAllowed bool) bool {

	var pattern *regexp.Regexp
	if emptyAllowed {
		pattern = regexp.MustCompile("^[A-Za-z0-9_]*$")
	} else {
		pattern = regexp.MustCompile("^[A-Za-z0-9_]+$")
	}

	return pattern.MatchString(queryInput)
}

package premux

import (
	"strings"
)

const (
	PathRoot              = "/"
	PathDelimiter         = PathRoot
	ParameterDelimiter    = ":"
	PatternDelimiterStart = "["
	PatternDelimiterEnd   = "]"
	PatternWildcard       = "(.+)"
)

// ExpandPath separates a PathDelimiter-delimited string into a slice of strings.
func ExpandPath(path string) []string {
	var r []string

	for _, str := range strings.Split(path, PathDelimiter) {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

// DeriveLabelPattern derives from a given label a regex pattern.
// e.g. :id[^\d+$] => ^\d+$
// e.g. :id => (.+)
func DeriveLabelPattern(label string) string {
	start := strings.Index(label, PatternDelimiterStart)
	end := strings.Index(label, PatternDelimiterEnd)

	// If the label doesn't contain a pattern, default to the wildcard pattern.
	if start == -1 || end == -1 {
		return PatternWildcard
	}

	return label[start+1 : end]
}

// DeriveParameterKey derives from a given label a regex pattern.
// e.g. :id[^\d+$] → id
// e.g. :id        → id
func DeriveParameterKey(label string) string {
	start := strings.Index(label, ParameterDelimiter)
	end := strings.Index(label, PatternDelimiterStart)

	if end == -1 {
		end = len(label)
	}

	return label[start+1 : end]
}

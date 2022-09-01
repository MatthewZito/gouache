package utils

// Contains determines whether a given list of strings `l` contains a given target string `t`.
func Contains(l []string, t string) bool {
	for _, c := range l {
		if c == t {
			return true
		}
	}

	return false
}

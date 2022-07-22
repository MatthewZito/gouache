package premux

import (
	"fmt"
	"regexp"
	"sync"
)

// RegexCache maintains a thread-safe cache for compiled regular expressions.
type RegexCache struct {
	state sync.Map
}

// MakeCache constructs and returns a pointer to a new RegexCache.
func MakeCache() *RegexCache {
	cache := &RegexCache{
		state: sync.Map{},
	}

	return cache
}

// Get retrieves a compiled regex from the RegexCache, or creates one and caches it if not extant.
func (rc *RegexCache) Get(pattern string) (*regexp.Regexp, error) {
	v, ok := rc.state.Load(pattern)
	if ok {
		// Verify the validity of the cached regex.
		regex, ok := v.(*regexp.Regexp)
		if !ok {
			// @todo delete entry?
			return nil, fmt.Errorf("the given pattern %s is not a valid regular expression", pattern)
		}
		// Return the cached regex.
		return regex, nil
	}

	// Compile the regex and add to cache if valid.
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	rc.state.Store(pattern, regex)
	return regex, nil
}

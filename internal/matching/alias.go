package matching

import "strings"

// AliasMatch compares candidate against aliases and returns the associated canonical value.
//
// Comparison is case-insensitive and ignores leading/trailing whitespace.
func AliasMatch(candidate string, aliases map[string]string) (string, bool) {
	normalizedCandidate := strings.TrimSpace(strings.ToLower(candidate))
	for alias, canonical := range aliases {
		if normalizedCandidate == strings.TrimSpace(strings.ToLower(alias)) {
			return canonical, true
		}
	}
	return "", false
}

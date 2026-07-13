package matching

import "strings"

// ExactMatch compares candidate against a list of available values.
//
// Comparison is case-insensitive and ignores leading/trailing whitespace.
// It returns the original matching value from available when found.
func ExactMatch(candidate string, available []string) (string, bool) {
	normalizedCandidate := strings.TrimSpace(strings.ToLower(candidate))
	for _, item := range available {
		if normalizedCandidate == strings.TrimSpace(strings.ToLower(item)) {
			return item, true
		}
	}
	return "", false
}

package matching

// Resolve returns the canonical value for candidate from available or aliases.
//
// It first looks for an exact match in available, then falls back to aliases.
func Resolve(candidate string, available []string, aliases map[string]string) (string, bool) {
	if matched, ok := ExactMatch(candidate, available); ok {
		return matched, true
	}
	return AliasMatch(candidate, aliases)
}

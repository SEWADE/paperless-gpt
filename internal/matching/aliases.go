package matching

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

// LoadAliases loads aliases from a JSON object at path.
//
// Alias keys are normalized case-insensitively with leading and trailing
// whitespace removed. Canonical values are preserved exactly as provided.
func LoadAliases(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return map[string]string{}, nil
	}
	if err != nil {
		return nil, err
	}

	aliases := map[string]string{}
	if err := json.Unmarshal(data, &aliases); err != nil {
		return nil, err
	}

	normalizedAliases := make(map[string]string, len(aliases))
	for alias, canonical := range aliases {
		normalizedAliases[strings.TrimSpace(strings.ToLower(alias))] = canonical
	}
	return normalizedAliases, nil
}

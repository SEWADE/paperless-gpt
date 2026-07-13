package matching

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadAliases(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    map[string]string
		wantErr bool
	}{
		{
			name: "normalizes keys and preserves values",
			content: `{
				"  Globus Baumarkt GmbH & Co. KG  ": "GLOBUS BAUMARKT",
				"TELEKOM DEUTSCHLAND": " Deutsche Telekom "
			}`,
			want: map[string]string{
				"globus baumarkt gmbh & co. kg": "GLOBUS BAUMARKT",
				"telekom deutschland":           " Deutsche Telekom ",
			},
		},
		{
			name:    "empty object",
			content: `{}`,
			want:    map[string]string{},
		},
		{
			name:    "invalid JSON",
			content: `{"globus":`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "aliases.json")
			if err := os.WriteFile(path, []byte(tt.content), 0o600); err != nil {
				t.Fatal(err)
			}

			got, err := LoadAliases(path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("LoadAliases() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAliases() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestLoadAliasesMissingFile(t *testing.T) {
	got, err := LoadAliases(filepath.Join(t.TempDir(), "missing.json"))
	if err != nil {
		t.Fatalf("LoadAliases() error = %v, want nil", err)
	}
	if !reflect.DeepEqual(got, map[string]string{}) {
		t.Errorf("LoadAliases() = %#v, want empty map", got)
	}
}

func TestLoadAliasesIOError(t *testing.T) {
	_, err := LoadAliases(t.TempDir())
	if err == nil {
		t.Fatal("LoadAliases() error = nil, want I/O error")
	}
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("LoadAliases() error = %v, want non-not-exist I/O error", err)
	}
}

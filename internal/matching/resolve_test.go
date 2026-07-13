package matching

import "testing"

func TestResolve(t *testing.T) {
	cases := []struct {
		name      string
		candidate string
		available []string
		aliases   map[string]string
		want      string
		found     bool
	}{
		{
			name:      "exact match returns available canonical value",
			candidate: "invoice",
			available: []string{"Invoice", "Receipt"},
			aliases:   map[string]string{"invoice": "Invoices"},
			want:      "Invoice",
			found:     true,
		},
		{
			name:      "exact match takes precedence over alias",
			candidate: "receipt",
			available: []string{"Receipt"},
			aliases:   map[string]string{"receipt": "Receipts"},
			want:      "Receipt",
			found:     true,
		},
		{
			name:      "alias fallback returns canonical value",
			candidate: "bill",
			available: []string{"Invoice", "Receipt"},
			aliases:   map[string]string{"bill": "Invoice"},
			want:      "Invoice",
			found:     true,
		},
		{
			name:      "alias comparison ignores case and whitespace",
			candidate: "  RECEIPT COPY ",
			available: []string{"Invoice"},
			aliases:   map[string]string{" receipt copy ": "Receipt"},
			want:      "Receipt",
			found:     true,
		},
		{
			name:      "no exact or alias match",
			candidate: "statement",
			available: []string{"Invoice", "Receipt"},
			aliases:   map[string]string{"bill": "Invoice"},
			want:      "",
			found:     false,
		},
		{
			name:      "empty canonical alias value is found",
			candidate: "unknown",
			aliases:   map[string]string{"unknown": ""},
			want:      "",
			found:     true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := Resolve(tc.candidate, tc.available, tc.aliases)
			if ok != tc.found || got != tc.want {
				t.Errorf("Resolve(%q, %q, %q) = %q, %v; want %q, %v",
					tc.candidate, tc.available, tc.aliases, got, ok, tc.want, tc.found)
			}
		})
	}
}

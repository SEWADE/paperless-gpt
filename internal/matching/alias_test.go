package matching

import "testing"

func TestAliasMatch(t *testing.T) {
	cases := []struct {
		name      string
		candidate string
		aliases   map[string]string
		want      string
		found     bool
	}{
		{
			name:      "exact alias returns canonical value",
			candidate: "invoice",
			aliases:   map[string]string{"invoice": "Invoices"},
			want:      "Invoices",
			found:     true,
		},
		{
			name:      "case insensitive alias",
			candidate: "ReCeIpT",
			aliases:   map[string]string{"receipt": "Receipts"},
			want:      "Receipts",
			found:     true,
		},
		{
			name:      "candidate and alias whitespace ignored",
			candidate: "  bill  ",
			aliases:   map[string]string{" bill ": "Bills"},
			want:      "Bills",
			found:     true,
		},
		{
			name:      "missing alias",
			candidate: "statement",
			aliases:   map[string]string{"invoice": "Invoices"},
			want:      "",
			found:     false,
		},
		{
			name:      "empty alias",
			candidate: "  ",
			aliases:   map[string]string{"": "Uncategorized"},
			want:      "Uncategorized",
			found:     true,
		},
		{
			name:      "empty canonical value is found",
			candidate: "unknown",
			aliases:   map[string]string{"unknown": ""},
			want:      "",
			found:     true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := AliasMatch(tc.candidate, tc.aliases)
			if ok != tc.found || got != tc.want {
				t.Errorf("AliasMatch(%q, %q) = %q, %v; want %q, %v",
					tc.candidate, tc.aliases, got, ok, tc.want, tc.found)
			}
		})
	}
}

package matching

import "testing"

func TestExactMatch(t *testing.T) {
	cases := []struct {
		candidate string
		available []string
		want      string
		found     bool
	}{
		{"  Foo ", []string{"foo", "bar"}, "foo", true},
		{"BAR", []string{"foo", "Bar", "baz"}, "Bar", true},
		{"missing", []string{"foo", "bar"}, "", false},
		{"  baz", []string{"baz  ", "qux"}, "baz  ", true},
	}

	for _, c := range cases {
		got, ok := ExactMatch(c.candidate, c.available)
		if ok != c.found || got != c.want {
			t.Errorf("ExactMatch(%q, %q) = %q, %v; want %q, %v",
				c.candidate, c.available, got, ok, c.want, c.found)
		}
	}
}

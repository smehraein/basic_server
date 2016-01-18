package sanitize

import "testing"

func TestSelect(t *testing.T) {
	cases := []struct {
		in      string
		choices []string
		want    bool
	}{
		{"apple", []string{"apple", "banana", "orange"}, true},
		{"banana", []string{"apple", "banana", "orange"}, true},
		{"egg", []string{"apple", "banana", "orange"}, false},
	}
	for _, c := range cases {
		got := Select(c.in, c.choices)
		if got != c.want {
			t.Errorf("Select(%q) with choices(%q) == %t, want %t", c.in, c.choices, got, c.want)
		}
	}
}

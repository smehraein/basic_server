package sanitize

import "testing"

func TestRadio(t *testing.T) {
	cases := []struct {
		in      int
		choices []int
		want    bool
	}{
		{1, []int{1, 2}, true},
		{2, []int{1, 2}, true},
		{3, []int{1, 2}, false},
	}
	for _, c := range cases {
		got := Radio(c.in, c.choices)
		if got != c.want {
			t.Errorf("Radio(%q) with choices(%q) == %t, want %t", c.in, c.choices, got, c.want)
		}
	}
}

package sanitize

import "testing"

func TestEmail(t *testing.T) {
  cases := []struct {
    in string
    want bool
  }{
    {"soroushmehraein@gmail.com", true},
    {"soroushmehraein", false},
  }
  for _, c := range cases {
    got := Email(c.in)
    if got != c.want {
      t.Errorf("Email(%q) == %t, want %t", c.in, got, c.want)
    }
  }
}

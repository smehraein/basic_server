package sanitize

import "testing"

func TestName(t *testing.T) {
  cases := []struct {
    in string
    want bool
  }{
    {"soroushmehraein", true},
    {"SOROUSHMEHRAEIN", true},
    {"/soroushmehraein", false},
    {".soroushmehraein", false},
    {"<soroushmehraein", false},
    {">soroushmehraein", false},
    {"", false},
  }
  for _, c := range cases {
    got := Name(c.in)
    if got != c.want {
      t.Errorf("Name(%q) == %t, want %t", c.in, got, c.want)
    }
  }
}

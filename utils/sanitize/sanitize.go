package sanitize

import (
  "regexp"
)

func Email(s string) bool {
  if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, s); !m {
    return false
  } else{
    return true
  }
}

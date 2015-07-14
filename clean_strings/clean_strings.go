package clean_strings

import (
  "regexp"
  "log"
)

func CleanStrings(dirty_strings ...string) []string{

  // Setup Regexp
  reg, err := regexp.Compile("[^A-Za-z]")
  if err != nil {
    log.Fatal(err)
  }

  var clean_strings = make([]string, len(dirty_strings))

  for i := 0; i < len(dirty_strings); i++ {
    clean_strings[i] = reg.ReplaceAllString(dirty_strings[i], "")
  }
  return clean_strings
}
package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)
func Capitalize(st string) []string {
  var newString1 string
  var newString2 string
  var slice []string

	//slice := make([]string, 0, 2)

  for i, j := range st {
    if i%2 != 0 {
      newString1 += strings.ToUpper(string(j))
	  fmt.Println(newString1)
      newString2 += strings.ToLower(string(j))
	  fmt.Println(newString2)
      continue
    }

    if i%2 == 0 {
      newString2 += strings.ToUpper(string(j))
      newString1 += strings.ToLower(string(j))
      continue
    }
  }

  slice = append(slice, newString2, newString1)
  return slice
}

func TestCapilalize(t *testing.T) {
	result := Capitalize("prayoga")
	fmt.Println(result)
}
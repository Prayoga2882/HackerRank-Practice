package hackerrank_test

import (
	"fmt"
	"testing"
	"unicode"
)

func Solve(s string) []int {
  
	var results = []int{0,0,0,0} 
  
  for _, char := range s { //
    fmt.Println(char)
    if(unicode.IsUpper(char)) {
      results[0]++
    } else if (unicode.IsLower(char)) {
      results[1]++
    } else if (unicode.IsDigit(char)) {
      results[2]++
    } else {
      results[3]++
    }
  }
  return results
}

func TestSolve(t *testing.T) {
	result := Solve("agoy")
	fmt.Println(result)
}
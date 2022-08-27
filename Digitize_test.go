package hackerrank_test

import (
	"fmt"
	"testing"
)

func Digitize(n int) []int {
	// your code here
	var result []int
	fmt.Println(n)
	for {
	  result = append(result, n % 10)
	  n = n / 10
	  if n == 0 {
		return result
	  }
	}
  }


func TestDigitize(t *testing.T) {
	result := Digitize(35231)
	fmt.Println(result)
}

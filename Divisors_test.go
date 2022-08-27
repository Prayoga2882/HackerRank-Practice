package hackerrank_test

import (
	"fmt"
	"testing"
)

func Divisors(n int)int{
	result := 0
	  for i:=1; i<=n; i++ {
		if n%i == 0 {
		  result += 1
		}
	  }
	return result
  }

func TestDivisors(t *testing.T) {
	result := Divisors(3)
	fmt.Println(result)
}
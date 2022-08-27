package hackerrank_test

import (
	"fmt"
	"testing"
)

func GetMiddle(s string) string {
	n := len(s)
   if n==0 {
	 return s      
   }
   return s[(n - 1) / 2 : n / 2 + 1]
}

func TestGetMiddleString(t *testing.T) {
	result := GetMiddle("Prayoga")
	fmt.Println(result)
}
package hackerrank_test

import (
	"fmt"
	"testing"
)


func FakeBin(x string) string {
   var result []rune
   for i:=0; i<len(x); i++ {
		if x[i] < '5' {
			result = append(result, '0')
		} else if x[i] > 5 {
			result = append(result, '1')
		}
   }
   return string(result)
}

func TestFakeBin(t *testing.T) {
	result := FakeBin("16")
	fmt.Println(result)
}
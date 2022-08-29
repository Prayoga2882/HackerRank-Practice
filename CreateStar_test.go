package hackerrank_test

import (
	"fmt"
	"testing"
)

func Star(){
	result := ""

	for i:=5; i>0; i-- {
		result = result + "*"
		fmt.Println(result)
	}
	// fmt.Println(result)
}

func TestStar(t *testing.T) {
	Star()
}
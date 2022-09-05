package hackerrank_test

import (
	"fmt"
	"testing"
)

func Star() {
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

//	result := ""
//
//	for i := 5; i > 0; i-- {
//		result = result + "*"
//		fmt.Println(result)
//	}
//}

func TestStar(t *testing.T) {
	Star()
}

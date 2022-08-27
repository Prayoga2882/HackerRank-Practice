package hackerrank_test

import (
	"fmt"
	"testing"
)

func FindMultiples(integer, limit int) []int {
	result := []int{}
	for i := integer; i <= limit; i += integer {
		result = append(result, i)
		fmt.Println(result)
	}
	return result
}

func TestFindMultiples(t *testing.T) {
	result := FindMultiples(2, 6)
	fmt.Println(result)
}


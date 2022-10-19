package hackerrank_test

import (
	"fmt"
	"testing"
)

func PositiveSum(arr []int) int {
	var result int
	for _, element := range arr {
		if element > 0 {
			result += element
		}
	}
	return result
}

func TestPositiveSum(t *testing.T) {
	result := PositiveSum([]int{1, 2, -1, 3})
	fmt.Println(result)
}

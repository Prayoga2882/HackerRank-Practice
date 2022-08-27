package hackerrank_test

import (
	"fmt"
	"testing"
)

func ReduceButGrow(arr []int) int {
	result := 1

	for _, data := range arr {
		result *= data
	}
	return result
}

func TestReduceButGrow(t *testing.T) {
	result := ReduceButGrow([]int{1,2,3,4,5})
	fmt.Println(result)
}
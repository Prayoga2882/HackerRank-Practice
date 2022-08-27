package hackerrank_test

import (
	"fmt"
	"sort"
	"testing"
)

func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	result := [2]int{arr[0], arr[len(arr)-1]}
	return result
}

func TestMinMax(t *testing.T) {
	result := MinMax([]int{1,2,3,4,5})
	fmt.Println(result)
}
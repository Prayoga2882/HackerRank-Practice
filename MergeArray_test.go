package hackerrank_test

import (
	"fmt"
	"sort"
	"testing"
)

func MergeArrays(arr1, arr2 []int) []int {
	arr := append(arr1,arr2...)
	sort.Ints(arr)
	var res []int
	res = append(res,arr[0])
	for i:=1;i<len(arr);i++ {
	  if arr[i] != arr[i-1]{
		res = append(res,arr[i])
	  }
	}
	return res
  }

func TestMergeArray(t *testing.T) {
	result := MergeArrays([]int{1,2,4,5,6}, []int{1,2,4,6,3,2})
	fmt.Println(result)
}
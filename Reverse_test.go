package hackerrank_test

import (
	"fmt"
	"testing"
)

// append itu nambahin ke dalam slice
// kalaupun index pertama di hapus dan di append(ditambahkan index pertama)
// maka data yg di tambahkan ya masuk ke dalam slice

func Reverse(xs []int) []int {
	if len(xs) == 0 {
		return xs
	}
	fmt.Println(xs)
	return append(Reverse(xs[1:]), xs[0])
}

func Reverse2(value []int) []int {
	var result []int
	for i := len(value) - 1; i >= 0; i-- {
		result = append(result, value[i])
	}
	return result
}

func TestReverse(t *testing.T) {
	result := Reverse([]int{1, 2, 3, 4, 5})
	fmt.Println(result)

	//result := Reverse2([]int{1, 2, 3, 4, 5})
	//fmt.Println(result)

}

package hackerrank_test

import (
	"fmt"
	"testing"
)

// If the key(values of the slice) is not equal
// to the already present value in new slice (list)
// then we append it. else we jump on another element.

// setiap angka yg di looping kan mendapat nilai nya true
// jadi kalo nilai itu ketemu lagi ya gak akan masuk kondisi
// krna nilai nya udah di set jadi true

func RemoveDuplicate(arr []int) []int {
	var keysMap = map[int]bool{}
	var list []int

	for _, element := range arr {
		value := keysMap[element]
		if !value {
			keysMap[element] = true
			list = append(list, element)
		}
	}
	return list
}

func TestRemoveDuplicate(t *testing.T) {
	result := RemoveDuplicate([]int{1, 1, 2, 3, 4})
	fmt.Println(result)
}

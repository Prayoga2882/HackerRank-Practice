package hackerrank_test

import (
	"fmt"
	"testing"
)

func RemoveDuplicate(arr []int) []int {
	keys := make(map[int]bool)
	var list []int

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.

	// setiap angka yg di looping kan mendapat nilai nya true
	// jadi kalo nilai itu ketemu lagi ya gak ada masuk kondisi
	// krna nilai nya udah di set jadi true

	for _, element := range arr {
		value := keys[element]
		if !value {
			fmt.Println(!value)
			keys[element] = true
			list = append(list, element)
		}
	}
	return list
}

func TestRemoveDuplicate(t *testing.T) {
	data := []int{1, 2, 2, 3, 1}
	result := RemoveDuplicate(data)
	fmt.Println(result)
}

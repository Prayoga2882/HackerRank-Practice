package hackerrank_test

import (
	"fmt"
	"testing"
)

func MultipleArr() {
	sample := [][]int{{1, 2, 3}, {3, 2, 1}}

	//fmt.Printf("columns : %d\n", len(sample[1]))
	//fmt.Println(sample)

	for _, element := range sample {
		fmt.Println(element)
		for _, element2 := range element {
			fmt.Println(element2)
		}
	}
}

func TestMultipleArr(t *testing.T) {
	MultipleArr()
}

package hackerrank_test

import (
	"fmt"
	"testing"
)

func ArrayMultiDimensi() {
	array := [4][2]string{
		{"singa", "macan"}, {"babi","gukguk"},{"harimau", "macan"},
	}

	for i:=0; i<len(array); i++{
		fmt.Println(i)
		fmt.Println(array)
		fmt.Println(array[i])
	}
}

func TestArrMulti(t *testing.T) {
	ArrayMultiDimensi()
}
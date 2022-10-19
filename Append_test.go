package hackerrank_test

import (
	"fmt"
	"testing"
)

func AppendInt(value []int) []int {
	var pow []int
	for _, v := range value { // setiap range data/v itu adalah nilai dari yang di iterasi
		fmt.Println(v)
		pow = append(pow, v)
	}
	return pow
}

func Append(value string) []string {
	var fruits []string

	for i := 0; i < len(value); i++ {
		fruits = append(fruits, value)
	}
	return fruits
}

func TestAppend(t *testing.T) {
	// result := Append("prayoga")
	// fmt.Println(result)
	//result := AppendInt([]int{0, 0, 0})
	//fmt.Println(result)
}

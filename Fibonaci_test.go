package hackerrank_test

import (
	"fmt"
	"testing"
)

func Fibonacci(n int) []int {
	var slice = []int{1, 2, 3, 4, 5}

	for i := 0; i < n; i++ {
		if slice[i] == 2 {
			slice[i] = 0
		}
	}
	result := 0
	for _, element := range slice {
		result += element
	}
	fmt.Println(result)
	return slice
}

func TestFibonacci(t *testing.T) {
	result := Fibonacci(4)
	fmt.Println(result)
}

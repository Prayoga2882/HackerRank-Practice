package hackerrank_test

import (
	"fmt"
	"sort"
	"testing"
)

func SmallestIntegerFinder(numbers []int) int {
	fmt.Println(numbers)
  	sort.Ints(numbers)
	fmt.Println(numbers)
  	return numbers[0]
}

func TestSmallerInteger(t *testing.T) {
	data := []int{1,2,4,124,341,656,-124,-155,-413}
	SmallestIntegerFinder(data)
}
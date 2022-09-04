package hackerrank_test

import (
	"fmt"
	"testing"
)

func Map() {
	data := make(map[int]string)
	data[1] = "me"
	data[2] = "you"

	var value, isExist = data[1]

	if isExist {
		fmt.Println(value)
	}
}

func TestMap(t *testing.T) {
	Map()
}

package hackerrank_test

import (
	"fmt"
	"testing"
)

func Map(){
	data := make(map[int]string)
	data[0] = "aku"
	data[1] = "kamu"

	var value, isExist = data[1]

	if isExist {
		fmt.Println(value)
	}
}

func TestMap(t *testing.T) {
	Map()
}
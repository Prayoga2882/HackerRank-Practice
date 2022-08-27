package hackerrank_test

import (
	"fmt"
	"testing"
)

func RepeatedStr(rep int, value string) string {
	result := ""
	for i:=1; i<=rep; i++ {
		result =result + value
	}

	return result
}

func TestRepeatedStr(t *testing.T) {
	result := RepeatedStr(5, "hai")
	fmt.Println(result)
}
package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func StringToArray(str string) []string {
	char := []string{str}
	return char
}

func StringToArrayYeah(str string) []string {
	result := strings.Split(str, "")
	return result
}

func TestStringToArray(t *testing.T) {
	// result := StringToArray("agoy pra")
	result := StringToArray("agoy pra")
	fmt.Println(result)
}
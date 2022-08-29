package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func UpperCase(str string) string {
	return strings.ToUpper(str)
}

func TestUpper(t *testing.T) {
	result := UpperCase("agoy")
	fmt.Println(result)
}
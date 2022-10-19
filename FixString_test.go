package hackerrank_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func FixString(str string) string {
	numberLower := len(regexp.MustCompile("[a-z]").FindAllStringIndex(str, -1))
	fmt.Println(numberLower)
	numberUpper := len(regexp.MustCompile("[A-Z]").FindAllStringIndex(str, -1))
	fmt.Println(numberUpper)
	if numberLower >= numberUpper {
		return strings.ToLower(str)
	}
	return strings.ToUpper(str)
}

func TestFixString(t *testing.T) {
	//FixString("cOdEc")
	result := FixString("CoDeW")
	fmt.Println(result)
}

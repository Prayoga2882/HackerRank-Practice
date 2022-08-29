package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func Spacify(str string) string {
	word := strings.Split(str, "")
	return strings.Join(word, " ")
}

func TestSpacify(t *testing.T) {
	result := Spacify("agoy")
	fmt.Println(result)
}
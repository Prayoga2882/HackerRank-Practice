package hackerrank_test

import (
	"fmt"
	"testing"
)

func RemoveChar(word string) string{
	return word[1 : len(word)-1]
}

func TestRemoveChar(t *testing.T) {
	result := RemoveChar("login")
	fmt.Println(result)
}

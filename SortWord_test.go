package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func FindShort(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}

func TestFindShot(t *testing.T) {
	result := FindShort("dollar it gone")
	fmt.Println(result)
}

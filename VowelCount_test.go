package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func GetCount(str string) int {
	count := 0
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, data := range vowels {
		count += strings.Count(str, data)
	}
	return count
  }

func TestGetCount(t *testing.T) {
	result := GetCount("Aku bakar hutan")
	fmt.Println(result)
}
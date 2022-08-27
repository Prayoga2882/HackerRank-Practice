package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func isCaseSensitif(str string) bool {
	str = strings.ToUpper(str)
	for i:=0; i<len(str);i++ {
	   if str[i]!=str[len(str)-i-1] {
		return false
		}
	}
	return true
}

func TestIsCaseSensifit(t *testing.T) {
	result := isCaseSensitif("a")
	fmt.Println(result)
}
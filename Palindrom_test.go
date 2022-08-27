package hackerrank_test

import (
	"fmt"
	"regexp"
	"testing"
)

func IsPalindrome(str string) string {
    result := []byte{}
    for i := len(str)-1; i >= 0; i-- {
        result = append(result, str[i])
    }
    return string(result)
}

func GetAlfabet(str string) string{
	re := regexp.MustCompile("[^A-Za-z]")
	result := re.ReplaceAllString(str,"")
	return result
}


func TestPalindrom(t *testing.T) {
	result := GetAlfabet("dawdd98a7wd98qy2098ryq039 8ryq3907rncq73tnq397rtq7")
	fmt.Println(result)
	final := IsPalindrome(result)
	fmt.Println(final)
}
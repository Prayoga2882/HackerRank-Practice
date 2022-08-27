package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func GetTitleString(str string) string {
	// var result string
	arr := strings.Split(str," ")
	fmt.Println(arr)
    return strings.ToUpper(string(arr[0][0]) + "." + string(arr[0]))
}

func TestGetTitleString(t *testing.T) {
	result := GetTitleString("agoy lucu")
	fmt.Println(result)
}

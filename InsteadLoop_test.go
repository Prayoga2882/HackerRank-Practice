package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func InsteadLoop(){
	Word := ""
	
    for _, i := range("ABCDE") {
        for _,j := range("ABCDE") {
            Word = string(i) + string(j)
            fmt.Println(Word)
            if Word == "DC" {
                break
            }
        }
        if Word == "DC" {
            break
        }
    }
}

func AppendLoop(str string) []string {
    var result []string

    for _, element := range str{
        elementStr := string(element)
        if elementStr == "a" {
            code := strings.ReplaceAll(str, elementStr, "b")
            result = append(result, code)
        }
    }

    return result
}

func TestInsteadLoop(t *testing.T) {
	// InsteadLoop()

    result := AppendLoop("agoy")
    fmt.Println(result)
}
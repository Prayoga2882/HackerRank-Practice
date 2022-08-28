package hackerrank_test

import (
	"fmt"
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

func TestInsteadLoop(t *testing.T) {
	InsteadLoop()
}
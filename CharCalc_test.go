package hackerrank_test

import (
	"fmt"
	"testing"
)


func Calc(s string) int {
    n := 0
    for _,c := range s {
        if (c/10) == 7 {n++}
        if (c%10) == 7 {n++}
    }
    return 6*n
}

func TestCalc(t *testing.T) {
	result := Calc("aku gilaa")
	fmt.Println(result)
}
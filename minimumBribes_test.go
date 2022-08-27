package hackerrank_test

import (
	"fmt"
	"testing"
)

func maximum(a, b int32) int32 {
    if a >= b {
        return a
    }
    return b
}

func MinimumBribes(q []int32) {
    // Write your code here
    count := int32(0)
    for i := int32(len(q)) - 1; i >= 0; i-- {
        if q[i]-(i+1) > 2 {
            fmt.Println("Too chaotic")
            return
        }
        for j := maximum(0, q[i]-2); j < i; j++ {
            if q[j] > q[i] {
                count++
            }
        }
    }
    fmt.Println(count)
}

func TestMinimumBribes(t *testing.T) {
    var data []int32
    MinimumBribes(data)
}
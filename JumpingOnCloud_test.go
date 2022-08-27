package hackerrank_test

import "testing"

func JumpingOnClouds(c []int32) int32 {
    length := len(c)
    var count int32

    for i := 0; i < length - 1; {
        if i+2 == length || c[i+2] == 1{
            i++
            count++
        }else {
            i += 2
            count++
        }
    }

    return count
}

func TestJumpinOnClouds(t *testing.T) {
    var data []int32
    JumpingOnClouds(data)
}
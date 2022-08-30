package hackerrank_test

import (
	"fmt"
	"testing"
)

func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
	x := dadYearsOld - sonYearsOld * 2 
	if x >= 0 { 
		fmt.Println(x)
		return x 
	}
	return x 
  }

func TestTwiceOld(t *testing.T) {
	result := TwiceAsOld(22, 2)
	fmt.Println(result)
}
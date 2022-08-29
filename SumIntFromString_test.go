package hackerrank_test

import (
	"fmt"
	"strconv"
	"testing"
)

func SumMix(arr []any) int {
	sum := 0
	for _, val := range arr {
		fmt.Println("agoy : ",val)
	 switch val := val.(type){
	   case int:
		sum += val
	   case string:
		k,_ := strconv.Atoi(val)
		sum += k
	 }
	 
	}
	return sum
  }

func SumData(str []any) int {
	result := 0

	for _, element := range str {
		code, _ := strconv.Atoi(fmt.Sprintf("%v", element))
		result += code
	}
	return result
}

func TestSumInt(t *testing.T) {
	result := SumMix([]any{"123"})
	fmt.Println(result)
}
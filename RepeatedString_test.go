package hackerrank_test

import (
	"fmt"
	"testing"
)

func RepeatedString(s string, n int64) int64 {
	inputLength := int64(len(s))
	occurrenceOfAInInput := OccurrenceOfA(s)
	numberOfSubstring := int64(n / inputLength)

	numberOfAs := numberOfSubstring  * occurrenceOfAInInput
	reminder := n % inputLength // a

	for i := int64(0); i < reminder; i++ {
		if s[i]  == 'a' {
			numberOfAs++
		}
	}

	return numberOfAs
}

func OccurrenceOfA(s string) int64  {
	var count int64
	for i := range s {
		if s[i] == 'a' {
			count++
		}
	}

	return count
}

func TestRepeatedString(t *testing.T) {
	data1 := "Agoy"
	data2 := 123
	result := RepeatedString(data1, int64(data2))
	fmt.Print(result)
}
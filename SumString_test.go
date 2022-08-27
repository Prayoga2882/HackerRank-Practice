package hackerrank_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func SumOfIntegersInString(strng string) int {
	result := 0
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(strng, -1))

	for _, data := range re.FindAllString(strng, -1) {
		data, _ := strconv.Atoi(data)
		result += data
	}
	return result
  }

func TestSumString(t *testing.T) {
	result := SumOfIntegersInString("121oindo13idn13o09f139f")
	fmt.Println(result)
}
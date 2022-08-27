package hackerrank_test

import (
	"sort"
	"strings"
	"testing"
)

var s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	result := strings.Join(chars, "***")
	return result
  }

func TestSplit(t *testing.T){
	TwoSort(s)
}
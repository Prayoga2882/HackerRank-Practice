package hackerrank_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortString(t *testing.T) {
	strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
}
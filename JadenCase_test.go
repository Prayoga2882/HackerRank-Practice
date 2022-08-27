package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func ToJadenCase(str string) string {
	return strings.ToTitle(str)
  }

  func TestJadenCase(t *testing.T) {
	result := ToJadenCase("prayoga boedihartoyo")
	fmt.Println(result)
  }


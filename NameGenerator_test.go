package hackerrank_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func NameGenerator(word string) string {
	first := word[1:]
	log.Println(first)
  	last := word[len(word)-1:]
	log.Println(last)

  if first != last {
    return "The " + strings.ToTitle(word)
  }

  return strings.ToTitle(word) + word[1:]

}

func TestNameGenerator(t *testing.T) {
	result := NameGenerator("agoy")
	fmt.Println(result)
}
package hackerrank_test

import (
	"fmt"
	"testing"
)

func ForLoop() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
}

func ForLoopIndex(value string) {
	result := ""
	for i:=0; i<len(value); i++ {
		result = result + value
		fmt.Println(result)
	}
	fmt.Println(result)
}

func TestForLoop(t *testing.T) {
	ForLoop()
	// ForLoopIndex("Prayoga ")
}
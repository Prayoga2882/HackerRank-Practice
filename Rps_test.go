package hackerrank_test

import (
	"fmt"
	"testing"
)

func Rps(p1, p2 string) string {
	beats := map[string]string{ 
	  "rock": "scissors", 
	  "paper": "rock",
	  "scissors": "paper",
	}
	
	if beats[p1] == p2 {
	  return "Player 1 won!"
	} else if beats[p2] == p1 {
	  return "Player 2 won!"
	}
	return "Draw!"
  }

func TestRps(t *testing.T) {
	result := Rps("paper", "scissors")
	fmt.Println(result)
}

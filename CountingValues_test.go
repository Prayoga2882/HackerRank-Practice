package hackerrank_test

import (
	"fmt"
	"strings"
	"testing"
)

func CountingValleys(steps int, path string) (valleyCounter int) {
	pathList := strings.Split(path, "")
	  //fmt.Printf("pathlist: %s", pathList)
	  currentElevation := 0
	  var enteredValley bool
  
	  for i := 0; i < steps; i++ {
		  if pathList[i] == "U" {
			  currentElevation++
		  } else {
			  currentElevation--
		  }
		  if currentElevation < 0 && enteredValley == false {
			  enteredValley = true
			  valleyCounter++
		  }
		  if currentElevation >= 0 && enteredValley == true {
			  enteredValley = false
		  }
	  }
	  return valleyCounter
  }

func TestCountingValues(t *testing.T) {

tests := []struct {
	steps int
	path  string
	want  int
}{
	{8, "UDDDUDUU", 1},
	{12, "DDUUDDUDUUUD", 2},
}
for _, tc := range tests {
	t.Run(fmt.Sprintf("(%v,%v)", tc.path, tc.steps), func(t *testing.T) {
		got := CountingValleys(tc.steps, tc.path)
		if got != tc.want {
			t.Errorf("Return:  %v; want %v", got, tc.want)
		}
	})
}
}
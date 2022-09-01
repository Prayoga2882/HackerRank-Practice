package hackerrank_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func IntToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func SliceToString(str []string) string {
	return strings.Join(str, "")
}

var badWords = []string{"fuck", "shit", "tai"}

func ReplaceString(message string) (result string) {
	for _, element := range badWords {
		if message == element {
			log.Println("exist")
			result = "****"
			log.Println(result)
			return
		}
	}
	fmt.Println("finish")
	return
}

func TestReplaceString(t *testing.T) {
	//result := ReplaceString("melvin")
	//fmt.Println(result)
	ReplaceString("shit")
}

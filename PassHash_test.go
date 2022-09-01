package hackerrank_test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func PassHash(str string) string {
	data := []byte(str)
	md5ed := fmt.Sprintf("%x", md5.Sum(data))
	return md5ed
}

func TestPassHash(t *testing.T) {
	result := PassHash("melvin")
	fmt.Println(result)
}
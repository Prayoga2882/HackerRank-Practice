package hackerrank_test

import (
	"fmt"
	"testing"
	"time"
)

func TestExpiry(t *testing.T) {
	now := time.Now().UTC()
	expiry := time.Date(2009, 12, 1, 0, 0, 0, 0, time.UTC)
	expiresSoon := now.After(expiry.AddDate(0, 0, -12))
	fmt.Println("Now :", now)
	fmt.Println("expiry :", expiry)
	fmt.Println("expiresSoon :", expiresSoon)
}
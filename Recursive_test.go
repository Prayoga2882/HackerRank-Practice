package hackerrank_test

func Recursive(n int) int {
	if n == 0 {
		return 0
	}
	return n * Recursive(n-1)
}

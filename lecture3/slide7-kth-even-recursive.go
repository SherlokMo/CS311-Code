package main

import "fmt"

// ## Pseudo code
// Analysis:
//		Time Complexity: O(n)
//		Space Complexity: O(n) (N call-stack)
// Algorithm even(k):
//	if k <- 1:
//		return 2
//	return even(k <- k - 1) + 2

func even(k int) int {
	if k == 1 {
		return 2
	}

	return even(k-1) + 2
}

func main() {
	fmt.Println(even(7))
}

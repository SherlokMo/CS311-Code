package main

import "fmt"

// 	Equation b^e = b * b^e-1
// 	Find where e = 8
// 	constrains: b = 2
// 	base case:
// 	resursive call limit e = 1 // return 0

func solution(b, e int) int { // o(N)
	if e == 1 {
		return 1
	}

	return b * solution(b, e-1)
}

func main() {

	fmt.Println(solution(2, 8))
	// Output 128
}

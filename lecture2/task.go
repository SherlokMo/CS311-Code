package main

import "fmt"

/**
	#Pseudocode

	Algorithm solution(A) {
		Avg, min, max <- A[0]
		for i <- 1 to 4
			Avg <- Avg + A[i]
			if A[i] > max
				max <- A[i]
			if A[i] < min
				min <- A[i]

		Avg <- Avg / 5

		return Avg, min, max
	}
**/

func solution(A []int) (float64, int, int) {

	Avg, min, max := A[0], A[0], A[0]

	for i := 1; i <= 4; i++ {
		Avg += A[i]
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	return float64(Avg) / 5.0, min, max
}

func main() {

	numbers := []int{6, 4, 9, 2, 10}

	average, min, max := solution(numbers)

	fmt.Println("Average is: ", average)
	fmt.Println("Min is: ", min)
	fmt.Println("Max is: ", max)

}

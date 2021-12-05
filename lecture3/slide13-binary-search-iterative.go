package main

import "fmt"

/**
	#PseudoCode

	Algorithm binarySearch(A, n, target):
		low <- 0, high <- n - 1
		while low <= high:
			mid <- (low + high) / 2
			if A[mid] = target:
				return mid
			else if A[mid] > target:
				high = mid - 1
			else:
				low = mid + 1
		endwhile

		return -1
	endAlgorithm

**/

func binarySearch(A []int, target int) int {

	for low, high := 0, len(A)-1; low <= high; {
		mid := (low + high) / 2

		if A[mid] == target {
			return mid
		} else if A[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binarySearch(arr, 10))
}

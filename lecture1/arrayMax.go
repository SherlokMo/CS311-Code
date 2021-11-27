package main

func arrayMax(A []int, n int) int {

	max := A[0]
	for i := 1; i < n; i++ {

		if A[i] > max {
			max = A[i]
		}
	}

	return max
}

func main() {

	arr := []int{9, 1, 2, 45, 80, 100, 4}

	arrayMax(arr, len(arr))

}

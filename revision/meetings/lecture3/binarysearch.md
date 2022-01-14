## Pseudocode for binary search algorithm

## Iterative
```
	Algorithm binarySearch(A, n, target):
        Input: A Array of numbers, n size of the array, target the number we want to search
        Output: index of the number we want to reach it

 		low <- 0, high <- n - 1
		while low <= high do
			mid <- (low + high) / 2
			if A[mid] = target then
				return mid
			else if A[mid] > target then
				high <- mid - 1
			else:
				low <- mid + 1
		return -1

    Space O(1)
    Time  O(log(Ng))
```
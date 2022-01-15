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
    Time  O(log(N))
```
# Recursive binary search
```
	Algorithm binarySearch(arr, target, low, high)
		Input: Arr is a sorted array, target is what we are searching for, low is lower boundry, high 		   is the boundry limit.
		Output: index of the number we want to reach it otherwise we return nil (or -1)

		mid <- (low+high)/2  #note that mid is always casted to integer
		if arr[mid] = target then
			return mid
		elsif arr[mid] > target then
			high <- mid - 1
			return binarySearch(arr, target, low, high)
		elsif arr[mid] < target then
			low <- mid + 1
			return binarySearch(arr, target, low, high)

		return nil #Note that we are always returning something in each if statment, we don't need an 			   else for this return.
	Complexity:
		Space: O(log(N)) #We could have a log(N) call stacks at the same time (worst-case)
		Time:  O(log(N))
```
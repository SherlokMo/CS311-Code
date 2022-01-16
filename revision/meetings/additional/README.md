## Here are some additional pseudo codes from exams

* Linear Search
    ```
        Algorithm linearSearch(A, n, target)
            Input: A is an array of n size, target is what we are searching for
            Output: index of the target, nil if the target is not found 

            for i <- 0 to n - 1 do
                if A[i] = target then
                    return i
            
            return nil

        Complexity:
            Space: O(1)
            Time:  O(n)
    ```

* bubble sort
```
    Algorithm bubbleSort(arr, n)
        Input: Array arr of n size
        Output sorted array

        for i <- 0 to n - 1 do
            for j <- 0 to n - i - 1 do
                if arr[j] > arr[j + 1]  then
                    swap array[j] between arr[j + 1]

        return arr

    Complexity:
        Space: O(1)
        Time:  O(n^2) 
```
* Insertion sort
```
    Algorithm insertionSort(arr, n)
        Input: Array arr of n size
        Output sorted array

        for i <- 1 to n - 1 do
            temp <- arr[i]

            j <- i - 1
            while j >= 0 and arr[j] > temp do
                arr[j + 1] <- arr[j]
                j <- j - 1
            
            arr[i] <- temp
        return arr

    Complexity:
        Space: O(1)
        Time:  O(n^2)
```

* Selection sort

```
    Algorithm selectionSort(Arr, n)
        Input: Array arr of n size
        Output: Sorted array

        for i <- 0 to n - 2 do
            minIdx <- i
            for j <- i + 1 to n - 1 do
                if arr[j] < arr[minIdx] then
                    minIdx <- j
            swap between arr[i], arr[minIdx]

        return arr

    Complexity:
        Space: O(1)
        Time:  O(n^2)
```
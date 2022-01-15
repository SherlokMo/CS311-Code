# Important note
* Merge Sort divide the array into two halfes in n/2 ratio while quick sort splits the array into any ratio based on the pivot.
* Merge Sort worst case is O(nlog(n)) while quick sort worst case is O(n^2)
* Merge Sort is prefered in large datasets and the best way to sort an LinkedLists. while quick sort works better on minimum dataset.
* Merge Sort adds a new array to store the sorted array, quick sort sorts the array in place. 


# Merge Sort Pseudo Code & analysis

``` 
    Algorithm mergeSort(arr, n)
        Input: Arr is array on n intergers.
        Output: Sorted picture of Arr.

        if(n <= 1) 
            return arr
        
        (LeftHalf, RightHalf) <- partition(arr, n/2)

        C1 <- LeftHalf.size
        C2 <- RightHalf.size
        S <- merge(mergeSort(LeftHalf, C1), mergeSort(RightHalf, C2))
        return S

    Complexity:
        Space: O(nlog(n))
        Time:  O(nlog(n))
```

# Quick Sort 
Quick sort uses a random pivot to split the array there are several ways to chose it:
  * Always pick first element as pivot. 
  * Always pick last element as pivot (implemented below).
  * Pick a random element as pivot.
  * Pick median as pivot.

Quick sort complexity changes by chosing a better pivot.
    * If we choose the middle element we could end-up with the perfect pivot and the analysis could be O(nlog(n)), But in worst case scenrios our time complexity would be O(N^2)

# Quick Sort pseudo code

```
    Algorithm quickSort(Arr, n)
        Input: Array arr of n slots.
        Output: Sorted array

        if n <= 1 then
            return arr
        
        pivot = arr[n - 1]
        (LittleHalf, BiggerHalf) <- Parition(Arr, n, pivot)

        sortedArr <- merge(quickSort(LittleHalf, LittleHalf.size), pivot, quickSort(BiggerHalf, BiggerHalf.size))

        return sortedArr
    
    Complexity (worst-case): 
        Space: O(n)
        Time:  O(n^2)
```
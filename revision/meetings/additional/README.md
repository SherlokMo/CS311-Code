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
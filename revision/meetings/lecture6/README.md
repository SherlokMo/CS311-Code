# Lecture 6 topics
* What is backtracking?
* Problems
  * [Travelling sales person (TSP)](#travelling-sales-person)
  * [Subset Sum](#subset-sum)
  * [Longest Increasing subsequence](#longest-increasing-subsequence)

# Travelling sales person
Solving travelling sales man in backtracking approach requires an understanding of Hamilition cycle. So.. what is hamilition cycle exactly? <br/>
Hamilition cycle is a backtracking approach to not re-visit the visited nodes again in the parent tree, so that the ending and starting node are the same. Each node should be visited **Only once!** except the starting and ending node. the problem is to find a minimum weight Hamiltonian Cycle.

```
  global answer <- infinety
  Algorithm TSP(graph, visited, currPosition, n, count, cost)
      Input: graph: is a adjacency matrix, visited: nodes that are visited on this recursance, currPosition: current Position inside the graph, count:counter of how many nodes are visited, cost: cost of this track 
      Output: minimum hamilition cycle cost

      // base case
      if count = n and graph[currPosition][0] exists do
        answer <- min(ans, cost + graph[currPosition][0])
        backtrack to the parent
      
      for i <- 0 to n - 1 do 
        if visited[i] is false and graph[currPosition][i] exists do
          visited[i] <- true
          tsp(graph, visited, i, n, count+1, cost + graph[currPosition][i])

          visited[i] <- false // when backtracking

    Complexity: 
      Space: O(N)
      Time:  O(N!)   
```


# Subset sum
Let sum wanted to reach is 8 and we have a set [1, 3, 4, 5] We want to have a set of the subsets that can reach that sum
```
  Input:  [1, 3, 4, 5], 8
  Output: [[1, 3, 4], [3, 5]]
```

```
  this psuedo code works in a sorted set, so if the input is not sorted, we shall sort it first.
  global result <- empty set
  Algorithm subset(set, n, currentSet, currentIdx, target)
    Input: Set witn n size, currentSet of the recursion, currentIdx: is the index in set, target is the sum left
    Output: subset of set that sum up to target

    if sum = 0 then
      append currentSet in result
      backtrack to the root of this reccurance
    
    if currentIdx = n then
      backtrack to the root of this reccurance

    for i <- currentIdx to n - 1 do
      if set[i] <= target then
        append set[i] into currentSet
        subset(set, n, currentSet, i, target - set[i])
        pop set[i] from currentSet
      else
        backtrack to root of this reccurance
    
  Complexity:
    Space: O(N)
    Time:  O(2^N)
```

# Longest-increasing subsequence
Given an integer array `nums`, return the length of the longest strictly increasing subsequence. </br>
A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

```
  global max <- 1

  Algorithm LIS(arr, n, i, prev)
    if n = i then
      return 0

    excluded <- LIS(arr, n, i+1, prev)

    included <- 0
    if arr[i] > prev then
      include <- 1 + LIS(arr, n, i+1, arr[i])

    return max(included, excluded)

  Complexity:
    Space: O(n)
    Time:  O(2^N)
```
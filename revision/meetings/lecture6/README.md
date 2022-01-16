# Lecture 6 topics
* What is backtracking?
* Problems
  * [Travelling sales person (TSP)](#travelling-sales-person)
  * [Subset Sum](#subset-sum)

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
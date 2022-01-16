# Lecture 6 topics
* What is backtracking?
* Problems
  * [Travelling sales person (TSP)](#travelling-sales-person)
  * N-queens
  * different combinations


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
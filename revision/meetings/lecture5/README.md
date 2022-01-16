# What is a graph?

A graph is simply a Node with a pointer to another Node. When these Nodes combine together they make a graph! For example:<br/>
```
        20->10->90->9
        |           |
        25->5->2->50
```
A Graph consists of **Vertices** and **Edges** which will later be called an `V and E` So we can represent our Graph to be `G=(V,E)` an **undirected connected graph**. 

# Simple structure for a Graph
```
    struct node<Type> {
        val        <Type>
        neighbours hashMap<*node<Type>, int> #second argument is the weight in case of a weighted graph
    }
```

# Traversal
We can traverse a graph in multiple ways like: <br/>
* BFS (Breadth-first search)
* DFS (Depth-first search)  </br>
But we have to note that in certain problems we should traverse our graph in a specific way such as `Greeady algorithms`, if we are looking for a minumum weight there should be some constrains.

# What are spanning trees?
A spanning tree is a Sub graph let's denote it by `t=(V,E)`. It's a subset of the main graph but it has no cycles. But why do we need them? <br/>
We need spanning trees when we want to:
* Connect and link `N` cities by an electric network in minimum cost.
* Connect telephone lines
We can observe that our objective is to connect links together without having a cycle as not to waste **resources**. Minimum spanning trees have direct applications in the design of networks, including computer networks, telecommunications networks, transportation networks, water supply networks, and electrical grids.

# Famous Algorithms for finding minimum-cost spanning tree
In minimum-cost spanning tree we can observe that we are using `Greeady Algorithm`.. but why? <br/>
Greeady algorithm is concerned by selecting the minimum local optimum at each step to find a global optimum. which is what we want in this type of problem. Let's get back to our two famouse algorithms: </br>
* Kruskal's Algorithm (consists of edges only)
* Prime's Algorithm   (consists of nodes & edges)

# Prime's Algorithm
Start with any node in the graph, and repeatedly choose the minmum weight to add it. This node should point to another node that is has not been chosen earlier.<br/>
Resource for studying: [Open Link](https://www.geeksforgeeks.org/prims-minimum-spanning-tree-mst-greedy-algo-5/)

# Kruskal Algorithm
Start with **no node or edges** and repeatedly add the minimum weight that does not creat a cycle. <br/>
Resource for solving technique: [Open Link](https://www.geeksforgeeks.org/kruskals-minimum-spanning-tree-algorithm-greedy-algo-2/)

## Pseudo codes

# Breadth first search for finding spanning trees
```
    Algorithm BFST(V)
        Input: Vertices V
        Output: Spanning tree

        declear queue, visitedMap

        Push V in queue
        visitedMap[V] <- true
        while queue is not empty do
            currentVertex <- Pop from queue
            output(currentVertex)
            for currentVertex.neighbours as vertex do
                if visitedMap[vertex] is nil then
                    visitedMap[vertex] <- true
                    Push vertex in queue
                    Insert vertex to spanningTree
    
    Complexity analysis:
        V stands for the number of vertices in our graph
        Space: O(V)
        Time:  O(V+E)
```
In case of adjency matrix in a graph the complexity changes to:
```
    Space: O(N)
    Time:  O(N^2)
```
**an adjacency matrix is a square matrix used to represent a finite graph**

# DFST (Depth First Search And Traversal)
I'm going to implement this pseudo code in **Recursive Approach**, but it could also be solved using normal `Stack` datastructure like the above problem but instead of queue we use **Stack**.
```
    Algoithm DFST(V, visited)
        Input: V is the start vertex, visited is the visited vertices.
        Output: spanning tree

        for V.neighbours as vertex do
            if visited[vertex] is nil then
                visited[vertex] <- true
                DFST[vertex, visited]
                Insert vertex to spanningTree
    
    Complexity:
        Space: O(V)
        Time:  O(V+E)
```

In case of adjency matrix in a graph the complexity changes to:
```
    Space: O(N)
    Time:  O(N^2)
```

# Dijkstra Shortest Path algorithm
Dijkstra’s algorithm is very similar to Prim’s algorithm for minimum spanning tree. Like Prim’s MST, we generate a SPT (shortest path tree) with a given source as a root. We maintain two sets, one set contains vertices included in the shortest-path tree, other set includes vertices not yet included in the shortest-path tree. At every step of the algorithm, we find a vertex that is in the other set (set of not yet included) and has a minimum distance from the source. <br />
Depending on how the algorithm is implemented and what data structure are used the time complexity is typically **O(E*Log(V))** which is competitive against other shortest path algorithms. Dijkstra algorithm finds the minimum distance between every vertex from the start node. </br>

```
    Algorithm dijkstra(g, v, s, e)
        Input: g an adjacency list of the weighted graph that has V nodes, starts from S node, and the end node is e
        Output: Shortest distance

        declear visited, priorityQueue
        distances <- fill array of V size with infinity value 
        distances[s] <- 0
        Insert (s, 0) to priorityQueue
        while priorityQueue is not empty do
            index, minValue <- pull from priorityQueue
            visited[index] <- true
            if distances[index] < minValue then
                exit this iteration and continue
            for g[index] as edge do 
                if visited[edge.to] then
                    exit this iteration and continue
                newDistance <- distances[index] + edge.cost
                if newDistance < distances[edge.to]
                    distances[edge.to] <- newDistance
                    Insert (edge.to, newDistance) to priorityQueue
        return distances

    Complexity:
        Space: O(V)
        Time:  O(E*Log(V))
```


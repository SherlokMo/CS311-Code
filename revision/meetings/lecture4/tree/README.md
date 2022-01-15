# What is a tree?

A tree is can be simplified as a pair of nodes that has pointers to another nodes.
```
        node
    -          -
    node       node

```

# Node structure
The node can be expressed with 3 things (In case of a binary tree). </br>
* value of the node
* left pointer to another node
* right pointer to the right node
```
    struct node {
        val       int|string|hashmap|...|any #any datatype
        leftNode  *node 
        rightNode *node
    }
```

# Binary tree
A binary tree in which each node has at least two children. <br/>
* A tree in which each node has precisely two children. (Every sub-tree should follow).
* Binary trees are recursive that means that each sub-tree should also contain atmost 2 childrens.
* Binary trees could be not symmetrical.

# Binary search tree (BST)
Binary search tree is a special condition of binary trees as it points to nodes in order. this mean that a binary search tree the `left < root < right` example: `25<50<65` in which we can search for numbers/strings easily inside it with time complexity of *O(log(N))* logarithmic function.
```
                50
            25       60
        10     27  55   65   
```

If we want to search for number "10", we can access it through `50>25>10`. <br/>

*Complexity*
We could have a space complexity of the entire highet of the tree, and logarthimic search in the best case (like binary search as we divide the tree in halves) thus: 
```
    h stands for height of the tree, N 
    Space: O(h) 
    Worst case:
        Time:  O(h)
    Best case:
        Time:  O(Log(n))
```

# Tree traversal approaches
* In-Order `left>root>right`   (This is a **DFS** search)
* Post-Order `left>right>root` (This is a **DFS** Search)
* Pre-Order `root>left>right`
* Breadth-First Search (**BFS**)

# Examples of traversal
Our tree is going to be represented with the below: <br/>
```
        1
    2       3
4      5   -  -    
```

Different traversing: <br/>
* In-Order: `4-2-5-1-3`
* Pre-Order: `1-2-4-5-3`
* Post-Order: `4-5-2-3-1`
* Breath-First: `1-2-3-4-5`

# Pseudo code

Let's assume our node data structure is this:
```
    struct Node {
        val any
        right *Node
        left  *Node
    }
```

* In-Order 
```
    Algorithm inOrderSearch(root)
        Input: root node of the tree
        Output: Printing nodes value

        if root is nil do
            return void
        
        inOrderSearch(root.left)
        output(root.val)
        inOrderSearch(root.right)
    
    Complexity:
        n Stands for nodes, h is the height of the tree
        Space: O(h) (call-stacks in memory)
        Time:  O(n)
```

* Post-Order 
```
    Algorithm postOrderSearch(root)
        Input: root node of the tree
        Output: Printing nodes value

        if root is nil do
            return void

        postOrderSearch(root.left)
        postOrderSearch(root.right)
        output(root.val)
    
    Complexity:
        n Stands for nodes, h is the height of the tree
        Space: O(h) (call-stacks in memory)
        Time:  O(n)
```

* Pre-Order 
```
    Algorithm preOrderSearch(root)
        Input: root node of the tree
        Output: Printing nodes value

        if root is nil do
            return void

        output(root.val)
        preOrderSearch(root.left)
        preOrderSearch(root.right)
    
    Complexity:
        n Stands for nodes, h is the height of the tree
        Space: O(h) (call-stacks in memory)
        Time:  O(n)
```

* Breadth-First Search

```
    Algorithm BFS(root)
        Input: root node of the tree
        Output: Pirinting BFS.

        declear queue
        queue.push(root)
        while queue is not empty do
            currentNode <- queue.pop()
            queue.push(currentNode.left)
            queue.push(currentNode.right)
            output(currentNode.value)
        
    Complexity:
        V stands for Vertices
        Space: O(V)
        Time:  O(V)
```

* Depth-First Search **Iterative Approach**
```
    Algorithm DFS(root)
        Input: root node of the tree
        Output: Pirinting DFS.
        
        declear stack
        stack.push(root)
        while stack is not empty do
            currentNode <- stack.pop()
            stack.push(currentNode.left)
            stack.push(currentNode.right)
            output(currentNode.value)
        
    Complexity:
        V stands for Vertices
        Space: O(V)
        Time:  O(V)
```
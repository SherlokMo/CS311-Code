# What is a tree?

A tree is can be simplified as a pair of nodes that has pointers to another nodes.
```
        node
    -          -
    node       node

```

# Node structure
The node can be expressed with 3 things (In case of a binary tree)
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
A binary tree in which each node has at least two children.
    * A tree in which each node has precisely two children. (Every sub-tree should follow).
    * Binary trees are recursive that means that each sub-tree should also contain atmost 2 childrens.
    * Binary trees could be not symmetrical.

# Binary search tree (BST)
Binary search tree is a special condition of binary trees as it points to nodes in order. this mean that a binary search tree the left < root < right (ex: 25<50<65) in which we can search for numbers/strings easily inside it with time complexity of *O(log(N))* logarithmic function.
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

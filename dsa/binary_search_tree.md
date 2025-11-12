# Binary Search Tree

<br>
<br>
<br>
<br>
<br>

## Problems

<br>
<br>
<br>

### Kth smallest element

- Problem: Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
- Link: https://leetcode.com/problems/kth-smallest-element-in-a-bst/

---

- **Algorithm**: Inorder traversal
- Inorder traversal of a BST iterate through the nodes in a sorted order.
- Maintain a variable `count=1` (1-indexed).
- Perform inorder traversal and check if the current node is `kth` node.
- If not, increment count `count++` and continue the traversal.
- Also implement early stopping when `kth` node is found.

<br>
<br>
<br>

### Kth smallest element in a volatile BST (⭐)

- Problem: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
- Link: https://leetcode.com/problems/kth-smallest-element-in-a-bst/

---

- **Algorithm**: Left & right count
- For each node, keep track of how many nodes are to it's left and to it's right.
- We can use this `left` and `right` count to find the `kth` smallest node in `O(logn)` time complexity.
- When a node is inserted or deleted, we just need to update the counts of it's parent elements.

<br>
<br>
<br>

### 

<br>
<br>
<br>

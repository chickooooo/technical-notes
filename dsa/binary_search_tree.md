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
- This transforms the BST into an **Order Statistic Tree**.
- We can use this `left` and `right` count to find the `kth` smallest node in `O(logn)` time complexity.
- When a node is inserted or deleted, we just need to update the counts of it's parent elements (while going down the path).

<br>
<br>
<br>

### LCA of a BST

- Problem: Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
- Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

---

- **Algorithm**: Left or right
- If `p` or `q` is same as current node, then current node is the LCA.
- If one is smaller and one is greater (values), then current node is the LCA.
- If both are smaller, go left.
- If both are greater, go right.

<br>
<br>
<br>

### Validate BST

- Problem: Given the root of a binary tree, determine whether it is a valid binary search tree (BST), where all left descendants are strictly less than the node and all right descendants are strictly greater.
- Link: https://leetcode.com/problems/validate-binary-search-tree/

---

- **Algorithm**: Node value within limits
- Maintain 2 variable `left` and `right`. These represents left and right limit of each node.
- Start with `left = -INF` and `right = +INF`.
- Check if the node value is within these limits. If not, return `false`.
- Go to left child and update `right = current.val`.
- Go to right child and update `left = current.val`.
- Recursively check if both left **and** right children are within limits.
- If current node is `nil`, return `true`.

<br>
<br>
<br>

### Construct BST from preorder traversal

- Problem: Given a preorder traversal array of a BST, reconstruct the BST and return its root.
- Link: https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

---

- **Algorithm**: Construct binary tree from preorder & inorder.
- Get the inorder traversal by sorting the preorder array.
- Iterate the preorder array. For each node, use the inorder array to find the nodes to it's left and right.
- Recursively construct the left & right subtree and attach them to the current node.

<br>
<br>
<br>

### Two sum on a BST

- Problem: Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.
- Link: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

---

- **Algorithm**: Inorder and reverse inorder traversal (using stack)
- Use a stack to iterate the tree using the **inorder** approach (ascending order).
- Use another stack to iterate the tree using the **reverse-inorder** approach (descending order).
- The top nodes in each stack represents the `start` and `end` nodes.
- Use the standard **Two Sum** approach to find the required 2 nodes.
- If the `start` and `end` pointers point to the same node, we have reached the end condition. 

<br>
<br>
<br>

### BST iterator

- Problem: Design an iterator for a binary search tree that traverses the nodes in in-order sequence.
- Link: https://leetcode.com/problems/binary-search-tree-iterator/

---

- **Algorithm**: Inorder traversal using stack
- Use a stack to store current node and it's left subtree.
- If `len(stack) > 0`, next element is present. Next element is the top element of the stack.
- Pop the top element, insert left subtree of it's right child in the stack and return popped value.

<br>
<br>
<br>

### 

<br>
<br>
<br>

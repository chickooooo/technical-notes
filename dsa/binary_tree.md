# Binary Tree

<br>
<br>
<br>
<br>
<br>

## Problems

<br>
<br>
<br>

### Level order traversal

- Problem: Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
- Link: https://leetcode.com/problems/binary-tree-level-order-traversal/

---

- **Algorithm**: Level order traversal
- Create a `temp` array and a `queue`.
- If the root is `nil`, return empty slice. Otherwise set `queue = [root, nil]`.
- Start removing one item at a time from the start of the queue.
- If the item is not `nil`, add it's value to `temp`. Then add left and right children of that item to the `queue`, if they exist. 
- If the item is `nil`, copy `temp` and add it to `output`. Then clear the `temp`. If the `queue` is not empty, add back `nil` and continue the process.
- Do so until the `queue` is empty.

<br>

- Note: Here `nil` item in the queue represents the end of that level.

<br>
<br>
<br>

### Maximum depth of binary tree

- Problem: Given the root of a binary tree, return its maximum depth. A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
- Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

---

- **Algorithm**: Max of left or right
- If `root` is nil, return `0`.
- Max depth of current node is max of (left subtree depth or right subtree depth) + 1.
- Recursively perform above logic from the root and return the answer.

<br>
<br>
<br>

### Diameter of binary tree

- Problem: Given the root of a binary tree, return the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root. The length of a path between two nodes is represented by the number of edges between them.
- Link: https://leetcode.com/problems/diameter-of-binary-tree/

---

- **Algorithm**: Max of left or right
- Create a variable `maximum` to keep track of maximum diameter. 
- Recursively find the max depth of each node: [Maximum depth of binary tree](#maximum-depth-of-binary-tree).
- In the recursive function, add one extra check for diameter.
- If 'current node + no. of nodes on left + no. of nodes on right' is greater than `maximum`, then update `maximum`.
- At last, return `maximum - 1`.

<br>
<br>
<br>

### 

<br>
<br>
<br>

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

### Level Order Traversal

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

### 

<br>
<br>
<br>

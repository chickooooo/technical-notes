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

### Same tree

- Problem: Given the roots of two binary trees p and q, write a function to check if they are the same or not. Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
- Link: https://leetcode.com/problems/same-tree/

---

- **Algorithm**: Bool of left and right
- For a pair of nodes to be identical:
    - Either both should be `nil`.
    - Or both should be not `nil` and should have the same value.
- Check if a pair of nodes is identical. If it is, recursively check the left and right pair.
- Also implement early stoping. If a pair is not identical, return `false` from there, don't evaluate other pairs.

<br>
<br>
<br>

### Zigzag level order traversal

- Problem: Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
- Link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

---

- **Algorithm**: Level order traversal
- Use standard [Level order traversal](#level-order-traversal).
- We will need an extra `forward` flag. Initially `forward` will be `true`.
- During LOT, when we encounter `nil` element, we will use this flag to determine whether to add `temp` to `output` in forward direction or reverse direction.
- After insertion, flip the `forward` flag. Rest of the template remains the same.

<br>
<br>
<br>

### Vertical order traversal

- Problem: You are given the root of a binary tree. Each node has a (row, col) position where left children go to (row+1, col–1) and right children to (row+1, col+1). You must return the vertical order traversal: group nodes by column from leftmost to rightmost, and within each column sort nodes first by row (top to bottom) and then by value if they share the same row and column.
- Link: https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/

---

- **Algorithm**: Preorder traversal with node position
- Maintain a `hashmap` that will hold nodes having the same `col`. `{col: [(row,node)]}`.
- Start with `row=0` & `col=0`. Use standard Preorder traversal.
- At each node, add `row` + `node.val` in `hashmap[col]`.
- Keep doing so, until whole tree is traversed.
- Now iterate from `min_col` to `max_col`. While iterating:
    - Sort `hashmap[col]`.
    - Add sorted node values to the `output`.
- At last return the `output`.

<br>
<br>
<br>

### Maximum path sum

- Problem: A path in a binary tree is a sequence of connected nodes with no repeats and does not need to include the root. The path sum is the total of its node values. Given a binary tree root, return the maximum path sum of any non-empty path.
- Link: https://leetcode.com/problems/binary-tree-maximum-path-sum/

---

- **Algorithm**: Return global max + local max
- For each node, return 2 values **global max** & **local max**.
- Global max:
    - Max path sum of that subtree.
    - Max of: LGM, RGM, C, C+LLM, C+RLM, C+LLM+RLM.
- Local max:
    - Max path sum, if that node is to be included.
    - Max of: C, C+LLM, C+RLM

<br>
<br>
<br>

### Maximum width of binary tree

- Problem: Given a binary tree root, return its maximum width, defined as the largest number of nodes (including gaps) between the leftmost and rightmost non-null nodes at any level.
- Link: https://leetcode.com/problems/maximum-width-of-binary-tree/

---

- **Algorithm**: Level order traversal + Node positions
- Node positions always follow this rule (1-indexed):
    - Left child position = parent position * 2
    - Right child position = (parent position * 2) + 1
- These positions are preserved even if there are missing nodes.
- Perform **Level order traversal** and find the max level width using these positions.

<br>
<br>
<br>

### Morris inorder traversal

- Problem: Given a binary tree. Find the inorder traversal of the tree without using recursion or extra stack space.
- Link: https://www.geeksforgeeks.org/problems/inorder-traversal-iterative/1

---

- **Algorithm**: Morris traversal
- This algorithm uses threads / temporary links that point back to the parent node.
- The binary tree is distorted while traversing down and is put back while going up.

---

- If the left node exists, find the rightmost node of the left node.
    - If rightmost node is `None`, add current node there. Go left.
    - If current node is already there, remove it. Print and go right.
- If the left node does not exist, print curent node and go to right node.

<br>
<br>
<br>

### Morris preorder traversal

- Problem: Given a Binary tree. Find the preorder traversal of the tree without using recursion. Try solving this with O(1) auxiliary space.
- Link: https://www.geeksforgeeks.org/problems/preorder-traversal-iterative/1

---

- **Algorithm**: Morris traversal
- Use same algorithm as Morris inorder traversal.
- Just adjust when nodes are printed and how the current pointer moves.
- Core working of temporary links stays the same.

<br>
<br>
<br>

### Right side view of binary tree

- Problem: Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
- Link: https://leetcode.com/problems/binary-tree-right-side-view/

---

- **Algorithm**: Preorder traversal
- Use an array to store the latest node value at each level.
- After preorder traversal, values at each level will be of right most nodes.

<br>
<br>
<br>

### Maximum product of splitted binary tree

- Problem: Given the root of a binary tree, split the binary tree into two subtrees by removing one edge such that the product of the sums of the subtrees is maximized. Return the maximum product of the sums of the two subtrees.
- Link: https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/

---

- **Algorithm**: Return subtree sum at each node
- At each node, we can either cutoff the left edge or right edge.
- The left node will give the sum of left subtree and similarly right node.
- Calculate maximum product and return it.

<br>
<br>
<br>

### Subtree containing deepest nodes

- Problem: Given the root of a binary tree, return the smallest subtree such that it contains all the deepest nodes in the original tree.
- Link: https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/

---

- **Algorithm**: Pass depth down, compute LCA while going up
- Perform postorder traversal of the binary tree.
- While traversing, keep track of the `depth` while going down the tree.
- At each node, check if left and right depth is same and `>= max_depth`.
- If so, update `max_depth` and `lca_node` with the current node.
- At the end of postorder traversal, `lca_node` will be the required node.

<br>
<br>
<br>

### All nodes at distance k

- Problem: Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.
- Link: https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

---

- **Algorithm**: Graph BFS
- First convert the tree into a graph by creating a `{node: parent}` hashmap.
- Then perform BFS from the start node and find all the nodes at distance `k`.
- While performing BFS, keep track of previous node and ignore it to avoid cycles.

<br>
<br>
<br>

### Symmetric tree

- Problem: Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
- Link: https://leetcode.com/problems/symmetric-tree/

---

- **Algorithm**: Traverse 2 trees simultaneously
- Use a single recursive function to traverse 2 trees together.
- Trees 1: `root.left`. Tree 2: `root.right`.
- Both nodes should be absent or present with same value.
- If the above condition satisfies, got to left of node 1 and right of node 2. And then right of node 1 and left of node 2.
- Keep doing recursively until whole tree is traverse or the condition fails.

<br>
<br>
<br>

### 

<br>
<br>
<br>

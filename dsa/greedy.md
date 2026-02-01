# Greedy

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Jump game

- Problem: You’re given an integer array nums where each element represents the maximum jump length from that position. Starting at index 0, determine whether you can reach the last index. Return true if you can, otherwise false.
- Link: https://leetcode.com/problems/jump-game/

---

- **Algorithm**: Next closest index to reach (O(n))
- Define a `reach` variable. It will hold the next closest index to reach from the current index.
- Set `reach = n-1`.
- Start iterating from `i = n-2` to `i = 0`.
- At each step, check if we can reach till the `reach` index.
    - If yes, set `reach = i`.
- At last, check if `reach == 0`. 

---

- **Algorithm**: Dynamic programming (O(n^2))
- Create a boolean array `setArray` of length `n`. This will tell whether we can reach the last index from the `ith` index.
- Set `setArray[n-1] = true`.
- Start iterating from `i = n-2` to `i = 0`.
    - At each step, again iterate from `j = i+1` to `j = i+nums[i]`.
        - If found `setArray[j] == true`, set `setArray[i] == true`.
        - Break the iteration.
- At last, check if `setArray[0] == true`.

<br>
<br>
<br>

### Distribute candy

- Problem: Given an integer array ratings representing children standing in a line, give candies such that:
    - Every child gets at least one candy.
    - A child with a higher rating than an adjacent child gets more candies.
Return the minimum total number of candies required.
- Link: https://leetcode.com/problems/candy/

---

- **Algorithm**: Left-right traversal + Right-left traversal
- **Intuition**: Each child should have more candies than less priority children to the left and right.
- Traverse from left to right, if current `priority` is greater than previous `priority`, increment candy `count`.
- Do same from right to left.
- For each index, take `max` of `left_count[i]` & `right_count[i]`.
- Return the total candy count.

<br>
<br>
<br>

### Jump game 2

- Problem: You are given a 0-indexed integer array nums, where each element represents the maximum number of steps you can jump forward from that index. Starting at index 0, you can move to any reachable index within the allowed jump range. The goal is to determine the minimum number of jumps needed to reach the last index (n - 1).
- Link: https://leetcode.com/problems/jump-game-ii/

---

- **Algorithm**: Monotonic stack
- Iterate from the back of the array.
- Maintain a strictly increasing monotonic stack: `(dist, index)`.
- If the current distance is lesser, we don't need more distance elements ahead of current element.
- After iteration, return the distance of the topmost item in the stack.

---

- **Algorithm**: Greedy / BFS
- Perform BFS on a 1D array. At each level, find out the elements in the next level.
- The min number of jumps is equal to the number of levels we need to cross to reach `n-1`th element.
- At each level, find the max index that can be reached from the elements of that level.
- This max index will be the right boundary of the next level.

<br>

```
0 1 2 3 4 5
2 1 1 2 3 4
_ ___ _ ___
```

<br>
<br>
<br>

### Maximum matrix sum

- Problem: You are given an n × n integer matrix. You may perform an operation any number of times where you choose two adjacent cells (sharing a side) and multiply both values by −1. Return that maximum possible matrix sum after all operations.
- Link: https://leetcode.com/problems/maximum-matrix-sum/

---

- **Algorithm**: Simple traversal
- If we try out various scenarios, we can observe that a pair of negative numbers can always be made positive.
- That means, even number of negative numbers will always become positive.
- If the matrix has odd number of negative numbers, at last only 1 negative number will be left.
- To maximize matrix sum, we can shift the last negative sign to the smallest num in the matrix.
- So, the problem simplifies to finding matrix sum, negative count and smallest absolute value.

<br>
<br>
<br>

### 

<br>
<br>
<br>

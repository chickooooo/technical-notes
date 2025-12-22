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

### Jump Game

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

### 

<br>
<br>
<br>

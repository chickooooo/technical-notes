# Dynamic Programming

<br>
<br>
<br>
<br>
<br>

## Problems

<br>
<br>
<br>

### Climbing stairs

- Problem: You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
- Link: https://leetcode.com/problems/climbing-stairs/

---

- Assumption: `f(i)` represents the number of ways to reach the last step, starting from `ith` step.
- To find: `f(0)`
- Functional equation: `f(i) = f(i+1) + f(i+2)`
- Range: `i ∈ [0, n]`

<br>

Top-down approach

- Create a 1D `memo` of size n+1 [0, n]. Set `memo[-1] = 1`.
- Start from `i=0`, if `i > n` return 0.
- If `f(i)` is already memoized, return the memoized value.
- Otherwise, use the **Functional equation** to calculate value of `f(i)`.
- Memoize `f(i)` and return the response.

<br>

Bottom-up approach

- Create a 1D `memo` of size n+1 [0, n].
- Set `memo[n] = 1` and `memo[n-1] = 1`.
- Iterate for `i ∈ [n-1, 0]` and populate the `memo`.
- Return `memo[0]`.

<br>

Space optimised

- At each step `i`, we only need `i+1` and `i+2`.
- So, set `last = 1` & `sec_last = 1`.
- Loop for `n-1` iterations and calculate `current`.
- After that update the values of `last` to `sec_last` and `sec_last` to `current`.
- Return `sec_last` at the end.

<br>
<br>
<br>

### House robber

- Problem: Given an integer array nums where nums[i] is the money in house i, pick a subset of non-adjacent houses to maximize the total stolen amount and return that maximum sum.
- Link: https://leetcode.com/problems/house-robber/

---

- Assumption: `f(i)` represents the maximum amount we can collect, starting from `ith` house and going to the right.
- To find: `f(0)`
- Functional equation: `f(i) = max(A[i] + f(i+2), f(i+1))`
- Range: `i ∈ [0, n-1]`

<br>

Top-down approach

- Create a 1D `memo` of size n [0, n-1].
- Start from `i=0`, if `i >= n` return 0.
- If `f(i)` is already memoized, return the memoized value.
- Otherwise, use the **Functional equation** to calculate value of `f(i)`.
- Memoize `f(i)` and return the response.

<br>

Bottom-up approach

- Create a 1D `memo` of size n+2 [0, n+1].
- Set `memo[n+1] = 0` and `memo[n] = 0`.
- Iterate for `i ∈ [n-1, 0]` and populate the `memo`.
- Return `memo[0]`.

<br>

Space optimised

- At each step `i`, we only need `i+1` and `i+2`.
- So, set `last = 0` & `sec_last = 0`.
- Loop from `n-1` to `0` and calculate `current`.
- After that update the values of `last` to `sec_last` and `sec_last` to `current`.
- Return `sec_last` at the end.

<br>
<br>
<br>

### Partition equal subset sum

- Problem: Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.
- Link: https://leetcode.com/problems/partition-equal-subset-sum/

---

- Assumption: `f(i, j)` returns true is it is possible to pick m elements from nums, whose sum is `j`, starting from `i` and going to right. Otherwise false.
- To find: `f(0, k)`. Where `k = sum(nums) / 2`
- Functional equation: `f(i, j) = f(i+1, j - A[i]) or f(i+1, j)`
- Range: `i ∈ [0, n-1]` and `j ∈ [k, 0]`

<br>

Top-down approach

- Create a 2D `memo` of size `n * (k+1)`.
- Start from `i = 0, j = k`.
- Return true if `j == 0` and return false if `j < 0 or i == n`.
- If `f(i, j)` is already memoized, return the memoized value.
- Otherwise, use the **Functional equation** to calculate value of `f(i, j)`.
- Memoize `f(i, j)` and return the response.

<br>

Bottom-up approach

- Create a 2D `memo` of size `(n+1) * (k+1)`.
- Set `memo[*][0] = True` (priority) and `memo[n][*] = False`.
- Iterate for `i ∈ [n-1, 0]`, `j ∈ [0, k]`.
- Populate the `memo` using the **Functional equation**.
- Return `memo[0][k]`.

<br>

Space optimised

- At each step `i`, we only need the next row `i+1`.
- So, set `temp = [True] + [False] * k` & `memo = [True] + [False] * k`.
- Iterate for `i ∈ [n-1, 0]`, `j ∈ [0, k]`. And populate the `temp`.
- Here `temp` refers to current row and `memo` refers to the next row.
- After that update the values of `memo` to `temp` and `temp` to `[True] + [False] * k`.
- At last, return `memo[k]`.

<br>
<br>
<br>

### Unique paths

- Problem: A robot moves from the top-left to the bottom-right of an m × n grid, moving only right or down. Find the number of unique paths it can take to reach the destination.
- Link: https://leetcode.com/problems/unique-paths/

---

- Assumption: `f(i, j)` represents the number of unique paths the robot can take to reach the destination starting from `ith` row and `jth` column.
- To find: `f(0, 0)`
- Functional equation: `f(i, j) = f(i+1, j) + f(i, j+1)`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`

<br>

Top-down approach

- Create a 2D `memo` of size `m * n`. Set `memo[m-1][n-1] = 1`.
- Start from `i = 0, j = 0`.
- Return `0` if `i == m` or `j == n`.
- If `f(i, j)` is already memoized, return the memoized value.
- Otherwise, use the **Functional equation** to calculate value of `f(i, j)`.
- Memoize `f(i, j)` and return the response.

<br>

Bottom-up approach

- Create a 2D `memo` of size `(m+1) * (n+1)`. Set `memo[m][n-1] = 1`.
- Iterate for `i ∈ [m-1, 0]`, `j ∈ [n-1, 0]`.
- Populate the `memo` using the **Functional equation**.
- Return `memo[0][0]`.

<br>

Space optimised

- At each step `i`, we only need the next row `i+1`.
- So, set `temp := make([]int, n+1)` & `current := make([]int, n+1)`. Also set `temp[n-1] = 1`.
- Iterate for `i ∈ [m-1, 0]`, `j ∈ [n-1, 0]`. And populate the `temp`.
- Here `temp` refers to next row and `current` refers to the current row.
- After that update the values of `temp` to `current` and `current` to `make([]int, n+1)`.
- At last, return `temp[0]`.

<br>
<br>
<br>

### Unique paths II

- Problem: Find the number of unique paths from the top-left to the bottom-right of a grid, moving only right or down, while avoiding cells marked as obstacles (1).
- Link: https://leetcode.com/problems/unique-paths-ii/

---

- Assumption: `f(i, j)` represents the number of unique paths the robot can take to reach the destination starting from `ith` row and `jth` column.
- To find: `f(0, 0)`
- Functional equation:
    - if `A[i][j] == 0` :`f(i, j) = f(i+1, j) + f(i, j+1)`
    - if `A[i][j] == 1` :`f(i, j) = 0`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`

<br>

Top-down approach

- Create a 2D `memo` of size `m * n`. Set `memo[m-1][n-1] = 1`.
- Start from `i = 0, j = 0`.
- Return `0` if `i == m` or `j == n`.
- If `f(i, j)` is already memoized, return the memoized value.
- Otherwise, use the **Functional equation** to calculate value of `f(i, j)`.
- Memoize `f(i, j)` and return the response.

<br>

Bottom-up approach

- Create a 2D `memo` of size `(m+1) * (n+1)`. Set `memo[m][n-1] = 1`.
- Iterate for `i ∈ [m-1, 0]`, `j ∈ [n-1, 0]`.
- Populate the `memo` using the **Functional equation**.
- Return `memo[0][0]`.

<br>

Space optimised

- At each step `i`, we only need the next row `i+1`.
- So, set `temp := make([]int, n+1)` & `current := make([]int, n+1)`. Also set `temp[n-1] = 1`.
- Iterate for `i ∈ [m-1, 0]`, `j ∈ [n-1, 0]`. And populate the `temp`.
- Here `temp` refers to next row and `current` refers to the current row.
- After that update the values of `temp` to `current` and `current` to `make([]int, n+1)`.
- At last, return `temp[0]`.

<br>
<br>
<br>

### Minimum path sum

- Problem: Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path. You can only move either down or right at any point in time.
- Link: https://leetcode.com/problems/minimum-path-sum/

---

- Assumption: `f(i, j)` represents minimum sum we can get while going to the bottom right corner, starting from `ith` row and `jth` column.
- To find: `f(0, 0)`
- Functional equation: `f(i, j) = min(f(i+1, j), f(i, j+1)) + A[i][j]`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`

<br>

Top-down approach

- Create a 2D `memo` of size `m * n`. Set `memo[m-1][n-1] = A[m-1][n-1]`.
- Start from `i = 0, j = 0`.
- Return `math.MaxInt` if `i == m` or `j == n`.
- If `f(i, j)` is already memoized, return the memoized value.
- Otherwise, use the **Functional equation** to calculate value of `f(i, j)`.
- Memoize `f(i, j)` and return the response.

<br>

Bottom-up approach

- Create a 2D `memo` of size `(m+1) * (n+1)`. Set `memo[m][n-1] = 0`.
- Iterate for `i ∈ [m-1, 0]`, `j ∈ [n-1, 0]`.
- Populate the `memo` using the **Functional equation**.
- Return `memo[0][0]`.

<br>

Space optimised

- At each step `i`, we only need the next row `i+1`.
- So, set `temp := make([]int, n+1)` & `current := make([]int, n+1)`. All the values of `temp` and `current` should be `math.MaxInt` inititally. Also set `temp[n-1] = 0`.
- Iterate for `i ∈ [m-1, 0]`, `j ∈ [n-1, 0]`. And populate the `temp`.
- Here `temp` refers to next row and `current` refers to the current row.
- After that update the values of `temp` to `current` and `current` to `make([]int, n+1)` (all MaxInt values).
- At last, return `temp[0]`.

<br>
<br>
<br>

### Paths in matrix whose sum is divisible by k

- Problem: Given an m×n grid and an integer k, count the number of paths from (0,0) to (m−1,n−1) moving only right or down such that the sum of values along the path is divisible by k. Return the count modulo 109 + 7.
- Link: https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/

---

- **Algorithm**: DP on grid. Top-left to bottom-right.
- In tabulation approach, we need to store `{num:count}` in each cell. Making it 3D DP problem.
- This num will range from 0 to k-1, repesenting the number of paths with `count = path_sum % k`.
- `f(i,j) = f(i+1,j) + f(i,j+1)`, here addition represents the summation of dictionary values, taking in consideration `grid[i][j]`.
- At last, return the value of `0` key from `memo[0][0]`.
- This represents all the paths starting from `(0,0)` that have `path_sum % k == 0`.

<br>
<br>
<br>

### Coin change

- Problem: Given an array of coin denominations and a target amount, find the minimum number of coins needed to make that amount. If it’s not possible using any combination of the coins, return -1.
- Link: https://leetcode.com/problems/coin-change/

---

- Assumption: `f(i, j)` represents the minimum number of coins required to make the amount `j`, picking coins from and after index `i`.
- To find: `f(0, amount)`
- Functional equation: `f(i, j) = min(f(i, j-A[i])+1, f(i+1, j))`
- Range: `i ∈ [0, n-1]`, `j ∈ [0, amount]`
- Base conditions:
    - If `j < 0` or `i == n`, then answer is `INF`.
    - If `j == 0`, then answer is 0.

---

- **Algorithm**: Pick or not pick
- Solve with top-down, bottom-up and space optimization approach.
- In the 3rd approach, use a single array `memo`, no need for `temp` array.
- Iterate `j` from `coins[i]` to `amount+1`, keeping previous values unchanged.

<br>
<br>
<br>

### Target sum

- Problem: Given an array of integers, count how many ways you can assign + or - to each number so that the resulting expression equals the target value.
- Link: https://leetcode.com/problems/target-sum/

---

- Standard DP problem. Follow standard steps.
- Functional equation, top-down, bottom-up, space optimization.
- Here, `j` ranges from `[-total, total]`. Adjust this range to `[0, 2*total]`.
- Make this adjustment to the `target` and the return value as well.

<br>
<br>
<br>

### Longest common subsequence

- Problem: Given two strings text1 and text2, return the length of their longest common subsequence (characters in the same order, not necessarily consecutive). Return 0 if none exists.
- Link: https://leetcode.com/problems/longest-common-subsequence/

---

- Assumption: `f(i, j)` represents the LCS starting from `ith` element in text1 and `jth` element in text2.
- To find: `f(0, 0)`
- Functional equation:
    - If `text1[i] == text2[j]`: `f(i, j) = f(i+1, j+1) + 1`
    - Else `f(i, j) = max(f(i+1, j), f(i, j+1))`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`
- Base conditions:
    - If `i == m` or `j == n`, then answer is `0`.

<br>
<br>
<br>

### Max dot product of two subsequences

- Problem: Given two arrays nums1 and nums2. Return the maximum dot product between non-empty subsequences of nums1 and nums2 with the same length.
- Link: https://leetcode.com/problems/max-dot-product-of-two-subsequences/

---

- Approach is similar to **Longest common subsequence**.
- It just has 1 more condition handling in the functional equation.
- We will add previous product to the sum only if it is greater than 0.

---

- Assumption: `f(i, j)` represents max dot product starting from index `i` in nums1 and index `j` in nums2, going right.
- To find: `f(0, 0)`
- Functional equation:
    - `f(i, j) = max(A, B, C)`
    - If `f(i+1, j+1) > 0`, `A = (N1[i] * N2[j]) + f(i+1, j+1)`.
    - Else `A = N1[i] * N2[j]`.
    - `B = f(i+1, j)`
    - `C = f(i, j+1)`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`
- Base conditions:
    - If `i == m` or `j == n`, then answer is `0`.

<br>
<br>
<br>

### Minimum ASCII delete sum

- Problem: Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.
- Link: https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/

---

- Approach is similar to **Longest common subsequence**.
- Calculate postfix ASCII sum of s1 & s2 for `O(1)` lookup for `sum(s1[i:])`.
- Rest of the algorithm is identical to LCS with minor tweaks.

---

- Assumption: `f(i, j)` represents the minimum delete sum starting from `ith` index in s1 and `jth` index in s2.
- To find: `f(0, 0)`
- Functional equation:
    - If `s1[i] == s2[j]`: `f(i, j) = f(i+1, j+1)`
    - Else `f(i, j) = min(s1[i] + f(i+1, j), s2[j] + f(i, j+1))`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`
- Base conditions:
    - If `i == m and j == n`, then answer is `0`,
    - Else if `i == m`, then answer is `sum(s2[j:])`.
    - Else if `j == n`, then answer is `sum(s1[i:])`.

<br>
<br>
<br>

### Maximal square

- Problem: Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
- Link: https://leetcode.com/problems/maximal-square/

---

- Intuition? Minimize the problem size. See how we can solve the base case, then 1 level above that case and so on. This will build up the intuition.
- This problem is slightly different than other DP problems.
- Other problems will give the optimal value after end of recursion/iteration.
- In the problem, we have to keep track of optimal value while iterating.


---

- Assumption: `f(i, j)` represents the max square side having top-left corner at `(i, j)`.
- To find: `max(f(i, j)) ^ 2`
- Functional equation:
    - If `A[i][j] == 0`: `f(i, j) = 0`
    - Else `f(i, j) = min(f(i+1, j+1), f(i, j+1), f(i+1, j)) + 1`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`
- Base conditions:
    - If `i == m` or `j == n`, then answer is `0`.

<br>
<br>
<br>

### Minimum insertion to make palindrome

- Problem: Given a string s. In one step you can insert any character at any index of the string. Return the minimum number of steps to make s palindrome.
- Link: https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

---

- Assumption: `f(i, j)` represents the minimum insertions for string starting from `ith` index and ending at `jth` index.
- To find: `f(0, n-1)`
- Functional equation:
    - If `s[i] == s[j]`: `f(i, j) = f(i+1, j-1)`
    - Else `f(i, j) = min(f(i+1, j), f(i, j-1)) + 1`
- Range: `i ∈ [0, n-1]`, `j ∈ [n-1, 0]`
- Base conditions:
    - If `i >= j`, then answer is `0`.

### Minimum falling path sum

- Problem: Given an n × n integer matrix, find the minimum sum of a falling path from the top row to the bottom. From each cell, you may move to the cell directly below or diagonally down-left or down-right in the next row.
- Link: https://leetcode.com/problems/minimum-falling-path-sum/

---

- Assumption: `f(i, j)` represents the minimum sum of falling path starting from `(i, j)th` index in the grid. 
- To find: `min(f(0, j))`
- Functional equation: `f(i, j) = min(f(i+1, j), f(i+1, j-1), f(i+1, j+1)) + A[i][j]`
- Range: `i ∈ [n-2, 0]`, `j ∈ [n-1, 0]`
- Base conditions:
    - If `j < 0` or `j == n`, then answer is `INF`.

<br>
<br>
<br>

### Distinct subsequences

- Problem: Given two strings s and t, return the number of distinct subsequences of s which equals t.
- Link: https://leetcode.com/problems/distinct-subsequences/

---

- Assumption: `f(i, j)` represents the no. of distinct subsequences starting from `ith` index of `s` and `jth` index of `t`.
- To find: `f(0, 0)`
- Functional equation:
    - If `s[i] == t[j]`, then `f(i, j) = f(i+1, j+1) + f(i+1, j)`
    - Else `f(i, j) = f(i+1, j)`
- Range: `i ∈ [0, m-1]`, `j ∈ [0, n-1]`
- Base conditions:
    - If `i == m & j == n` then answer is `1`.
    - If `i == m` then answer is `0`.
    - If `j == n` then answer is `1`.

<br>
<br>
<br>

### Buy sell stock with cooldown

- Problem: Given daily stock prices, find the maximum profit with unlimited transactions, allowing only one stock at a time and requiring a one-day cooldown after each sale.
- Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

---

- Assumption: `f(i, j)` represents the max profit starting at index `i` and have bought a stock at index `j`. If no stock bought previously then `j = -1`.
- To find: `f(0, -1)`
- Functional equation:
    - If `j == -1`, then `f(i, j) = max(f(i+1, i), f(i+1, -1))`
    - Else `f(i, j) = max(f(i+2, -1) + A[i] - A[j], f(i+1, j))`
- Range: `i ∈ [n-1, 0]`, `j ∈ [i-1, -1]`
- Base conditions:
    - If `i >= n` then answer is `0`.

---

- Note: Make adjustments to `j` range, to handle `-1`.

---

`O(n)` time complexity solution possible.

<br>
<br>
<br>

### 

<br>
<br>
<br>

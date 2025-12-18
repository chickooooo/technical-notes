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

### 

<br>
<br>
<br>

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

###

<br>
<br>
<br>

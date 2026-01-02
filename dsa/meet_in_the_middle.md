# Meet in the middle

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Closest Subsequence Sum

- Problem: Given an integer array nums and an integer goal, find a subsequence whose sum is closest to goal, and return the minimum absolute difference between the sum and goal.
- Link: https://leetcode.com/problems/closest-subsequence-sum/

---

- **Algorithm**:  Meet in the middle
- Split array in 2 equal halves.
- Generate subset sum of each half.
- Sort the second half.
- Iterate over H1 and find lower bound and upper bound in H2 using binary search.
    - lower bound: `y >= goal - x`
    - upper bound: `y <= goal - x`
- Return the minimum absolute difference.
- TC: `n * (2^(n/2))`

<br>
<br>
<br>

### 

<br>
<br>
<br>
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

### Closest subsequence sum

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

### Partition array to minimize sum difference

- Problem: Given an array nums of 2n integers, split it into two arrays of size n such that the absolute difference between their sums is minimized. Return that minimum difference.
- Link: https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/

---

- **Algorithm**: Meet in the middle
- The high level approach is same as standard MITM algorithm.
- When creating subset sum, group the output by number of elements included in that sum.
- When performing binary search, use only `lcount` & `rcount` groups such that `lcount + rcount = n/2`.
- Here, we only need to perform lower bound binary search. Upper bound is handled automatically.
- At last, return minimum difference * 2.

<br>
<br>
<br>

### 

<br>
<br>
<br>

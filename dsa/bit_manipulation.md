# Bit Manipulation

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Single number

- Problem: Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
- Link: https://leetcode.com/problems/single-number/

---

- **Algorithm**:  XOR: A ^ A = 0
- Create a variable `xor` initially set to 0.
- Perform XOR of all the numbers with this `xor`.
- Identical numbers will cancel out each other, leaving only the single number.

<br>
<br>
<br>

### Construct minimum bitwise array I

- Problem: Given an array nums of n prime numbers, build an array ans of length n such that for each index i:
    - `ans[i] OR (ans[i] + 1) = nums[i]`
    - `ans[i] is as small as possible`
If no such value exists, set `ans[i] = -1`.
- Link: https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/

---

- **Algorithm**: Flip a bit
- After dry running on multiple inputs, we can observe that the required number can be obtained by flipping the leftmost set bit starting from the right.
- For even numbers, the answer is `-1`.
- A brute-force solution with `O(n*k)` time complexity will do the work.

```
0 1 0 1 0 1 1 1 1
          _  <- flip this
```

<br>
<br>
<br>

### 

<br>
<br>
<br>

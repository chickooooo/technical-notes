# Bit Manipulation

<br>
<br>
<br>
<br>
<br>

## Concepts:

### Check if `ith` bit is set.

- If `num & (1 << i) == 0`, then `ith` bit is not set.
- If the value is non-zero, then it is set.

```
num = 22
i = 4

   010110   <-  22
&  010000   <-  1 << 4

   010000   non zero -> set
```

<br>
<br>
<br>

### Check if number is even or odd

- If `num & 1 == 1`, then the number is odd.
- If `num & 1 == 0`, then the number is even.

```
num = 22

   010110   <-  22
&  000001   <-   1

   000000   zero -> even
```

<br>
<br>
<br>

### Set/Unset `ith` bit

- To set/unset `ith` bit, do XOR of `num` and `(1 << i)`.

```
num = 22
i = 3

   010110   <-  22
^  001000   <-  1 << 3

   011110   num with ith bit set
```

<br>
<br>
<br>

### Unset rightmost set bit

- `n & (n-1)` unsets the rightmost set bit of `n`.

```
num = 22

   010110   <-  22
&  010101   <-  21

   010100   num with rightmost set bit unset
```

---

- This technique can be used to find the number of set bits in `O(k)` time complexity. Where `k` represents the number of set bits.

<br>
<br>
<br>

### 

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

### Construct minimum bitwise array II

- Problem: Given an array nums of n prime numbers, build an array ans of length n such that for each index i:
    - `ans[i] OR (ans[i] + 1) = nums[i]`
    - `ans[i] is as small as possible`
If no such value exists, set `ans[i] = -1`.
- Link: https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/

---

- **Algorithm**: Flip a bit
- Same approach as problem 1, but using bit manipulation instead of for-loop.
- Finding `jth` unset bit: `num & (1 << j) == 0`.
- Unsetting `(j-1)th` bit: `num ^ (1 << (j - 1))`.

<br>
<br>
<br>

### Bit flips to convert number

- Problem: Given two integers start and goal, return the minimum number of bit flips to convert start to goal. A bit flip of a number x is choosing a bit in the binary representation of x and flipping it from either 0 to 1 or 1 to 0.
- Link: https://leetcode.com/problems/minimum-bit-flips-to-convert-number/

---

- **Algorithm**: Count set bits
- Take XOR of given two numbers to find the mismatching bits.
- All the mismatching bits will be set `1` in the new number.
- Count the number of set bits using **Unset rightmost set bit** technique.
- Return the count of set bits.

<br>
<br>
<br>

### 

<br>
<br>
<br>

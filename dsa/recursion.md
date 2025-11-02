# Recursion

<br>
<br>
<br>
<br>
<br>

## Problems

<br>
<br>
<br>

### Combination sum

- Problem: Find all unique combinations of numbers from a list that sum up to a given target, where each number can be used multiple times.
- Link: https://leetcode.com/problems/combination-sum/

---

- **Algorithm**: `i` iterates over array
- Sort the given array first.
- Start with `index=0` and `total=0`.
- Iterate over the array starting from `i=index`.
- While iterating add the `ith` element in the `temp` array and increase the `total`.
- Keep doing so recursively until `total` is equal to or greater than the target.
    - If equal, make copy of `temp` and add it to `output` array. Return `false`.
    - If greater, return `false`. `false` indicates that the upper loop should not continue.
- After each recursive call, remove the last element from `temp` array.
- At the end of recursive function, return `true`.

---

- Returning `true` & `false` helps with early stoping and avoid unnecessary recursion calls.
- Early stoping can only be implemented if the array is sorted.

<br>
<br>
<br>

### 

<br>
<br>
<br>

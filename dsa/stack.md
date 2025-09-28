## Important Concepts

<br>
<br>
<br>

### Prefix, Postfix, Infix
- Prefix to ANY:
    - Iterate backwards on input.
    - Reverse the output (if needed).

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Infix to postfix conversion
- Use stack to store the operators, as per their priority order.
- Keep on adding operands to the output string.

<br>
<br>
<br>

### Postfix to prefix conversion
- Use stack to store operands.
- When an operator arrives, replace last operand with `{operator}{left}{right}`.

<br>
<br>
<br>

### Postfix to infix conversion
- Use stack to store operands.
- When an operator arrives, replace last operand with `({left}{operator}{right})`.

<br>
<br>
<br>

### Prefix to postfix conversion
- Iterate from the back.
- Use stack to store operands.
- When an operator arrives, replace last operand with `{operator}{left}{right}`.
- Return **reversed** output.

<br>
<br>
<br>

### Prefix to infix conversion
- Iterate from the back.
- Use stack to store operands.
- When an operator arrives, replace last operand with `({left}{operator}{right})`.
- Return the only item present in the stack

<br>
<br>
<br>

### Trapping rain water
- Create a decreasing monotonic stack. This stack will hold `[height, width]`
- When a smaller element appears, add `[element, 1]` to the stack.
- When a bigger number appears:
    - Calculate the water limit `min(stack[0][0], element)`.
    - Pop each pair from stack where `stack[-1][0] <= element`. Calculate the water stored.
    - At the same time, keep updating the `width` of the current element.
    - At last add `[element, width]` to the stack.

---

- O(1) space complexity approach
 
<br>
<br>
<br>

### Asteroid collision

- Problem: Given an array of integers representing asteroids moving in space (positive = right, negative = left), simulate their collisions. When a right-moving and a left-moving asteroid collide, the smaller one explodes (or both if equal). Return the final state of the asteroids after all collisions.

---

- **Algorithm**: Monotonic stack.
- Iterate through the array, if an asteroid is moving right, add it to the `stack`.
- If an asteroid is moving left, perform the collisions.

<br>
<br>
<br>

### Sum of subarray minimums

- Problem: Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr. Since the answer may be large, return the answer modulo 10^9 + 7.
- Link: https://leetcode.com/problems/sum-of-subarray-minimums/description/

---

- **Algorithm**: Sliding window + Monotonic stack
- Create a stack and `prefix` that holds the sum of elements in stack.
- While iterating through the array if `num >= stack[-1]`, add `(item, 1)` to the stack and update `prefix`.
- If `num <= stack[-1]`, remove all small or equal items from stack.
- While removing, note the count `freq` of items removed. Also remove the sum from `prefix`.
- At last, add `(item, freq)` to the stack and update `prefix`.
- At each step, add `prefix` to the `total` sum.

---

Why modulo (10^9 + 7) ?
- Large numbers can lead to integer overflow. To avoid that modulo is taken.
- The number 10^9 + 7 (i.e., 1000000007) is a large prime.
- Doing modulo retains the original number that are less than it.

<br>
<br>
<br>

### Sum of subarray ranges

- Problem: You are given an integer array nums. The range of a subarray of nums is the difference between the largest and smallest element in the subarray. Return the sum of all subarray ranges of nums.
- Link: https://leetcode.com/problems/sum-of-subarray-ranges/description/

---

- **Algorithm**: Sliding window + Monotonic stack
- Calculate sum of subarray minimums: [Sum of subarray minimums](#sum-of-subarray-minimums).
- Similarly, calculate sum of subarray maximums.
- Subarray range sum = Maximum sum - Minimum sum.
- `∑(ai​-bi​) = ∑ai ​- ∑bi`

<br>
<br>
<br>

### Smallest integer after removing k digits

- Problem: Given string num representing a non-negative integer num, and an integer k, return the smallest possible integer after removing k digits from num.
- Link: https://leetcode.com/problems/remove-k-digits/

---

- **Algorithm**: Monotonic stack
- Maintain an increasing monotonic stack. Can contain duplicates.
- Remove atmost k elements when creating the stack.
- If less than k elements were removed, remove remaining elements from the end of the stack.
- Handle leading zeros.

### Important Concepts:

<br>
<br>

**Prefix Sum + Hash Map (Subarray Sum Pattern)**

Overview:
- This pattern is used to count the number of subarrays with a given sum `k`.
- Works for arrays with **positive, zero, or negative** numbers.

Key Properties:
- Time Complexity: O(n)
- Space Complexity: O(n) (to store prefix sums)
- Supports negative numbers (unlike sliding window which usually requires non-negative numbers).

Core Idea:
- Maintain a hashmap: `{prefix_sum: count}` to store how often each prefix sum occurs.
- For each element:
    - Update running `prefix_sum`.
    - Check if `prefix_sum - k` exists in the map → this indicates a subarray summing to k.
    - Add the count of `prefix_sum - k` to the answer.
    - Update the map with the current `prefix_sum`.

Use Cases:
- Algorithmic problems:
    - Count subarrays with sum = k
    - Count binary subarrays with sum = k
    - Find max length subarray with sum = k
    - Subarray Sums Divisible by K
- Real-World Applications:
    - Detect periods where cumulative profit/loss equals a target.

<br>
<br>
<br>

### Problems:

<br>
<br>

**Max consecutive ones after k flips**
- Simple variable sliding window problem.
- The window increases while `1` or `k > 0`.
- The window shrinks till we get a flip to move on.
- Handle `k = 0` edge case.
- Use `queue` data structure for `O(1)` shrinking.

<br>

**Longest substring after k replacements**
- Variable sliding window problem.
- Create a frequency array to track frequencies.
- If the current element is `majority` element, increase the window.
- If replacements available, increase the window.
- Otherwise, shrink the window.
- Time complexity: O(n*26)
- Space compelxity: O(26)
---
- Approach 2

<br>

**Longest substring without repeating characters**
- Problem: Given a string `s`, find the length of the longest substring without duplicate characters.
- Simple variable sliding window problem.
- The window increases while `jth` element is not in `hashset`.
- The window decreases while `jth` element is in `hashset`.
- While decreasing, the `ith` element is removed from the hashset.

<br>

**Binary subarrays matching the sum**
- Problem: Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.
- Prefix sum sliding window problem.
- Maintain a hashmap of `{prefix_sum: count}`.
- At each step, count the number of subarrays with sum equal to `prefix - k`.
- Update hashmap for current `prefix_sum`.





<br>
<br>
<br>
<br>
<br>

- Problem: subarray sum greater than k

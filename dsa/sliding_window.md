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

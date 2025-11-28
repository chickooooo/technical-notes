## Important Concepts:

<br>
<br>
<br>

### Prefix Sum + Hash Map (Subarray Sum Pattern)

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

### Inverse sliding window

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Max consecutive ones after k flips
- Simple variable sliding window problem.
- The window increases while `1` or `k > 0`.
- The window shrinks till we get a flip to move on.
- Handle `k = 0` edge case.
- Use `queue` data structure for `O(1)` shrinking.

<br>
<br>
<br>

### Longest substring after k replacements
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
<br>
<br>

### Longest substring without repeating characters
- Problem: Given a string `s`, find the length of the longest substring without duplicate characters.
- Simple variable sliding window problem.
- The window increases while `jth` element is not in `hashset`.
- The window decreases while `jth` element is in `hashset`.
- While decreasing, the `ith` element is removed from the hashset.

<br>
<br>
<br>

### Binary subarrays matching the sum

- Problem: Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

---

- **Algorithm**: Prefix sum + hash map
- Prefix sum sliding window problem.
- Maintain a hashmap of `{prefix_sum: count}`.
- At each step, count the number of subarrays with sum equal to `prefix - k`.
- Update hashmap for current `prefix_sum`.

<br>
<br>
<br>

### Subarrays containing k odd numbers
- Problem: Given an array of integers nums and an integer k. Return the number of continuous subarrays that contains k odd numbers.

---
- **Algorithm**: Simplify the given problem to a simpler problem. Fixed size sliding window.
- Create a new list containing all odd indexes.
- Perform fixed sized (k) sliding window on this list.
- At each step,
    - left = no. of even numbers to the left (after last odd number)
    - right = no. of even numbers to the right (before next odd number)
    - Add all combinations `1 + left + right + left*right` to the total

```
2 2 2 1 2 2 1 2 1 2 2 1 2   |   k=3
      ___________
         window
-----             ---
 left            right
```

---

Approach 2

<br>
<br>
<br>

### Substrings containing a,b,c atleast once

- Problem: Given a string s consisting only of characters a, b and c. Return the number of substrings containing at least one occurrence of all these characters a, b and c.
- Link: https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

---

- **Algorithm**: Variable sliding window + Counting at each step.
- Use 2 pointers `i & j`.
- Increment `j` till you find the first window containing `a,b,c`.
- This window and all the windows to the right (by increasing `j`) are valid. Add this count to `total`.
- Now decrement `i`. If the window is valid, add the same above count to `total`.
- Continue decrement of `i` till the window becomes invalid.
- Again start from 2nd step.

<br>
<br>
<br>

### Count subarrays with given XOR

- Problem: Given an array of integers arr[] and a number k, count the number of subarrays having XOR of their elements as k.
- Link: https://www.geeksforgeeks.org/problems/count-subarray-with-given-xor/1

---

- **Algorithm**: Prefix sum + hash map
- Maintain a hashmap of `{prefix_xor: count}`. Set `hashmap[0] = 1`.
- At each step, perform XOR with current element and update the value of `prefix_xor`.
- Count the number of subarrays with XOR equal to `prefix_xor ^ k`.
- Increment the count of current `prefix_xor` in the hashmap.

<br>
<br>
<br>

### Maximum sum using elements from start / end

- Problem: Given an array cardPoints and an integer k, pick exactly k cards from either the start or end of the array to maximize your total score. Return the maximum score possible.
- Link: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

---

- **Algorithm**: Inverse sliding window
- Calculate `total` sum of the array.
- Create a window of size `n-k`. Subtract the sum of this window from the `total`.
- Slide this window over the array and find the maximum value of total.
- The idea is to maintain a fixed size window that will subtract elements instead of adding them.

<br>
<br>
<br>

### Longest substring with atmost k distinct characters

- Problem: You are given a string 'str' and an integer ‘K’. Your task is to find the length of the largest substring with at most ‘K’ distinct characters.
- Link: https://www.naukri.com/code360/problems/distinct-characters_2221410

---

- **Algorithm**: Variable sliding window + hashmap
- Use a `hashmap` to maintain the distinct elements in the window and their frequency.
- If current element is already in window or `len(hashmap)` is less than `k`, then add the `jth` element to the window and increase the window `j+=1`.
- After each window increase, update the `maximum`.
- If current element is not in window, then `while len(hashmap)==k`, remove the `ith` element from the window and shrink the window `i-=1`.
- At last, return the `maximum`.
- At each point in time, the window will contain atmost k distinct characters.

<br>
<br>
<br>

### Count subarrays with k distinct integers (⭐)

- Problem: Given an integer array nums and an integer k, return the number of good subarrays of nums. A good array is an array where the number of different integers in that array is exactly k.
- Link: https://leetcode.com/problems/subarrays-with-k-different-integers/

---

- **Algorithm**: f(k) = g(k) - g(k-1)
- Count of subarrays with k distinct integers = Count of subarrays with **atmost k** distinct integers - Count of subarrays with **atmost k-1** distinct integers.
- f(k) = Count of subarrays with k distinct integers
- g(k) = Count of subarrays with **atmost k** distinct integers

<br>

Algorithm to count subarrays with k distinct integers.

- **Algorithm**: Variable sliding window + hashmap
- Use a `hashmap` to maintain the distinct elements in the window and their frequency.
- If current element is already in window or `len(hashmap)` is less than `k`, then add the `jth` element to the window and increase the window `j+=1`.
- After each window increase, update the count `count += j - i`.
- If current element is not in window, then `while len(hashmap)==k`, remove the `ith` element from the window and shrink the window `i-=1`.
- At last, return the `count`.
- At each point in time, the window will contain atmost k distinct characters.

---

- Far-near pointer approach

<br>
<br>
<br>

### Maximum subarray sum with length divisible by k

- Problem: You are given an array of integers nums and an integer k. Return the maximum sum of a subarray of nums, such that the size of the subarray is divisible by k.
- Link: https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/

---

- **Algorithm**: Fixed size sliding window + Kadane's algorithm
- Use **Fixed size sliding window** algorithm to find sum of all subarrays of size k.
- For each subarray, check if a trailing subarray of size `k` exists.
    - If it exists and has `sum > 0`, include that sum in the sum of the current subarray.
    - This approach follows **Kadane's algorithm**; include previous sum only if it is positive.
- Keep track of maximum sum while iterating and return it.

<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
<br>

- Problem: subarray sum greater than k

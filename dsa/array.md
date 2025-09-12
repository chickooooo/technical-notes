
### Important Concepts

<br>
<br>

**Boyer-Moore Majority Voting Algorithm**

Overview:
- This algorithm is used to find the majority element in an array.
- The majority element in an array of size `n` is the element that appears more than `n/2` times.
- For example, in `[3, 3, 4, 2, 3, 3, 3]`, the majority element is `3`.

Key Properties:
- Time Complexity: O(n)
- Space Complexity: O(1)
- Requires only a single pass through the array (for candidate selection) and a second optional pass (for verification).

Core Idea:
- The algorithm "cancels out" different elements. If there's a majority element, it will survive this canceling process and emerge as the final candidate.

Use Cases:
- Voting systems.
- Real-time data streams (since it uses O(1) space).

<br>
<br>

**Kadane's Algorithm**

Overview:
- This algorithm is used to find the **maximum sum of a contiguous subarray** within a one-dimensional array of numbers.
- For example, in `[-2, 1, -3, 4, -1, 2, 1, -5, 4]`, the maximum sum is `6`, from subarray `[4, -1, 2, 1]`.

Key Properties:
- Time Complexity: O(n)
- Space Complexity: O(1)
- Requires only a single pass through the array.

Core Idea:
- At each index, the algorithm decides whether to **extend the current subarray** or **start a new subarray** at the current element.
- It maintains two variables:
  - `current_sum`: the maximum sum ending at the current index.
  - `max_sum`: the maximum sum found so far.

Use Cases:
- Financial analysis (e.g., finding the most profitable period).
- Performance metrics over time (e.g., longest streaks).
- Applied in dynamic programming problems involving subarrays or subsegments.

<br>
<br>
<br>

### Problems:

<br>
<br>

**Flatten nested array**
- Recursive approach:
    - While iterating, if a number is found, add it to the output.
    - If a list is found, call the flatten function again.
- Iterative approach (Stack):
    - Simulate stack.
    - For each list encountered, save the current index & new list in stack. `[(A, 1), (B, 2)]`
    - Start again from index 0.
- Iterative approach (Reverse):
    - Reverse the list and store it in a stack.
    - While iterating backwards, if a number is found, add it to the output.
    - If a list is found, reverse it and extend to the stack.

<br>
<br>

**Sort an array of 0's 1's and 2's**
- Problem involves rearranging items inplace.
- Use 2 pointers to swap items and place them at appropriate place.

<br>
<br>

**Stock Buy and Sell once**
- Do normal iteration:
    - If current value is lower, buy at this value.
    - If current value is higher, try to sell at this value.

<br>
<br>

**Majority Element (freq > n/2)**
- Can be done in O(n) time and O(n) space using hashmap.
- Use **Boyer-Moore Majority Voting Algorithm** for O(n) time and O(1) space.
- The core idea is, if we clash majority elements vs non-majority elements, there will always be atleast 1 majority element left.

<br>
<br>

**Find subarray with maximum sum**
- Kadane's algorithm.
- If the current prefix sum is negative, don't include it.
- If it is positive, you can include those elements.
---
- Divide and conquer

## Important Concepts

<br>
<br>
<br>

### Max of min / Min of max

How to identify:
- The problem asks to **maximise the minimum value** or **minimise the maximum value**.
- Cannot be solved using traditional array, greedy or other approaches.

Steps:
- The steps to solve and other details are same as **Binary search on range** problems.
- The crucial part is to identify such problems.

<br>
<br>
<br>

### Binary search on range

How to identify:
- The problem asks for the **minimum or maximum** value of something.
- There exists a range of possible answers with a clear lower and upper bound.
- The problem is about searching for a value that satisfies a constraint.

Steps:
1. Define the lower and upper limit of the search space.
2. Perform binary search on this search space and verify the constraints.
3. Store valid values in a variable and return the optimal value at last.

Key Properties:
- Time Complexity: O(nlogn)
- Space Complexity: O(1)
- Does **not** require the array to be sorted.

Use Cases:
- Problems involving **allocation, scheduling, or distribution** under constraints.
- Finding minimum time, maximum distance, minimum speed, etc.

Note:
- Always use only 2 conditions for moving the left & right pointer.
- It can be: `if a <= b {} else {}` or `if a >= b {} else {}`.
- Don't use `if a == b`, this leads to missing out some scenarios.

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Find element index in sorted array
- Core binary search algorithm.
- Setup:
    - left = 0, right = n-1
    - while left <= right
    - middle = (left + right) // 2

<br>
<br>
<br>

### Find element index / insert position in sorted array
- Find element index using binary search.
- After BS, if element not found, the `left` pointer will point to insert position.

<br>
<br>
<br>

### Search in rotated sorted array with duplicates
- If `nums[left] == nums[middle] == nums[right]`, we cannot decide which half to eliminate.
- So we reduce the search space like this: `left += 1, right -= 1`.
- If `nums[left] != nums[middle] != nums[right]`, we follow same approach as **Search in rotated sorted array**.
- Optimisaton:
    - If `nums[left] == nums[middle]`, then all elements between left & middle pointer (both inclusive) are same. Hence we can eliminate that search space.
    - If `nums[middle] == nums[right]`, then all elements between middle & right pointer (both inclusive) are same. Hence we can eliminate that search space.

<br>
<br>
<br>

### Capacity To Ship Packages Within D Days
- standard template of **Binary search on answers**.
- Define the lower `max(nums)` & upper limit `sum(nums)` . Perform binary search on this limit.
- At each step, iterate the array to find required days.
- Move the pointers depending on required & given days.

<br>
<br>
<br>

### Koko Eating Bananas
- standard template of **Binary search on answers**.
- Define the lower `1` & upper limit `max(nums)`. Perform binary search on this limit.
- At each step, iterate the array to find required hours.
- Move the pointers depending on required & given hours.

<br>
<br>
<br>

### Split into k subarrays such that maximum sum is minimised

- Problem: Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized. Return the minimized largest sum of the split.
- Link: https://leetcode.com/problems/split-array-largest-sum/

---

- Algorithm: Max of min / Min of max
- Define the lower `max(nums)` & upper limit `sum(nums)`. Perform binary search on this limit.
- At each step, iterate the array to find the groups formed.
- Move the pointers depending on required & formed groups.

<br>
<br>
<br>

### Search in matrix with sorted rows
- problem: first element of each row is greater than last element of previous row.
- Perform binary search on rows. Check if the target lies within the first and last element of the middle row.
- If it does, perform standard binary search on that row.

<br>
<br>
<br>

### Search in matrix with sorted rows and columns
- problem: no duplicates. both rows and columns are sorted in ascending order.
- Start from the top right corner. If the target matches, return `True`.
- If the target is greater, we know it can't be on that row, so `i += 1`.
- If the target is smaller, we know it can't be on that column, so `j -= 1`.
- TC: `O(m+n)`

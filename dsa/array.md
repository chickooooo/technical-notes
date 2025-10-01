
## Important Concepts

<br>
<br>
<br>

### Boyer-Moore Majority Voting Algorithm

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
<br>

### Kadane's Algorithm

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

### Merge Intervals

Overview:
- This algorithm deals with **merging overlapping intervals** in a list of intervals.
- Each interval is represented as a pair `[start, end]`.
- For example, given `[[1,3], [2,6], [8,10], [15,18]]`, the merged result is `[[1,6], [8,10], [15,18]]`.

Key Properties:
- Time Complexity: **O(n log n)** (due to sorting).
- Space Complexity: **O(n)** (to store the merged intervals).
- Sorting the intervals by start time simplifies the merging logic.

Core Idea:
- **Sort** all intervals by their starting point.
- Initialize a list with the first interval.
- For each remaining interval:
    - If it **overlaps** with the last interval in the merged list, **merge** them by updating the end time:
    ```py
    last = [
        min(last[0], current[0]),
        max(last[1], current[1]),
    ]
    ```
    - Otherwise, append the current interval to the merged list.

Use Cases:
- Calendar management (e.g., merging meeting times).
- Range compression (e.g., merging IP ranges or numeric ranges).
- Data cleaning (e.g., consolidating overlapping time intervals in logs).

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>

### Flatten nested array
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
<br>

### Sort an array of 0's 1's and 2's
- Problem involves rearranging items inplace.
- Use 2 pointers to swap items and place them at appropriate place.

<br>
<br>
<br>

### Stock Buy and Sell once
- Do normal iteration:
    - If current value is lower, buy at this value.
    - If current value is higher, try to sell at this value.

<br>
<br>
<br>

### Majority Element (freq > n/2)
- Can be done in O(n) time and O(n) space using hashmap.
- Use **Boyer-Moore Majority Voting Algorithm** for O(n) time and O(1) space.
- The core idea is, if we clash majority elements vs non-majority elements, there will always be atleast 1 majority element left.

<br>
<br>
<br>

### Find subarray with maximum sum
- Kadane's algorithm.
- If the current prefix sum is negative, don't include it.
- If it is positive, you can include those elements.
---
- Divide and conquer

<br>
<br>
<br>

### Set matrix row and column zero
- Problem:
    - Given an integer matrix containing positive, negative and zero integers. For each cell containing `0`, set its entire row and column `0`.
    - Using extra space is trivial, do it in constant space.
---
- We will use the first row and first column to identify whether that entire row or column will be 0.
- First check & store if the first row and column contains a zero.
- Next, for each `0` cell, set it's corresponding first row and column position to `0`.
- Then iterate the matrix again. Each cell that has it's first row and column as `0`, set that cell to `0`.
- At last, if first row contained `0` originally, set that row to zero. Same for first column.

<br>
<br>
<br>

### Rearrange items in alternate positive-negative order

- Problem: Given an array nums of even length, arrange the items such that:
    - Every consecutive pair of integers have opposite signs
    - For all integers with the same sign, the order in which they were present in nums is preserved.
    - The rearranged array begins with a positive integer.

---

- **Algorithm**: Two pointers
- Create a new array of same size. Set `pos = 0` and `neg = 1`.
- Iterate through the given array. If element is positive, copy it to the new array at `pos` index and increment `pos` by 2.
- If element is negative, copy it to the new array at `neg` index and increment `neg` by 2.

<br>
<br>
<br>

### Merge overlapping intervals

- Problem: Given an array of intervals where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
- Link: https://leetcode.com/problems/merge-intervals/

---

- **Algorithm**: Merge intervals
- Sort the intervals in ascending order.
- If `start` of the current interval is greater than `end` of the last interval, add it to the array.
- Otherwise, merge last interval and current interval.
```py
last = [
    min(last[0], current[0]),
    max(last[1], current[1]),
]
```

<br>
<br>
<br>

### Merge 2 sorted arrays

- Problem: You are given two integer arrays nums1 (m) and nums2 (n), sorted in non-decreasing order. Merge nums1 and nums2 into a single array sorted in non-decreasing order. There are n zeros at the end of nums1.
- Link: https://leetcode.com/problems/merge-sorted-array/

---

- **Algorithm**: Two pointers
- Set `i=m-1`, `j=n-1` & `k=m+n-1`. `i` will iterate `nums1` and `j` will iterate `nums2`.
- The larger number from `i` or `j` will be replaced with the zero present at `k`.
- Continue this replacement until all number of `nums2` are in `nums1`.
- Handle the edge case where `nums1` finishes before `nums2`.

<br>
<br>
<br>

### Pascal's triangle

- Problem: Given an integer numRows, return the first numRows of Pascal's triangle.
- Link: https://leetcode.com/problems/pascals-triangle/

---

- **Algorithm**: 2 for loops
- The outer for loop decides how many rows to add.
- The inner for loop create a new row by adding neighbouring elements from the previous row.

<br>
<br>
<br>

### Find the repeating and missing number

- Problem: Given an unsorted array arr[] of size n, containing elements from the range 1 to n, it is known that one number in this range is missing, and another number occurs twice in the array, find both the duplicate number and the missing number.
- Link: https://www.geeksforgeeks.org/problems/find-missing-and-repeating2512/1

---

- **Algorithm**: Using array index
- Iterate through the array.
- In the first iteration, for all `nums`, mark `(num-1)th` element in array as negative.
- While marking a number negative, if that number is already negative, that means `num` is duplicate.
- Iterate through the array again and find the number that is not negative.
- This number was not marked, that means `index+1` number is missing from the array.
- Return duplicate and missing number.

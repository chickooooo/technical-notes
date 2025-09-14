### Important Concepts

<br>
<br>

**Binary search on answers**

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

<br>
<br>
<br>

### Problems:

<br>
<br>

**Find element index in sorted array**
- Core binary search algorithm.
- Setup:
    - left = 0, right = n-1
    - while left <= right
    - middle = (left + right) // 2

<br>
<br>

**Find element index / insert position in sorted array**
- Find element index using binary search.
- After BS, if element not found, the `left` pointer will point to insert position.

<br>
<br>

**Search in rotated sorted array**
- [PLACE HOLDER]

<br>
<br>

**Search in rotated sorted array with duplicates**
- If `nums[left] == nums[middle] == nums[right]`, we cannot decide which half to eliminate.
- So we reduce the search space like this: `left += 1, right -= 1`.
- If `nums[left] != nums[middle] != nums[right]`, we follow same approach as **Search in rotated sorted array**.
- Optimisaton:
    - If `nums[left] == nums[middle]`, then all elements between left & middle pointer (both inclusive) are same. Hence we can eliminate that search space.
    - If `nums[middle] == nums[right]`, then all elements between middle & right pointer (both inclusive) are same. Hence we can eliminate that search space.

<br>
<br>

**Capacity To Ship Packages Within D Days**
- standard template of **Binary search on answers**.
- Define the lower & upper limit. Perform binary search on this limit.
- At each step, iterate the array to find required days.
- Move the pointers depending on required & given days.

<br>
<br>

**Search in matrix with sorted rows**
- problem: first element of each row is greater than last element of previous row.
- Perform binary search on rows. Check if the target lies within the first and last element of the middle row.
- If it does, perform standard binary search on that row.

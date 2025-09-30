## Important Concepts

<br>
<br>
<br>

### Binary search on range

How to identify:
- The problem asks for the **minimum or maximum** value of something.
- There exists a range of possible answers with a clear lower and upper bound.
- The problem is about searching for a value that satisfies a constraint.

---
- Note: Brute-force solution for such problems will give `O(range*n)` complexity.
- The range will be sorted, so instead of normal linear iteration, we can perform binary search to reduce the time complexity to `O(log(range)*n)`. 

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

### Binary search on rotated sorted array

<br>
<br>
<br>

### Binary search for decimal

How to identify:
- The problem requires finding a **real number (floating point)** answer, not just an integer.
- This is similar to binary search on integers but adapted to handle **floating-point precision**.
- Because floating point numbers are infinite within any range, we stop the search based on a **precision threshold**, not just when `left <= right`.

Steps:

1. Define the **lower and upper bounds** of the search space.
2. Choose a **precision level** (`e.g. 1e-6` if answer is required up to 5 decimal places).
3. Perform binary search using:
   `while right - left > precision:`
4. Check the **condition at mid**, and move either `left` or `right` accordingly.

Key Properties:

- Time Complexity: `O(log(range * 10^decimal_places))`
- Space Complexity: `O(1)`
- Works on **continuous ranges**, not just integers.

Pointer Movement:

```python
if condition(mid):
    left = mid
else:
    right = mid
```

Use Cases:

- Finding **square root**, **cube root**, or any **root of a monotonic function**.
- **Minimum time**, **distance**, or **rate** in physics or optimization problems.
- Problems involving **probability thresholds**, **geometry (e.g., circle/radius fitting)**, etc.

Note:

- Always set `precision = 1e-(p+1)` if the problem asks for `p` decimal places of accuracy.
- In floating point problems, use caution with comparison (avoid `==`).
- Final answer should often be **rounded** to the required decimal precision.

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Kth missing positive number

- Problem: Given an array arr of positive integers sorted in a strictly increasing order, and an integer k. Return the kth positive integer that is missing from this array.
- Link: https://leetcode.com/problems/kth-missing-positive-number/

---

- **Algorithm**: Standard binary search
- Perform standard binary search.
- At each step, calculate the count of missing numbers towards the left of middle. `missing_count = arr[index] - (index + 1)`.
- If missing count is less than `k`, eliminate left search space.
- If missing count is `>= k`, set `upper_bound = middle` and eliminate right search space.
- After binary search, if `upper_bound == -1`, that means the missing element is present after the last element of array: `missing_element = len(arr) + k`.
- Otherwise, `missing_element = arr[upper_bound] - (missing_count - k) - 1`. This simplifies to: `missing_element = upper_bound + k`

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

### Find minimum in rotated sorted array

- Problem: Given a sorted rotated array nums of unique elements, return the minimum element of this array. You must write an algorithm that runs in O(log n) time.
- Link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

---

- **Algorithm**: Binary search on rotated sorted array
- Check which part is sorted using the standard **Binary search on rotated sorted array** algorithm.
- The minimum element can be the starting element of this sorted part.
- Use binary search to shrink the search space and keep on updating the minimum element.

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

### Minimum Number of Days to Make m Bouquets

- Problem: You are given an integer array bloomDay, an integer m and an integer k. You want to make m bouquets. To make a bouquet, you need to use k adjacent flowers from the garden. The garden consists of n flowers, the ith flower will bloom in the `bloomDay[i]` and then can be used in exactly one bouquet. Return the minimum number of days you need to wait to be able to make m bouquets from the garden. If it is impossible to make m bouquets return -1.
- Link: https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/

---

- **Algorithm**: Binary search on range.
- Define the lower `1` & upper limit `max(nums)`. Perform binary search on this limit.
- At each step, iterate the array to find the bouquets made.
- Move the pointers depending on required & made bouquets.

<br>
<br>
<br>

### Split into k subarrays such that maximum sum is minimised

- Problem: Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized. Return the minimized largest sum of the split.
- Link: https://leetcode.com/problems/split-array-largest-sum/

---

- **Algorithm**: Max of min / Min of max
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

<br>
<br>
<br>

### ⭐ Minimize max distance to gas station

- Problem: You are given a sorted list of gas station positions on a number line. You can add k more stations to minimize the maximum distance between any two adjacent stations. The goal is to find the smallest possible maximum distance after adding the stations, rounded to 2 decimal places.
- Link: https://www.geeksforgeeks.org/problems/minimize-max-distance-to-gas-station/1

---

- **Algorithm**: Max of min / Min of max
- Define the lower limit: `0` and upper limit: `max(stations[i]-stations[i-1])`.
- Define the decimal precision (always required precision + 1): `precision = 1e-3`.
- Perform **Binary search for decimal**.
- At each step, iterate the array to find the stations required such that the max distance between adjacent stations is `middle`.
- Number of stations required between 2 adjacent stations: `math.ceil((stations[i]-stations[i-1]) / max_distance) - 1`
- Move the pointers depending on required & given stations.

### Important Concepts

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

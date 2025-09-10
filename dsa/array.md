
### Important Concepts

<br>
<br>

**Boyer-Moore Majority Voting Algorithm**

Overview
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

### Problems:

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

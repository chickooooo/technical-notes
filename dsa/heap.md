## Important Concepts

<br>
<br>
<br>
<br>
<br>

## Problems:

<br>
<br>
<br>

### Kth largest element in an array

- Problem: Given an integer array nums and an integer k, return the kth largest element in the array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
- Link: https://leetcode.com/problems/kth-largest-element-in-an-array/

---

- **Algorithm**: Kth largest/smallest element
- Maintain a `heap` of size `k`.
- For each element in the array, if `len(heap) < k`, push that element into the heap.
- Otherwise, push and pop that element to and from the heap.
- Note that the size of the `heap` should not exceed k.

<br>
<br>
<br>

### 

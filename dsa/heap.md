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

### Form pairs of k consecutive numbers

- Problem: Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size groupSize, and consists of groupSize consecutive cards. Given an integer array hand where `hand[i]` is the value written on the ith card and an integer groupSize, return true if she can rearrange the cards, or false otherwise.
- Link: https://leetcode.com/problems/hand-of-straights/

---

- **Algorithm**: Frequency heap
- Create a heap that holds `[element, frequency]`.
- Use `Counter()` from `collections` to populate this heap.
- While heap is not empty, pop `groupSize` elements from the heap at a time. If less than `groupSize` elements exist, return `False`.
- Verify these elements are consecutive. If not, return `False`.
- Next, decrement the `frequency` of extracted elements by `1` and add non-zero frequency elements back to the heap. 
- Keep doing so until the heap is empty.
- If heap is emptied, return `True`.

<br>
<br>
<br>

### Top k frequent elements

- Problem: Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
- Link: https://leetcode.com/problems/top-k-frequent-elements/

---

- **Algorithm**: Top k / bottom k elements
- Maintain a `heap` of size k containing `(freq, num)`.
- Iterate through `Counter(nums)` and if `len(heap) < k`, keep on pushing data into the heap.
- If `len(heap) == k`, push and pop data from the heap.
- At last, return `num` from the last k data items present in the heap.

<br>
<br>
<br>

### 

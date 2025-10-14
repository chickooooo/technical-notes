# Linked List

<br>
<br>
<br>
<br>
<br>

## Problems

<br>
<br>
<br>

### Palindrome linked list

- Problem: Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
- Link: https://leetcode.com/problems/palindrome-linked-list/

---

- **Algorithm**: Slow fast pointer + Reverse linked list
- Use **Slow fast pointer** technique to get the middle of the linked list.
- Use **Reverse linked list** technique to reverse the right half of the linked list.
- Use 2 pointers pointing to the start of both the halves.
- Iterate through the linked lists and check for palindrome.

<br>
<br>
<br>

### Group alternate nodes

- Problem: Given the head of a singly linked list, rearrange the nodes so that all nodes at odd indices come first, followed by nodes at even indices. Maintain the relative order within the odd and even groups. Return the reordered list. The first node is at index 1 (odd), the second at index 2 (even), and so on.
- Link: https://leetcode.com/problems/odd-even-linked-list/

---

- **Algorithm**: Two pointers
- Handle no node condition.
- Set `odd` pointer to first node (`o1`) and `even` pointer to second node (`e1`).
- While `even` and `even.next` exists, set `o1 -> o2`, `o2 -> e1`, `e1 -> e2`.
- Then, increment `odd` and `even` pointers.
- Make sure, no cycle is created in the linked list.

```
      - - - - - - -
      |           |
- - - - - - -     |
|     |     ↓     ↓
o1 -> e1 -> o2 -> e2
      ↑     |
      - - - -
```

<br>
<br>
<br>

### 

<br>
<br>
<br>

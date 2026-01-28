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

### Delete Node in a Linked List

- Problem: You’re given a node (not the last) in a singly linked list with all nodes having unique value. You must delete the given node without access to the head node.
- Link: https://leetcode.com/problems/delete-node-in-a-linked-list/

---

- **Algorithm**: Alter node values
- We want to remove current node, but we do not know what the previous node is.
- Instead we have pointer to the current node and current node is not the last node.
- So, we will copy the value of next node to current node and remove the next node.
- All the values are unique and the problem explicitly tells it is okay to not remove the node from the memory.

<br>
<br>
<br>

### Reverse linked list

- Problem: Given the head of a singly linked list, reverse the list, and return the reversed list.
- Link: https://leetcode.com/problems/reverse-linked-list/

---

- **Algorithm**: Reverse linked list (iterative)
- Set previous node to `nil` initially.
- Make current node point to previous node.
- Move current by one step and previous by one step.
- Keep doing so untils current becomes `nil`.
- At last, return the previous node. This is the new head now.

---

- **Algorithm**: Reverse linked list (recursive)
- Recursively call `current.Next` until we reach the last node.
- If current node is last node, set `newHead` to this node and return the node.
- After the recursive call, set received node's next to current node and return the current node.
- After the class stack is emptied, first node is still pointing to second node. So, set `first.Next = nil`.
- At last, return the `newHead`.

<br>
<br>
<br>

### Sort linked list

- Problem: Given the head of a linked list, return the list after sorting it in ascending order.
- Link: https://leetcode.com/problems/sort-list/

---

- **Algorithm**: Merge sort
- Use **Slow & Fast pointer** algorithm to get the middle of the linked list.
- Split linked list into 2 parts from the middle.
- Recursively keep on splitting each part until a single node is left.
- Now get the left sorted part and right sorted part and merge these two parts.
- The algorithm is same as **Merge two sorted linked lists**. No need for extra space.
- Keep doing so until we get a single sorted linked list.

<br>
<br>
<br>

### Delete the middle node of a linked list

- Problem: You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
- Link: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

---

- **Algorithm**: Slow fast pointer
- Create `slow=head`, `fast=head` & `prev=None` pointers. 
- Iterate through the linkedlist using slow-fast algorithm till the `fast` pointer reaches the end.
- After the iteration is complete, `slow` pointer will point to the middle node.
- Use `prev` pointer to remove the middle node and return the `head` of LL.

<br>
<br>
<br>

### LRU Cache

- Problem: Design an LRU (Least Recently Used) cache that supports get and put operations in O(1) average time.
- Link: https://leetcode.com/problems/lru-cache/

---

- **Algorithm**: Doubly linked list + hashmap
- Use a doubly linkedlist to remove a node from any position in `O(1)`.
- It also removes last node and inserts a new node at start in `O(1)`.
- Use a hashmap of `{key: node}`, to quickly find the node associated with that key.

<br>
<br>
<br>

### LFU Cache (⭐)

- Problem: Implement LFU cache.
- Link: https://leetcode.com/problems/lfu-cache/

---

- Use doubly linkedlist to store items by descending frequency.
- Use `{key:node}` for direct access of node using the key.
- Use `{freq:node}` to keep track of starting node for each frequency.
- Whenever a node is accessed / updated, increase it's frequency and update it's place in the doubly LL.
- Intuition is easy, implementation is hard due to edge cases handling.

<br>
<br>
<br>

### Detect loop in linked list

- Problem: Given head, the head of a linked list, determine if the linked list has a cycle in it.
- Link: https://leetcode.com/problems/linked-list-cycle/

---

- **Algorithm**: Slow fast pointer
- Start with `slow` pointer pointing to the head and fast pointer pointing to `head.next`.
- Move `slow` pointer by 1 step and `fast` pointer by 2 steps.
- If `fast` reaches the end, there is no loop.
- If there is a loop, eventually `slow` will be equal to `fast`.

---

- **Approach 2**: Use a `set` to keep track of visited nodes.
- While iterating if that node is present in the `set`, that means there is a loop.

<br>
<br>
<br>

### Starting point of linked list cycle

- Problem: Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
- Link: https://leetcode.com/problems/linked-list-cycle-ii/

---

- **Algorithm**: Slow fast pointer
- Start with `slow` pointer pointing to the head and fast pointer pointing to `head.next`.
- Move `slow` pointer by 1 step and `fast` pointer by 2 steps.
- If `fast` reaches the end, there is no loop. Otherwise, eventually `slow` will be equal to `fast`.
- Now, create 2 new pointers `new_start = slow.next` and `start = head`.
- Move these 2 pointers together by 1 step until they meet. The node they meet on is the starting node of the cycle.

---

- **Approach 2**: Use a `set` to keep track of visited nodes.
- Return the first duplicate node.

<br>
<br>
<br>

### 

<br>
<br>
<br>

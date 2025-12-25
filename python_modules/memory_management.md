# Python's Memory Management

<br>
<br>
<br>
<br>
<br>

## Index

- [Stack & Heap](#stack--heap)

<br>
<br>
<br>
<br>
<br>

## Stack & Heap

### Stack

- A stack is used to store:
    - Function call frames
    - Local variables
    - References to heap objects
- Operations on the stack happens in LIFO (last in first out) manner.
- Each function call get its own stack frame.

<br>
<br>

### Heap

- A heap is used to store the actual python objects.
- This contains both mutable (`list`, `dict`, `set`) and immutable (`int`, `str`, `bool`) objects.
- Objects remain in the heap as long as they are **referenced**.

<br>
<br>

### Example

```py
def add(a, b):
    c = a + b
    return c

x = 10
y = 20
z = add(x, y)
```

<br>

![Stack Heap](./stack_heap.png)

<br>

- `10` & `20` objects are created in heap.
- `x` & `y` contain reference to these objects and are stored in stack.
- `add()` function call creates a new stack frame.
- `a` & `b` parameters are pointing to the same `10` & `20` objects in the heap. **Python Interning**.
- After addition, a new object `30` is added to the heap.
- `c` is stored in stack which points to this `30`.
- After function returns, `z` is now holding reference to this `30` object.
- And the function frame is removed from the stack.
- At the end of the program, remaining references and objects are also removed from the stack and heap respectively.

<br>
<br>

### Key points

- Python uses stack for references & function calls.
- Python uses heap for all objects.
- Variables are references, not containers.
- Stack frames are destroyed after function returns.
- Heap objects live as long as their reference count is greater than 0.

<br>
<br>
<br>
<br>
<br>

## 

<br>
<br>
<br>
<br>
<br>

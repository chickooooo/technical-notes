## Index
- [Python Context Manager](#python-context-manager)
- [Python List vs Array](#python-list-vs-array)
- [Python Module vs Package](#python-module-vs-package)
- [Python Memory Management Model](#python-memory-management-model)
- [Python Asyncio package](#python-asyncio-package)
- [Python Multithreading vs Multiprocessing vs Asyncio](#python-multithreading-vs-multiprocessing-vs-asyncio)
- [Python *args vs **kwargs](#python-args-vs-kwargs)
- [Python Unpacking operators](#python-unpacking-operators)

<br>
<br>
<br>

### Python Context Manager

- Context Managers are python class that implement `__enter__` and `__exit__` methods.
- These methods are used to setup & cleanup the resource.
- They are primarily used with the with block in python.
- If an error occurs, returning False from `__exit__` function will propogate the error.
- If an error occurs, returning True from `__exit__` function will suppress the error.

<br>

```python
"""Custom Context Manager"""

class MyContextManager:
    def __enter__(self):
        print("enter")
        return self

    def __exit__(self, exc_type, exc_val, exc_tb) -> bool:
        print("exit")
        if exc_type:
            print("error")
            return False  # propogate error
        return True  # suppress error


with MyContextManager() as manager:
    print("inside")
    raise ValueError("kkk")

# Output:

# enter
# inside
# exit
# error
# Traceback (most recent call last):
#   File "/home/a/Desktop/projects/other/python_notes/python_code.py", line 17, in <module>
#     raise ValueError("kkk")
# ValueError: kkk
```

<br>

- We can also create function-based context managers using `contextlib` module.
- Use `@contextmanager` decorator to create a context manager function.
- The `yield` keyword divides the enter and exit secions of the context manager.

<br>

```python
"""Using contextlib"""

from contextlib import contextmanager

@contextmanager
def my_context_manager():
    print("enter")
    try:
        yield  # This is where the block inside `with` runs
    except Exception as e:
        print("error")
        raise e
    finally:
        print("exit")


with my_context_manager() as manager:
    print("inside")
    raise ValueError("kkk")

# Output:

# enter
# inside
# error
# exit
# Traceback (most recent call last):
#   File "/home/a/Desktop/projects/other/python_notes/python_code.py", line 18, in <module>
#     raise ValueError("kkk")
# ValueError: kkk

```

<br>
<br>
<br>

### Python List vs Array

List
- Ordered collection of items.
- Dynamic sizing. Can add and remove items.
- Can hold items of different datatypes.
- Slower than arrays as they are general purpose containers.

<br>

```python
my_list = [1, 2.0, "3", "four"]

print(my_list)  # [1, 2.0, '3', 'four']
```

<br>

Array
- Ordered collection of items.
- Dynamic sizing. Can add and remove items.
- Can hold items of same datatype.
- Faster than lists as they hold items of same datatype.
- Efficient memory usage and performance.

<br>

```python
from array import array

int_array = array("i", [1, 2, 3, 4])
float_array = array("f", [1.0, 2, 3.0, 4.0])

int_array.append(5)
int_array.remove(1)
float_array.pop()

print(int_array)  # array('i', [2, 3, 4, 5])
print(float_array)  # array('f', [1.0, 2.0, 3.0])
```

<br>

Prefer `List` when you want a flexible ordered collection of items.<br>
Prefer `Array` when you want to store large data of same type.

<br>
<br>
<br>

### Python Module vs Package

Module
- A module is simply a `.py` file.
- It can contain functions, classes, or variables.
- You can import and use a module in other Python files.

Package
- A package is a directory containing multiple Python files (modules).
- This directory must contain an `__init__.py` file. Even if it is empty.
- This `__init__.py` file differentiates a directory from a package.
- A package can contain sub-packages.

<br>
<br>
<br>

### Python Memory Management Model

Python's memory management model contains 4 main sections:
- Memory Allocation
- Reference Counting
- Garbage Collection
- Memory Pools

<br>

**Memory Allocation**
- In Python, memory is managed automatically.
- Developer don't have to manually allocate and deallocate memory to objects.
- Memory is allocated dynamically when an object is created.
- When an object is created, it is stored in the heap and a reference to that object is stored in a stack.

<br>

**Reference Counting**
- When a new reference to an object is made, it's reference count is increased by 1.
- When an existing reference goes out of scope, it's reference count is decreased by 1.
- When reference count of an object reaches 0, the object is removed from the memory.

<br>

**Garbage Collection**
- Reference counting handles most cases, but it can't handle cyclic references.
- Here, the garbage collector detects and breaks these cycles.
- Once the cyclic references are broken, the reference counter works as expected and deallocates the memory when the counter reaches 0.
- The garbage collector is triggered when certain thresholds are reached in terms of the number of objects and memory allocated.
- It runs in the background without affecting the program.

<br>

**Memory Pool**
- Python uses memory pool to allocate memory in chunks.
- This avoids frequent memory allocation & deallocation requests to the OS.
- For small objects, python uses it's own memory allocator `pymalloc` (python memory allocator) that works in memory blocks.
- For large objects, the allocator requests memory from the OS.

<br>
<br>
<br>

### Python Asyncio package

- Asyncio package is useful when dealing with asynchronous I/O bound operations.
- It uses event loop to schedule and run tasks.
- Event loop is also responsible for switching between tasks, when `await` keyword is encountered.
- Asyncio has a single thread and a single process design.

<br>

**Key Components**

- `async def`: Creates a coroutine.
- `await`: Pauses the coroutine until the awaited task finishes.
- Event loop: Manages and runs all the asynchronous tasks.
- While a coroutine is paused, event loop will start executing another coroutine.

<br>

```python
import asyncio

async def say_hello():
    print("Hello...")
    await asyncio.sleep(1)
    print("...world!")

async def main():
    await asyncio.gather(say_hello(), say_hello())

asyncio.run(main())
```

<br>
<br>
<br>

### Python Multithreading vs Multiprocessing vs Asyncio

**Multithreading**
- Multithreading spwans multiple threads in a same process.
- Ideal for blocking I/O bound tasks.
- Due to GIL, only 1 thread is executed by the python interpreter at a time.
- While one thread is waiting for I/O bound operation, another thread is picked for execution.
- OS is responsible for context switching.
- Use threading when you want to perform `blocking I/O operations` concurrently.
- Avoid using threads for CPU intensive tasks.

<br>

**Multiprocessing**
- Multiprocessing spwans multiple processes, each with it's own python interpreter and memory.
- Ideal for CPU bound tasks.
- Each process has it's own GIL, hence true parallelism is possible.
- Use when you want to do parallel computation across CPU cores.
- Multiprocessing has higher overhead due to inter-process communication and memory usage.

<br>

**Asyncio**
- Asyncio works with a single thread and a single process.
- Ideal for non-blocking I/O bound tasks.
- It uses event loop to manage and schedule asynchronous I/O tasks.
- When one task encounters `await` statement, another task is picked for execution.
- Developer is in control of context switching (using `await` keyword).
- Use asyncio when you want to perform `non-blocking I/O operations` concurrently.

<br>
<br>
<br>

### Python *args vs **kwargs

- Both `*args` & `**kwargs` are used to pass arbitrary number of arguments to a function.
- `*args` collect extra positional arguments as a tuple.
- `**kwargs` collect extra keyword arguments as a dictionary.

<br>

```python
# extra positional arguments

def abc(*args):
    print(type(args))
    print(args)

abc(1, 2, "ooo")

# Output:
# <class 'tuple'>
# (1, 2, 'ooo')
```

<br>

```python
# extra keyword arguments

def abc(**kwargs):
    print(type(kwargs))
    print(kwargs)

abc(a=1, b="ooo")

# Output:
# <class 'dict'>
# {'a': 1, 'b': 'ooo'}

```

<br>

```python
# extra positional + keyword arguments

def abc(*args, **kwargs):
    print(args)
    print(kwargs)

abc(1, 2, a="ooo")

# Output:
# (1, 2)
# {'a': 'ooo'}

```

<br>
<br>
<br>

### Python Unpacking operators

- In python `*` and `**` are often refered to as unpacking operators.
- `*` is used to unpack an iterable like a list or a tuple.
- `**` is used to unpack a dictionary like object.
- `*` and `**` are used for unpacking arguments in a function call.

<br>

```python
def greet(name, age):
    print(f"{name} is {age} years old.")

args = ("Alice", 30)
kwargs = {"name": "Bob", "age": 25}

greet(*args)  # Output: Alice is 30 years old.
greet(**kwargs)  # Output: Bob is 25 years old.
```

<br>

```python
# unpacking dictionary like object

from collections import OrderedDict

def greet(name, age):
    print(f"{name} is {age} years old.")

kwargs = OrderedDict({"name": "Bob", "age": 25})

greet(**kwargs)  # Output: Bob is 25 years old.
```

## Index
- [Python Decorators](#python-decorators)
- [Python Context Manager](#python-context-manager)
- [Python List vs Array](#python-list-vs-array)
- [Python Module vs Package](#python-module-vs-package)
- [Python Memory Management Model](#python-memory-management-model)
- [Python Asyncio package](#python-asyncio-package)
- [Python Multithreading vs Multiprocessing vs Asyncio](#python-multithreading-vs-multiprocessing-vs-asyncio)
- [Python *args vs **kwargs](#python-args-vs-kwargs)
- [Python Unpacking operators](#python-unpacking-operators)
- [Python Namespace and Scope](#python-namespace-and-scope)
- [Python Scope Resolution](#python-scope-resolution)
- [Python \_\_init\_\_ vs \_\_new\_\_](#python-\_\_init\_\_-vs-\_\_new\_\_)
- [Python Lambda Function](#python-lambda-function)
- [Python Deep Copy vs Shallow copy](#python-deep-copy-vs-shallow-copy)
- [Python Generators](#python-generators)
- [Python Iterator](#python-iterator)
- [Generator vs Iterator](#generator-vs-iterator)
- [Python Iterables](#python-iterables)
- [Python Exception Handling](#python-exception-handling)
- [Python GIL](#python-gil)
- [Python == vs is](#python--vs-is)
- [Python \_\_str\_\_ vs \_\_repr\_\_](#python-\_\_str\_\_-vs-\_\_repr\_\_)
- [Python with Statement](#python-with-statement)
- [Python Interning](#python-interning)
- [Python Metaclass](#python-metaclass)
- [Python Built-in Functions](#python-built-in-functions)
- [Python Dunder Methods](#python-dunder-methods)
- [Python Duck Typing](#python-duck-typing)
- [String Interpolation](#string-interpolation)
- [`match-case` Statement](#match-case-statement)
- [Python `__slots__`](#python-__slots__)
- [Python Descriptors](#python-descriptors)

<br>
<br>
<br>

### Python Decorators

- Decorators are just functions that return functions.
- They are used to add additional functionality to existing functions.
- For example, `@authenticated` decorator can be used to check if the user is authenticated before passing the request to the view.
- To implement this, create `authenticated()` method with authentication logic and add `@authenticated` above your API view function.
A decorator can also take arguments. The syntax looks like this: `@authenticated(admin_only=True)`.


<br>

```py
# Normal decorator

def my_decorator(func):
    def inner():
        print("before")
        func()
        print("after")
    return inner

@my_decorator
def my_func():
    print("Hello")

my_func()

# Output:
# before
# Hello
# after
```

```py
# Overridden function accepts parameters

def my_decorator(func):
    def inner(*args, **kwargs):
        print("before")
        func(*args, **kwargs)
        print("after")
    return inner

@my_decorator
def my_func(name: str):
    print("Hello", name)

my_func("john")

# Output:
# before
# Hello john
# after
```

```py
# Overridden function return values

def my_decorator(func):
    def inner(*args, **kwargs):
        print("before")
        res = func(*args, **kwargs)
        print("after")
        return res
    return inner

@my_decorator
def my_func(name: str):
    print("Hello", name)
    return 1

print(my_func("john"))

# Output:
# before
# Hello john
# after
# 1
```

```py
# Decorator accepts parameters

def my_decorator(num):
    def wrapper(func):
        def inner(*args, **kwargs):
            print("before", num)
            res = func(*args, **kwargs)
            print("after", num)
            return res
        return inner
    return wrapper

@my_decorator(10)
def my_func(name: str):
    print("Hello", name)
    return 1

print(my_func("john"))

# Output:
# before 10
# Hello john
# after 10
# 1
```

<br>
<br>
<br>

### Python Context Manager

- Context Managers are python class that implement `__enter__` and `__exit__` methods.
- These methods are used to setup & cleanup the resource.
- They are primarily used with the `with` block in python.
- If an error occurs, returning False from `__exit__` function will propogate the error.
- If an error occurs, returning True from `__exit__` function will suppress the error.

<br>

```py
class MyCM:
    def __init__(self, file: str) -> None:
        logger.info("init")
        self._file = file

    def __enter__(self):
        logger.info("opening file: %s", self._file)
        return self

    def __exit__(self, exc_type, exc_val, exc_tb) -> bool:
        logger.info("closing file: %s", self._file)
        # Suppress value error
        if exc_type is ValueError:
            return True
        # Propogate other errors
        return False

    def read(self):
        logger.info("reading file: %s", self._file)

with MyCM("Action") as manager:
    time.sleep(1)
    manager.read()
    time.sleep(1)

# Output:
# 00:00:00 - init
# 00:00:00 - opening file: Action
# 00:00:01 - reading file: Action
# 00:00:02 - closing file: Action
```

<br>

- We can also create function-based context managers using `contextlib` module.
- Use `@contextmanager` decorator to create a context manager function.
- The `yield` keyword divides the enter and exit sections of the context manager.

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

<br>
<br>
<br>

### Python Namespace and Scope

- A namespace is like a dictionary that holds mapping between a variable identifier and its value.
- There are 4 categories of namespace:
    - **Built-in namespace**: contains built-in variables & functions (`print()`, `None`, etc.)
    - **Global namespace**: contains module level variables & functions.
    - **Local namespace**: contains variables defined inside a function.
    - **Enclosing namespace**: refers to namespace of outer function (in case of nested functions).

<br>

```python
x = "global"

def outer():
    x = "enclosing"
    def inner():
        x = "local"
    inner()
```

<br>

- Scope defines visibility and lifespan of a variable.
- It also states where we can access a variable in the code.
- Again, there are 4 categories of scope:
    - **Built-in scope**: Python’s built-in names.
    - **Global scope**: Module level scope.
    - **Enclosing scope**: Outer function scope.
    - **Local scope**: Function level scope. 

- We can see the identifiers defined in global and local namespace using: `globals()` and `locals()` functions.

<br>

```python
my_var_1 = 100
print("Globals:", globals())

def abc():
    my_var_2 = 200
    print("Globals:", globals())
    print("Locals:", locals())

abc()

# Output:

# Globals: {..., 'my_var_1': 100}

# Globals: {..., 'my_var_1': 100, 'abc': <function abc at 0x7efa454231a0>}
# Locals: {'my_var_2': 200}
```

<br>
<br>
<br>

### Python Scope Resolution

- It refers to the process of finding a particular variable in the program.
- Scope resolution is based on the LEGB rule.
- Python first tries to find the variable in the **L**ocal namespace then **E**nclosing namespace then **G**lobal namespace and at last **B**uilt-in namespace.
- If the variable in not found after searching in all 4 namespaces, it raises `NameError`.

<br>

```python
x = "global"

def outer():
    x = "enclosing"

    def inner():
        x = "local"
        print(x)  # Prints "local"
    
    inner()
    print(x)  # Prints "enclosing"

outer()
print(x)  # Prints "global"
```

<br>
<br>
<br>

### Python \_\_init\_\_ vs \_\_new\_\_

- Both of these dunder methods are used in object creation.

`__new__`
- It creates a new object of the class.
- It is called before `__init__` method.
- Returns the newly created object.
- It is a class method.

<br>

`__init__`
- It initializes the newly created object.
- It is called after `__new__` method.
- Returns `None`.
- It is an instance method.

<br>

```python
class MyClass:
    def __new__(cls, *args, **kwargs):
        print("Calling __new__")
        instance = super().__new__(cls)
        return instance

    def __init__(self, value):
        print("Calling __init__")
        self.value = value

obj = MyClass(10)
```

<br>
<br>
<br>

### Python Lambda Function

- A lambda function in python is an anonymous, inline function.
- It is mostly used with buit-in functions like filter(), map(), etc.

<br>

```python
a = [1, 2, 3, 4, 5]

even = list(filter(lambda x: x % 2 == 0, a))
print(even)  # Output: [2, 4]
```

<br>
<br>
<br>

### Python Deep Copy vs Shallow copy

- `deepcopy()` recursively copies the outer iterable and the nested iterables as well.
- Shallow copy or simply `copy()` copies only the outer iterable.

<br>

```python
from copy import deepcopy, copy

a = [[1, 2], 3, 4]

a_scopy = copy(a)
a[0].append(5)
print(a_scopy)  # Output: [[1, 2, 5], 3, 4]

a = [[1, 2], 3, 4]

a_dcopy = deepcopy(a)
a[0].append(5)
print(a_dcopy)  # Output: [[1, 2], 3, 4]
```

<br>
<br>
<br>

### Python Generators

- Generators in python are lazily evaluated objects.
- The next loop is ran only when it is required.
- It uses `yield` keyword (to return a single value) instead of returning all values at once.
- Like normal iterables, generators **do not** evaluate all the results once and store them in memory.

<br>

```python
def count_to_five():
    i = 1
    while i <= 5:
        yield i
        i += 1

for i in count_to_five():
    print(i, end=" ")  # Output: 1 2 3 4 5 
```

<br>

#### How generators work?
- When `yield` is executed, the function pauses and returns a value.
- The next call to `next()` resumes execution from where it left off.

<br>

#### Generator expression

- A generator expression is a compact inline way to create a generator.
- It doesn't create the whole list in memory, it lazily produces items when needed.
- It is similar to list comprehension, but uses **parentheses** instead of sqaure brackets.

```py
(expression for item in iterable if condition)
```

<br>

#### Generator exhaustion

- A generator can be consumed only once.
- Once all values are yielded, it is **exhausted**.
- Iterating over the same generator again will yield nothing.

---

- If we want to reuse the values, store the values using **list comprehension**.
- We can also create a new generator by calling the function again `gen_2 = count_to_five()`.

<br>

#### StopIteration exception

- When a generator is exhausted, it raises a `StopIteration` exception internally.
- `for` loop uses this exception to stop the iteration.
- Any value returned from the generator will get attached to this exception.

<br>

#### Chaining generators

- Chaining generators refer to using one generator inside another generator.
- This is done using `yield from` keyword.
- Used to access the return value of first generator.

```py
def worker():
    yield 1
    yield 2
    return "done"

def manager():
    result = yield from worker()
    print(result)  # done
```

<br>

#### Key points

- Generators are memory efficient because they generate values on demand instead of storing them.
- Generators are not inherently thread-safe.

<br>
<br>
<br>

### Python Iterator

- An iterator is an object that allows us to iterate over a collection of elements one at a time.
- An iterator must implement these two methods:
    - `__iter__()`: Returns the iterator object itself.
    - `__next__()`: Returns the next element. Raises `StopIteration` indicating end of iteration.

```py
class MyIterator:
    def __init__(self, n: int) -> None:
        self._n = n
        self._i = 0
    
    def __iter__(self):
        # Return iterator object itself
        return self

    def __next__(self):
        if self._i < self._n:
            value = self._i
            self._i += 1
            return value

        raise StopIteration

def main():
    iterator = MyIterator(5)

    for i in iterator:
        print(i)

main()

# Output:
# 0
# 1
# 2
# 3
# 4
```

- Above example creates an iterator that iterates from `[0, n)`.
- **Note**: A `for` loop internally uses `iter()` and `next()` method.

<br>

Usecases:
- Reading large files line by line. Prevents loading the entire file into memory.
- Fetching database records in chunk instead of all at once.
- Lazily doing expensive computation when iterating over an iterable one by one.

<br>
<br>
<br>

### Generator vs Iterator

| Feature                | Iterator                                                | Generator                                                         |
| ---------------------- | ------------------------------------------------------- | ----------------------------------------------------------------- |
| **Definition**         | An object that implements `__iter__()` and `__next__()` | A special kind of iterator built using a function with `yield`    |
| **Creation**           | Requires writing a full class                           | Created using a simple function or generator expression           |
| **Memory Usage**       | Depends on implementation; may store large data         | Always lazy, produces one value at a time (very memory-efficient) |
| **Ease of Use**        | More code; verbose                                      | Very easy; minimal code                                           |
| **State Handling**     | You manage state manually                               | Python handles state automatically                                |
| **Infinite Sequences** | Possible but more manual work                           | Very natural and simple                                           |
| **Readability**        | Can be harder to maintain                               | Much more readable for sequential logic                           |

<br>

In short:
- Use a **Generator** for simple, linear and lazy iteration.
- Use an **Iterator** for complex, object-oriented iteration.

<br>
<br>
<br>

### Python Iterables

- An iterable is an object we can loop over.
- For an object to be an iterable, it must implement `__iter__()` or `__getitem__()` method.
- Most common iterables: List, Tuple, Dict, Set, String, etc.

<br>

- **Iterable**: Something you can iterate over.
- **Iterator**: The object that does the iteration one item at a time.
- When we use a `for` loop, Python does this internally:

```py
iterable = [1, 2, 3]

iterator = iter(iterable)   # gets an iterator
print(next(iterator))       # 1
print(next(iterator))       # 2
print(next(iterator))       # 3
```

<br>
<br>
<br>

### Python Exception Handling

- Exceptions are handled using try - except block.
- We can except specific exceptions like `ValueError`, `ZeroDivisionError`, etc.
- We can also except generic exceptions using `except Exception as e`.
- The code inside else is executed, only if no exception occured.
- The code inside finally is always executed, whether an exception occured or not.

<br>

```python
def exception_handling(num):
    try:
        print("try")
        10 / num
        print("end")
    except ZeroDivisionError:
        print("error")
    else:
        print("else")
    finally:
        print("finally")

exception_handling(1)
# try
# end
# else
# finally

exception_handling(0)
# try
# error
# finally
```

<br>

- The code inside `finally` block will always be executed. Even if we have `return` statement in `try` and `except` blocks.
- If there is a `return` statement in `finally`, this value will be returned else the value from `try` and `except` blocks will be returned.

<br>

```python
def exception_handling(num):
    try:
        print("try")
        10 / num
        print("end")
        return 1
    except ZeroDivisionError:
        print("error")
        return 2
    else:
        print("else")
    finally:
        print("finally")

res = exception_handling(1)
print(res)
# try
# end
# finally
# 1

res = exception_handling(0)
print(res)
# try
# error
# finally
# 2
```

<br>
<br>
<br>

### Python GIL

- GIL stands from Global Interpreter Lock.
- Each python process has a GIL lock associated with it.
- GIL allows only 1 thread execution at a time by the python interpreter.
- It prevents multithreading and the issues associated with it (like race conditions, data corruption, etc.).
- Due to limitation of having only single thread, true multithreading is not possible. Thus the performance of the applications can be reduced.
- But if we are dealing with I/O bound operations, we can leverage threading module & context switching.

<br>
<br>
<br>

### Python == vs is

- `is` check if the 2 object are same. i.e. they point to same memory location.
- `==` check if the 2 objects have same value. i.e. the content of memory location.

<br>

```python
a = [1, 2]
b = [1, 2]
print(a == b)  # True
print(a is b)  # False

a = None
b = None
print(a == b)  # True
print(a is b)  # True

a = 1
b = 1
print(a is b)  # True
# Integer Interning
```

<br>
<br>
<br>

### Python \_\_str\_\_ vs \_\_repr\_\_

- `__str__` is aimed for end user.
- It provides end user readable string.
- Use `str()` or `print()` to see this string.

<br>

- `__repr__` is aimed for developer.
- It provides developer readable string. This string can be used for object creation.
- use `repr()` or type the object name in REPL to see this string.

<br>

- Both of these provides string representation of objects.
- If `__str__` is not implement, `__repr__` is used. If that is also not implemented, default python memory string is used.

<br>

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def __str__(self):
        return f"{self.name}, {self.age} years old"

    def __repr__(self):
        return f"Person('{self.name}', {self.age})"

p = Person("Alice", 30)

print(str(p))  # Alice, 30 years old
print(repr(p))  # Person('Alice', 30)
```

<br>
<br>
<br>

### Python with Statement

- The `with` statement is used to simplify resource management.
- It makes sure resources are aquired & released properly.
- Example: opening & closing a file, connecting to database, etc.
- If any exception occurs within the with block, it is handled gracefully making sure the resource is closed properly.

<br>

```python
with open(“my_file.txt”, “r”) as file:
	content = file.read()
```

<br>
<br>
<br>

### Python Interning

- Interning is an optimization technique where identical immutable objects (like integers and strings) are reused to save memory and speed up comparisons.
- Python automatically interns small integers in the range -5 to 256.
- Larger integers may not be interned.

<br>

```python
a = 100
b = 100
a is b  # True

a = 1000
b = 1000
a is b  # False (usually)
```

<br>

- String literals that are short or valid identifiers ('hello', 'foo123') are interned.
- String literals with space in them may not be interned.

<br>

```python
a = "hello"
b = "hello"
a is b  # True

a = "hello world"
b = "hello world"
a is b  # Might be False

```

<br>

- Interning only applies to immutable types (e.g., int, str, bool, None, float (some cases)).
- It’s an optimization, not a guarantee.

<br>
<br>
<br>

### Python Metaclass

- A metaclass is a class of a class.
- It controls class creation behavior.
- Automatically modify class attributes or methods.
- Commonly used when we are creating a python package.
- A class can only have one metaclass.

<br>

Example:

```py
class MethodNameChecker(type):
    def __new__(cls, name, bases, attrs):
        for attr in attrs:
            if callable(attrs[attr]) and not attr.startswith('do_'):
                raise TypeError(f"Method {attr} must start with 'do_'")
        
        return super().__new__(cls, name, bases, attrs)

class MyActions(metaclass=MethodNameChecker):
    def do_run(self): pass
    def run_fast(self): pass  # Raises TypeError
```

<br>
<br>
<br>
<br>
<br>

### Python Built-in Functions

#### `zip()`

- Combines multiple iterables element-wise into tuples.
- `zip(a, b)` pairs `a[i]` with `b[i]` until the shortest iterable ends. 

```py
iter1 = "abc"
iter2 = [1, 2, 3]
iter3 = [n * n for n in iter2]

for a, b, c in zip(iter1, iter2, iter3):
    print(a, b, c)

zip_obj = zip(iter1, iter2, iter3)
print(zip_obj)  # <zip object at 0x797a33091ac0>
print(list(zip_obj))  # [('a', 1, 1), ('b', 2, 4), ('c', 3, 9)]
print(list(zip_obj))  # []

# Output:
# a 1 1
# b 2 4
# c 3 9
```

<br>
<br>

#### `filter()`

- Keeps only the elements of an iterable that satisfy a condition.
- Signature: `filter(function, iterable)`.
    - `function` is usually a lambda function that returns `True` or `False`.
    - Returns a lazy, single-use iterator.
- Note: In modern Python, **List Comprehension** is usually preferred.

```py
iter1 = [1,2,3,4,5]

filtered = filter(lambda x: x % 2 == 0, iter1)

print(type(filtered))  # <class 'filter'>
print(list(filtered))  # [2, 4]
print(list(filtered))  # []
```

<br>
<br>

#### `map()`

- Apply a function to each element of an iterable.
- Signature: `map(function, iterable)`.
    - `function` is usually a lambda function that returns something.
    - Returns a lazy, single-use iterator.
- Note: In modern Python, again **List Comprehension** is usually preferred.

```py
iter1 = [1,2,3,4,5]

doubled = map(lambda x: x*2, iter1)

print(type(doubled))  # <class 'map'>
print(list(doubled))  # [2, 4, 6, 8, 10]
print(list(doubled))  # []
```

<br>
<br>

#### `enumerate()`

- Gives access to element and its index.

```py
iter1 = "abc"

for i, l in enumerate(iter1):
    print(i, l)

enum_obj = enumerate(iter1)
print(enum_obj)  # <enumerate object at 0x7214f0e88700>
print(list(enum_obj))  # [(0, 'a'), (1, 'b'), (2, 'c')]
print(list(enum_obj))  # []

# Output:
# 0 a
# 1 b
# 2 c
```

<br>
<br>

#### `range()`

- Generates a sequences of numbers.
- This sequence can be in ascending or descending order with or without skipping elements.

```py
r1 = range(5)

print(type(r1))  # <class 'range'>
print(list(r1))  # [0, 1, 2, 3, 4]
print(list(r1))  # [0, 1, 2, 3, 4]

r2 = range(5, 10)
print(list(r2))  # [5, 6, 7, 8, 9]

r3 = range(1, 10, 2)
print(list(r3))  # [1, 3, 5, 7, 9]

r4 = range(5, 3)
print(list(r4))  # []

r5 = range(5, 3, -1)
print(list(r5))  # [5, 4]
```

<br>
<br>

#### `sorted()`

- Returns a sorted list from any iterable.
- We can also pass a sorting key.
- Sorting can be done in ascending or descending order.

```py
a = [2,3,1,4,3]

sorted_a = sorted(a)
print(type(sorted_a))  # <class 'list'>
print(list(sorted_a))  # [1, 2, 3, 3, 4]
print(list(sorted_a))  # [1, 2, 3, 3, 4]

sorted_a_rev = sorted(a, reverse=True)
print(list(sorted_a_rev))  # [4, 3, 3, 2, 1]


b = ["banana", "apple", "peach"]

sorted_b_len = sorted(b, key=len)
print(list(sorted_b_len))  # ['apple', 'peach', 'banana']

sorted_b_1 = sorted(b, key=lambda x: x[1])
print(list(sorted_b_1))  # ['banana', 'peach', 'apple']
```

<br>
<br>

#### `reversed()`

- Reverse the given iterable.
- Returns an iterator.

```py
a = [2,3,1,4,3]

reversed_a = reversed(a)
print(type(reversed_a))  # <class 'list_reverseiterator'>
print(list(reversed_a))  # [3, 4, 1, 3, 2]
print(list(reversed_a))  # []
```

<br>
<br>

#### `any()` and `all()`

`any()`

- Return True if `bool(x)` is True for any x in the iterable.
- If the iterable is empty, return False.

<br>

`all()`

- Return True if `bool(x)` is True for all values x in the iterable.
- If the iterable is empty, return True.

```py
a = [False, True, False]

print(any(a))  # True
print(all(a))  # False
```

<br>
<br>

#### `sum()`

- Returns the sum of a 'start' value (default: 0) plus an iterable of numbers.
- It will reject non-numeric types.

```py
a = [1,2,3]
print(sum(a))  # 6
```

<br>
<br>
<br>
<br>
<br>

### Python Dunder Methods

#### Object Initialization & Representation

`__new__`

- Responsible for creating the object.
- `__new__` is called before `__init__`.
- `__new__` is often used to implement singletons.

<br>
<br>

`__init__`

- Constructor: Called after `__new__` to initialize the already-created instance.
- Should not return anything except `None`; returning a value causes an error.
- Best place to set instance attributes and perform object setup.

<br>
<br>

`__del__`

- Acts as a destructor, called when an object is about to be garbage-collected.
- **Not reliable** for critical cleanup, because garbage collection timing is unpredictable.
- Should be used sparingly; prefer context managers (`with`) for deterministic cleanup.

<br>
<br>

`__str__`

- Defines the object’s human-readable string representation (used by `print()` and `str()`).
- Should return something clear and user-friendly, not necessarily detailed.
- Falls back to `__repr__` if not defined.

<br>
<br>

`__repr__`

- Defines the object’s official, unambiguous string representation (used in REPL, debugging).
- Should ideally return a string that can recreate the object, e.g., `ClassName(arg=value)`.
- If only `__repr__` is defined, it also serves as `__str__`.

<br>
<br>

```py
class MyClass:
    def __new__(cls, *args, **kwargs):
        print("creating new object")
        return super().__new__(cls)

    def __init__(self, num: int):
        print("initializing object")
        self.num = num

    def __del__(self):
        print("deleting object")

    def __str__(self) -> str:
        return f"Num is {self.num}"

    def __repr__(self) -> str:
        return f"MyClass(num={self.num})"


a = MyClass(10)
# creating new object
# initializing object

print(str(a))  # Num is 10
print(repr(a))  # MyClass(num=10)

# deleting object
```

<br>
<br>
<br>

#### Arithmetic Operators

| Method                      | Purpose                           |
| --------------------------- | --------------------------------- |
| `__add__(self, other)`      | `+`                               |
| `__sub__(self, other)`      | `-`                               |
| `__mul__(self, other)`      | `*`                               |
| `__truediv__(self, other)`  | `/`                               |
| `__floordiv__(self, other)` | `//`                              |
| `__mod__(self, other)`      | `%`                               |
| `__pow__(self, other)`      | `**`                              |

<br>

```py
class MyClass:
    def __init__(self, num: int):
        self.num = num

    def __add__(self, other):
        return self.num + other.num

    def __sub__(self, other):
        return self.num - other.num

    def __mul__(self, other):
        return self.num * other.num

    def __truediv__(self, other):
        return self.num / other.num

    def __floordiv__(self, other):
        return self.num // other.num

    def __mod__(self, other):
        return self.num % other.num

    def __pow__(self, other):
        return self.num ** other.num

a = MyClass(3)
b = MyClass(4)

print(a + b)  # 7
print(a - b)  # -1
print(a * b)  # 12
print(a / b)  # 0.75
print(a // b)  # 0
print(a % b)  # 3
print(a ** b)  # 81
```

<br>
<br>
<br>

#### Comparison Operators

| Method                | Purpose          |
| --------------------- | ---------------- |
| `__eq__(self, other)` | Equality (`==`)  |
| `__ne__(self, other)` | Not equal        |
| `__lt__(self, other)` | Less than        |
| `__le__(self, other)` | Less or equal    |
| `__gt__(self, other)` | Greater than     |
| `__ge__(self, other)` | Greater or equal |

<br>

```py
class MyClass:
    def __init__(self, num: int):
        self.num = num

    def __eq__(self, other) -> bool:
        return self.num == other.num 

    def __ne__(self, other) -> bool:
        return self.num != other.num 

    def __lt__(self, other) -> bool:
        return self.num < other.num 

    def __le__(self, other) -> bool:
        return self.num <= other.num 

    def __gt__(self, other) -> bool:
        return self.num > other.num 

    def __ge__(self, other) -> bool:
        return self.num >= other.num 

a = MyClass(3)
b = MyClass(4)

print(a == b)  # False
print(a != b)  # True
print(a >= b)  # False
print(a >= b)  # False
print(a < b)  # True
print(a <= b)  # True
```

<br>
<br>
<br>

#### Sequence & Map Behaviour

`__len__`

<br>
<br>

`__getitem__`

<br>
<br>

`__setitem__`

<br>
<br>

`__delitem__`

<br>
<br>

`__contains__`

<br>
<br>

`__iter__` and `__next__`

<br>
<br>
<br>

#### Attribute Access

`__getattr__`

<br>
<br>

`__getattribute__`

<br>
<br>

`__setattr__`

<br>
<br>

`__delattr__`

<br>
<br>

`__dir__`

<br>
<br>
<br>

#### Hashing & Truthiness

`__hash__`

<br>
<br>

`__bool__`

<br>
<br>
<br>
<br>
<br>

### Python Duck Typing

- Duck typing is a programming concept.
- Here, the type or class of an object is less important than the methods or behaviors it supports.
- If the object behaves correctly (e.g., has a certain method), it is accepted, regardless of its actual class.

```
If it looks like a duck and quacks like a duck, it’s a duck.
```

<br>

```py
class Duck:
    def quack(self):
        print("Quack!")

class Person:
    def quack(self):
        print("I'm pretending to be a duck!")

def make_it_quack(thing):
    thing.quack()   # Works as long as 'thing' has a quack() method

make_it_quack(Duck())    # Quack!
make_it_quack(Person())  # I'm pretending to be a duck!
```

- Both objects work because they have a `quack()` method.
- Advantages: 
    - Flexibility: Code works with any object providing the right interface.
    - Fewer type checks: Encourages cleaner, more Pythonic code.
    - Supports polymorphism naturally.

<br>
<br>
<br>
<br>
<br>

### String Interpolation

#### `.format()` method

- This method uses curly braces `{}` as placeholders within a string.
- It works by calling the `.format()` method on a string literal and pass the values as arguments.

Advantages:

- Supports positional and keyword arguments.
- Can be used on predefined string variables

Disadvantages:

- Verbose and hard to read with many variables.
- Slower than f-strings.

```py
template = "My name is {name}"
output = template.format(name="John")

print(type(template))  # <class 'str'>
print(template)  # My name is {name}

print(type(output))  # <class 'str'>
print(output)  # My name is John
```

<br>
<br>

#### f-string

- Also called formatted string.
- It is defined as a string literal and prefixed with a `f`.
- It allows to embed Python expressions directly inside the string.

Advantages:

- Evaluated at the runtime and has the fastest performance.
- Most readable and concise syntax.

Disadvantages:

- Expressions are evaluated immediately. Cannot be stored for later use.
- Potential security risk if used with direct user input.

```py
name = "John"
output = f"My name is {name}"

print(type(output))  # <class 'str'>
print(output)  # My name is John
```

<br>
<br>

#### `string.Template`

- It uses `$` as a placeholder within a string.
- It creates a template, which can be later used to substitute place holders with actual values.

Advantages:

- Safest for user-provided strings.
- It prevents arbitrary code execution.

Disadvantages:

- Does not support formatting or logic inside the string.
- Slower than f-strings.

```py
from string import Template

template = Template("My name is $name")
output = template.substitute(name="John")

print(type(template))  # <class 'string.Template'>
print(template)  # <string.Template object at 0x77a062f8c830>

print(type(output))  # <class 'str'>
print(output)  # My name is John
```

<br>
<br>

#### t-string

- Also called template string.
- It is defined as a string literal and prefixed with a `t`.
- A t-string returns a template object and not a string literal.
- It stores the `string` parts and `value` parts of the template separately. This helps in preventing SQL injection and similar risks.

Advantages:

- Has security of `string.Template` and simplicity of f-string.
- The string isn't fully built until it is ready for evaluation.

Disadvantages:

- More complex to understand and use.
- Slower than f-strings for simple, one-off string concatenation.

```py
name = "John"
template = t"My name is {name}"

print(type(template))  # <class 'string.templatelib.Template'>
print(template)  # <string.templatelib.Template object at 0x77a062f8c830>

print(template.strings)  # ('My name is ', '')
print(template.interpolations)  # (Interpolation('John', 'name', None, ''),)
print(template.values)  # ('John',)
```

<br>
<br>
<br>
<br>
<br>

### `match-case` Statement

- `match-case` statement is Python's version of a `switch` statement.
- It is more powerful than `switch` statement.
- It supports **structural pattern matching**, meaning it can match not only values but also data structures and patterns.
- Introduced in Python 3.10 

<br>
<br>

#### Types of matching

Literal Matching

- Matches exact values (like switch-case).

```py
match status:
    case 200:
        print("Success")
    case 404:
        print("Not Found")
    case _:  # default case
        print("Unknown")
```

<br>

OR Pattern (`|`)

- Matches multiple values in one case.

```py
match day:
    case "Sat" | "Sun":
        print("Weekend")
    case _:
        print("Weekday")
```

<br>

Variable Capture

- Captures the matched value into a variable.

```py
match value:
    case 5:
        print("five")
    case x:
        print(f"Value is {x}")
```

<br>

Guarded Matching (`if` condition)

- Adds an extra condition using `if`.

```py
match number:
    case x if x > 0:
        print("Positive")
    case x if x < 0:
        print("Negative")
    case _:
        print("Zero")
```

<br>

Sequence Matching (List / Tuple)

- Matches structure of sequences.

```py
match data:
    case [a, b]:
        print("Two elements")
    case [a, b, c]:
        print("Three elements")
```

<br>

Wildcard Pattern (`_`)

- Matches anything, but doesn’t capture it.
- Equivalent to `else`.

```py
case _:
    print("Default case")
```

<br>

Nested Pattern Matching

- Matches complex nested structures.

```py
match point:
    case (0, 0):
        print("Origin")
    case (x, 0):
        print("X-axis")
    case (0, y):
        print("Y-axis")
```

<br>

Dictionary Mapping Pattern

- Matches dictionary keys and values.
- Subset matching.

```py
match response:
    case {"status": 200, "data": data}:
        print("Success")
    case {"status": 404}:
        print("Not Found")
```

<br>

Class Pattern (Object Matching)

- Matches objects based on attributes.

```py
class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

match p:
    case Point(x=0, y=0):
        print("Origin")
```

<br>
<br>

#### Key points

- Patterns are checked from top to bottom.
- No fall-through behavior like C++, only first matching block runs.
- `_` should always be the last case.

<br>
<br>
<br>
<br>
<br>

### Python `__slots__`

- `__slots__` is a class-level attribute in Python.
- It allows us to explicitly declare which attributes a class instance is allowed to have.
- Normally, Python stores instance attributes in `__dict__` attribute.
- When we use `__slots__`, Python removes this `__dict__` attribute and stores instance attributes in a memory-efficient structure.

<br>

```py
class User:
    __slots__ = ("name", "age")

    def __init__(self, name, age):
        self.name = name
        self.age = age

u = User("Alice", 25)
print(hasattr(u, "__dict__"))  # False

# AttributeError: 'User' object has no attribute 'email'
u.email = "a@example.com"
```

<br>
<br>

#### Advantages

- Objects using `__slots__` consume less memory. Noticeable with thousands of instances.
- Attribute access is slightly faster as no dictionary lookup is needed.
- Cleaner and explicit class definition.

<br>
<br>

#### Limitations

- Removes `__dict__` attribute. This can cause issue at places where `__dict__` is used.
- Cannot add new attributes other than mentioned in `__slots__`.
- Slots are not automatically inherited. They need to be defined again in the child class.
- Not worth it for small objects.

<br>
<br>

#### Key Features

- Attributes are stored in a fixed, more memory-efficient structure.
- Slots does not affect property decorator.

```py
# Slots impact on inheritance

class Base:
    __slots__ = ("a",)

class Child(Base):
    __slots__ = ("b",)

c = Child()
c.a = 1
c.b = 2
```

<br>
<br>
<br>
<br>
<br>

### Python Descriptors

- Descriptors are Python objects that control attribute access (get, set, delete) in other objects.
- A Python object is called a descriptor if it defines one or more of these methods:
    - `__get__(self, instance, owner)`
    - `__set__(self, instance, value)`
    - `__delete__(self, instance)` (not `__del__`)
- Descriptors are the core mechanism behind properties, methods, static and class methods.

<br>
<br>

#### Working

- Descriptors work via attribute lookup rules.
- When we do `obj.attribute`, Python first check if that attribute has `__get__`, `__set__` or `__delete__` methods implemented.
- If any of these methods is present, descriptor protocol is invoked.
- If none of these methods exist, Python does normal attribute lookup using `__dict__`.

---

Preference order

1. Data descriptors: Attributes that have `__get__` + (`__set__` or `__delete__`) methods implemented.
2. Instance attributes: Normal instance attributes.
3. Non-data descriptors: Attributes that have only `__get__` method implemented.
4. Class attributes: Normal class attributes.

<br>
<br>

#### Data descriptors

- Data descriptors are descriptors that implement `__get__` and (`__set__` or `__delete`) methods.
- They have highest preference in attribute lookup.

```py
class Positive:
    def __set_name__(self, owner, name):
        self.private_name = "_" + name

    def __get__(self, instance, owner):
        return getattr(instance, self.private_name, 0)

    def __set__(self, instance, value):
        if value <= 0:
            raise ValueError("Value must be positive")
        setattr(instance, self.private_name, value)

    def __delete__(self, instance):
        raise AttributeError("Cannot delete this attribute")

class Product:
    price = Positive()


p = Product()
p.price = 50          # __set__
print(p.price)        # __get__

del p.price           # __delete__ -> error
```

<br>
<br>

#### Non-data descriptors

- Non-data descriptors are descriptors that implement only the `__get__` method.
- They have lowest preference (3rd) in attribute lookup.

```py
class FullName:
    def __get__(self, instance, owner):
        return f"{instance.first} {instance.last}"

class User:
    full_name = FullName()   # non-data descriptor

    def __init__(self, first, last):
        self.first = first
        self.last = last


u = User("John", "Doe")
print(u.full_name)          # John Doe

u.full_name = "Override"
print(u.full_name)          # Override (instance attribute wins)
```

<br>
<br>

#### Instance and owner

- `instance` is the object accessing the attribute.
- `owner` is the class where the descriptor is defined.

<br>

```py
class MyDescriptor:
    def __get__(self, instance, owner):
        print("instance:", instance)
        print("owner:", owner)
        return "value"

class MyClass:
    attr = MyDescriptor()


MyClass.attr
# instance: None
# owner: <class '__main__.MyClass'>

MyClass().attr
# instance: <__main__.MyClass object at 0x77b247e5ba30>
# owner: <class '__main__.MyClass'>
```

<br>
<br>

#### Advantages

- Code Reusability: Write logic once, reuse across classes.
- Code Abstraction: Internal implementations are abstracted from the user, providing a clean API.
- Centralized Control: Access rules and validation logic lives in one place. 

<br>
<br>

#### Limitations

- Complexity: Attribute lookups are hard to follow and debug.
- Overkill for simple cases: `@property` is usually enough for most cases.
- Limited awareness: Many developers don't know descriptors well.

<br>
<br>

#### Usecases

- Reusable attribute get and set logic.
- Attributes value validation before setting.
- Computing attributes before getting.
- Adding caching to attributes.

<br>
<br>

#### Key points

- Descriptors are heavily used in Django ORM fields, SQLAlchemy, etc.
- `@property` decorator is a built-in descriptor.
- Python internally uses descriptors for instance and class methods.

<br>
<br>
<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
<br>

### More

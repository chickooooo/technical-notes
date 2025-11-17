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

```python
"""
Simple decorator
"""

def authenticated(func):
    # add auth logic
    print("auth")
    # return view function
    return func


@authenticated
def my_view():
    # add view logic here
    print("view")

my_view()
# auth
# view


"""
Function arguments
"""

def authenticated(func):
    # add auth logic
    print("auth")
    # return view function
    return func

@authenticated
def my_view(num):
    # add view logic here
    print("view", num)

my_view(1)
# auth
# view 1


"""
Decorator arguments
"""

def authenticated(**kwargs):
    # access parameters
    print(kwargs)
    # add auth logic
    print("auth")

    def decorator(func):
        # return view function
        return func

    # return decorator
    return decorator

@authenticated(admin_only=True)
def my_view(num):
    # add view logic here
    print("view", num)

my_view(1)
# Output
# {'admin_only': True}
# auth
# view 1
```

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

### More

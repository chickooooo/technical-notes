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
# Python Threading

<br>
<br>
<br>
<br>
<br>

## Index

- [Threading in Python](#threading-in-python)
- [GIL impact on threading](#gil-impact-on-threading)
- [Typical threading workflow](#typical-threading-workflow)
- [ThreadPoolExecutor](#thread-pool-executor)
- [Future](#future)
- [Synchronization Primitives](#synchronization-primitives)

<br>
<br>
<br>
<br>
<br>

### Threading in Python

#### What is threading?

- Threading allows a program to run multiple threads (smaller units of a process) concurrently.
- The execution is concurrent, meaning tasks appear to run at the same time, even on a single CPU (through context switching).
- Each thread runs in the same memory space, sharing data and resources.
- This makes inter-thread communication easy but introduces challenges like race conditions.

<br>

#### When to use threading?

- Use threading to speed up I/O-bound tasks like:
    - Network requests
    - Disk read/write
    - Reading/writing to a file
    - Waiting for user input
- While one thread is waiting for I/O it releases the GIL, allowing other threads to execute.
- This gives appearance of multiple threads running simultaneously.

<br>

#### When NOT to use threading:

- Avoid threading for CPU-bound tasks like:
    - Image processing
    - Heavy computations, etc.
- Python's GIL allows only one thread to execute Python bytecode at a time, so threading won’t make CPU-heavy tasks faster.
- Instead use `multiprocessing` for parallel CPU work.

<br>
<br>
<br>

### GIL impact on threading

#### What is the GIL?

- GIL stands for Global Interpreter Lock.
- It is a mutex (a kind of lock) in the **CPython interpreter** that allows **only one thread** to execute Python bytecode at a time, even on multi-core processors.

<br>

#### Why does it exist?

- It simplifies memory management by preventing multiple threads from modifying Python objects simultaneously.
- It avoides race conditions in CPython’s internal memory structures.

<br>

#### How GIL impacts threading performance?

- For I/O-bound tasks, when a thread is waiting for a response, it releases the GIL, allowing other threads to perform their task. Hence for I/O-bound tasks, GIL does not degrade the performance.
- For CPU-bound tasks, as GIL allows only one thread to execute at a time (on a single core), parallel processing is not possible. Hence the performance is not efficient and can even degrade due to context switching.
- If we need true parallelism, use `multiprocessing`.

<br>
<br>
<br>

### Typical threading workflow

1. Define a task:
    - This task is referred to as **target** function.
    - It includes the code each thread will perform.
    - Ideally, it should be an I/O-bound task to benefit from threading.

2. Create the threads:
    - Create the required number of threads.
    - We can use `threading.Thread` class or `concurrent.futures.ThreadPoolExecutor`.
    - Each thread will have a:
        - `target`: function to execute.
        - `args` / `kwargs`: arguments to the function.
        - `name`: optional, helpful for debugging.

3. Start the threads:
    - Use `.start()` method on each thread to schedule it's execution.

4. Synchronize threads:
    - Use synchronization primitives like `Lock`, `RLock`, `Queue`, `Event`, `Semaphore`, `Condition`, `Barrier`, if needed.
    - These are used to safely access shared resources and manage communication between threads.

5. Wait for threads to complete:
    - Use `.join()` method on each thread so the main thread waits until all worker threads finish.

6. Collect the result:
    - If we want to collect the result from each thread, we can use shared data structures.
    - Make sure to use locks when writing to these shared resources.
    - If we are using `ThreadPoolExecutor`, it internally collects the results and makes them available to us.

<br>

```py
import time
import random
import threading
from typing import List


def download_file(url: str, contents: List[str], lock: threading.Lock):
    # Download file
    print(f"Download started for {url}")
    delay = random.randint(1, 5)
    time.sleep(delay)
    print(f"Download completed for '{url}' in {delay} seconds")

    # Get file content
    file_content = url.split("/")[-1]

    # Add to contents safely
    with lock:
        contents.append(file_content)


def manual_threading(urls: List[str]) -> List[str]:
    # Will hold threads
    threads: List[threading.Thread] = []
    # Will hold file contents
    contents: List[str] = []
    # Create a shared lock
    lock = threading.Lock()

    # Create and schedule threads
    for i, url in enumerate(urls):
        thread = threading.Thread(
            target=download_file,
            kwargs={
                "url": url,
                "contents": contents,
                "lock": lock,
            },
            name=f"thread_{i+1}",
        )
        thread.start()
        threads.append(thread)

    # Wait for all threads to complete
    print("Waiting for threads to complete...")
    for thread in threads:
        thread.join()

    return contents


if __name__ == "__main__":
    urls = [
        "https://s3.amazon.com/file_1.txt",
        "https://s3.amazon.com/file_2.txt",
        "https://s3.amazon.com/file_3.txt",
    ]

    contents = manual_threading(urls)
    print(contents)

# Output:
# Download started for https://s3.amazon.com/file_1.txt
# Download started for https://s3.amazon.com/file_2.txt
# Download started for https://s3.amazon.com/file_3.txt
# Waiting for threads to complete...
# Download completed for 'https://s3.amazon.com/file_2.txt' in 1 seconds
# Download completed for 'https://s3.amazon.com/file_3.txt' in 3 seconds
# Download completed for 'https://s3.amazon.com/file_1.txt' in 5 seconds
# ['file_2.txt', 'file_3.txt', 'file_1.txt']
```

<br>
<br>
<br>

### ThreadPoolExecutor

- `ThreadPoolExecutor` is a high level interface for managing a pool of threads.
- It simplifies multithreading. We just have to provide a target function and arguments that will be passed to this function.
- `ThreadPoolExecutor` efficiently manages the threads, their execution and returns back the result.
- It is part of `concurrent.futures` module.
- Useful when we require a pool of threads to perform tasks concurrently.

<br>
<br>

#### `.map()` method

- `.map()` method is used when we want to execute the same function multiple times using a different set of arguments.
- It takes a target function as an input and `n` iterables. Here `n` is equal to the number of arguments that function accepts.
- Each row of the iterables is passed as arguments to the function and then it is scheduled for execution.
- `.map()` method blocks until all the threads have finished their execution.
- After all executions are completed, `.map()` returns the results in the **same order** as that of input arguments.

```py
from concurrent.futures import ThreadPoolExecutor

def divide(a: int, b: int) -> float:
    return a / b

def main():
    a_list = [2, 2, 10]
    b_list = [1, 2, 2]

    with ThreadPoolExecutor(max_workers=2) as executor:
        result = executor.map(divide, a_list, b_list)

    for item in result:
        print(item)

if __name__ == "__main__":
    main()

# Output:
# 2.0
# 1.0
# 5.0
```

<br>

Exception in thread

```py
def main():
    a_list = [2, 2, 10]
    b_list = [1, 0, 2]

    with ThreadPoolExecutor(max_workers=2) as executor:
        result = executor.map(divide, a_list, b_list)
        print("here")

    for item in result:
        print(item)

# Output:
# here
# 2.0
# Traceback (most recent call last):
# ...
# ZeroDivisionError: division by zero
```

- In the above example, when we call `.map()`, it schedules 3 tasks.
- All these 3 tasks complete their execution and `"here"` is printed to the console.
- Now, when we iterate the `result`, the first result gets printed `2.0`.
- The second result raises an exception and the iterator and program stops there.
- Even though third result is computed and is present in the `result`.

<br>
<br>

#### `.submit()` method

- `.submit()` schedules a function to be executed in a thread pool and returns a `Future` object.
- A `Future` object represents the result of an asynchronous operation.

```py
import time
from typing import List
from concurrent.futures import ThreadPoolExecutor, Future

def divide(a: int, b: int) -> float:
    time.sleep(1)
    return a / b

def main():
    a_list = [2, 2, 10, 3]
    b_list = [1, 0, 2, 2]

    with ThreadPoolExecutor(max_workers=2) as executor:
        # Will hold future objects
        futures: List[Future] = []

        for a, b in zip(a_list, b_list):
            future = executor.submit(divide, a, b)
            futures.append(future)
    
    for future in futures:
        exc = future.exception()
        if exc is not None:
            print(exc)  # Print exception
        else:
            print(future.result())  # Print result

if __name__ == "__main__":
    main()

# Output:
# 2.0
# division by zero
# 5.0
# 1.5
```

- Above code creates 2 threads and submits 4 tasks to be executed on these threads.
- Each `.submit()` call returns a `Future` object, which is stored in `futures` array.
- Once all the threads have finished execution, we iterate on these futures and print the result or exception.

<br>
<br>

#### `.shutdown()` method

- `.shutdown()` method does the clean up and free resources once all tasks have been completed.
- By default, it blocks until all tasks are completed.
- When the `with` block of `ThreadPoolExecutor` exits, the executor automatically calls `.shutdown()` method and waits for all tasks to finish before cleaning up.

```py
executor.shutdown(wait=True, cancel_futures=False)
# wait until all tasks are completed
# cancels any futures that haven’t yet started executing
```

<br>
<br>
<br>

### Future

- A `Future` object represents the result of an asynchronous operation.
- It is used to check the status of a task and retrieve result after completion.

<br>

#### `.done()`

- Returns `True` if the task has finished execution else `False`.
- A task execution finishes when it completes successfully or raises exception.

```py
future.done()  # True
```

<br>

#### `.result()`

- Blocks until the result of the task is available and then returns it.
- If the task raised an exception, it will be re-raised when calling `.result()`.
- We can also set a timeout indicating, how long to wait until we get the result.
- If the timeout expires, it raises `concurrent.futures.TimeoutError`.

```py
future.result(timeout=None)
```

<br>

#### `.exception()`

- Returns the exception raised by the task, or `None` if the task completed successfully.

```py
future.exception()
```

<br>

#### `.add_done_callback()`

- Registers a callback function that is called when the task finishes, regardless of whether it completes successfully or raises an exception.
- This callback function takes one argument, the `Future` object.

```py
future.add_done_callback(my_callback)

def my_callback(future: Future):
    print(f"Task result: {future.result()}")
```

<br>
<br>
<br>

### Synchronization Primitives

#### Lock

- A Lock (also known as a mutex) is the simplest synchronization primitive.
- It ensures **mutual exclusion**. That means, only one thread can hold the lock at a time.
- It is used to ensure only one thread can access the shared resource at a time.

Key behaviour:
- If a thread tries to acquire an already-held lock, it blocks until the lock is released. 
- It prevents race condition on shared resources.

```py
import threading
from concurrent.futures import ThreadPoolExecutor

def increment(counter_map, lock):
    with lock:
        counter_map["key"] += 1

def main():
    counter_map = {"key": 0}
    lock = threading.Lock()

    with ThreadPoolExecutor(max_workers=100) as executor:
        for _ in range(100):
            executor.submit(increment, counter_map, lock)

    print("Final counter:", counter_map["key"])

main()

# Output:
# Final counter: 100
```

<br>
<br>

#### RLock

- RLock stands for **Reentrant Lock**.
- A same thread can acquire RLock multiple times without deadlocking.

Use when:
- A function holding a lock calls another function that needs the same lock.
- We need **recursive acquisition** by the same thread.

Key behaviour:
- Maintains an internal counter indicating how many times the lock was acquired.
- A thread must release the lock the same number of times to unlock it.

```py
import threading

class BankAccount:
    def __init__(self, balance: int) -> None:
        self._balance = balance
        # Normal Lock() will cause deadlock
        self._lock = threading.RLock()

    def deposit(self, amount: int):
        with self._lock:
            self._balance += amount

    def withdraw(self, amount: int):
        with self._lock:
            self._balance -= amount
    
    def transfer(self, amount: int, other_account: BankAccount):
        # We need lock to make this operation atomic
        # No other deposit or withdrawl should interfere with a transfer
        with self._lock:
            # Same lock will be accessed during withdrawl
            self.withdraw(amount)
            other_account.deposit(amount)
```

- In the above example, we are acquiring the same lock twice in `transfer()` method.
- Using a standard `Lock()` would cause a deadlock, hence we have used `RLock()`.

<br>
<br>
<br>

### 

<br>
<br>
<br>

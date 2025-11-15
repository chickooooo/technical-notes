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
- [Thread Pool Executor](#thread-pool-executor) 

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

### Thread Pool Executor

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

<br>
<br>
<br>

### 

<br>
<br>
<br>

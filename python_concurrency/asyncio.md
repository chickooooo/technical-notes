# Async IO Package

<br>
<br>
<br>
<br>
<br>

## Index

- [Event Loop](#event-loop)
- [Coroutine](#coroutine)
- [Task](#task)

<br>
<br>
<br>
<br>
<br>

### Event Loop

#### What is an event loop?

- Event loop is the core of Python's `asyncio` framework.
- It is an infinite loop that schedules and executes **asynchronous tasks**, **callbacks**, and **I/O events**.
- It works on the principle of **Cooperative Multitasking**: tasks yield control using `await`, allowing other tasks to run.
- It works as a task orchestrator: It picks a coroutine, runs it until it hits an awaitable, pauses it, and runs another.

<br>

#### How event loop works?

- The event loop follows a predictable cycle:
    - Register tasks (coroutines, callbacks, I/O watchers).
    - Pick the next ready task from the queue.
    - Run the task until it hits `await` on an I/O operation or other awaitable.
    - Suspend the task and schedule it to resume once the awaited operation is done.
    - Repeat this cycle until all tasks are completed.

---

Key characteristics:
- The loop switches the task only when coroutines explicitly `await` (cooperative multitasking, not preemptive).
- I/O operations (like network or disk) are handled asynchronously via OS-level event notifications.
- The event loop ensures efficient concurrency using the single main thread.

---

Example:

```py
import asyncio

async def task(name: str):
    print(f"Task {name} started")
    await asyncio.sleep(1)  # yields control to event loop
    print(f"Task {name} finished")

async def main():
    await asyncio.gather(task("A"), task("B"))

if __name__ == "__main__":
    asyncio.run(main())
```

<br>

#### `asyncio.get_event_loop()` vs `asyncio.run()`

`asyncio.run(main())`:
- Modern approach.
- Creates a new event loop, runs the given coroutine to completion, and closes the loop automatically.
- It ensures a clean, well-scoped event loop lifecycle.
- Preferred for **top-level** entry points.

---

`asyncio.get_event_loop()`:
- Older, legacy interface.
- Retrieves the current event loop or creates one if missing.
- Not recommended for general usage in modern code due to potential confusion and lifecycle issues.

<br>

#### Blocking vs non-blocking operations

Blocking Operations:
- Stop the entire event loop until they finish.
- Bad for asyncio because no other tasks can run during the block.
- Examples: `time.sleep()`, heavy CPU processing, synchronous I/O.

---

Non-blocking Operations:
- Yield control back to the event loop using `await`.
- Allow other tasks to run concurrently.
- Examples: `asyncio.sleep()`, asynchronous I/O.

<br>
<br>
<br>

### Coroutine

#### What is a coroutine?

- A coroutine is a special Python function defined with `async def`.
- A coroutine can be paused and resumed.
- Coroutines enable asynchronous, non-blocking execution by yielding control using `await`.
- They support **cooperative multitasking**, meaning they voluntarily yield control.

<br>

#### Creating a coroutine

- Use `async` keyword in-front of the function definition to create a coroutine.
- Unlike normal functions, calling a coroutine does not execute it immediately.
- It returns a **coroutine object** which must be awaited or scheduled on the event loop.
- A **coroutine object** has a type of `Coroutine[SendType, ThrowType, ReturnType]`:
    - `SendType`: The type of the value you can send into the coroutine.
    - `ThrowType`: The type of exceptions that may be thrown into the coroutine.
    - `ReturnType`: The type of the value the coroutine returns when finished. This last type is the one we normally care about.

```py
async def work(name: str) -> str:
    logger.info(f"work {name} started")
    await asyncio.sleep(1)
    logger.info(f"work {name} finished")
    return name

async def main():
    coro = work("A")
    res = await coro
    logger.info(res)
```

---

- If a coroutine is never awaited or scheduled using a task, Python will throw a runtime warning.

```py
async def main():
    work("A")

asyncio.run(main())
```

```
RuntimeWarning: coroutine 'work' was never awaited
```

<br>

#### Coroutine lifecycle

The lifecycle of a coroutine can be summarized in four major stages:

- Creation:
    - The moment we call an `async def` function, a coroutine is created.
    - Example: `coro = fetch_data()`. Here `coro` is a coroutine.
    - This coroutine is not executed yet.
- Schedule:
    - A coroutine is scheduled when it is given to the event loop.
    - We can use `await`, `asyncio.create_task()` or `asyncio.gather()` to schedule a coroutine.
    - Now the loop can run it.
- Execution:
    - After a coroutine is scheduled, event loop will eventually pick it up and start executing this function.
    - This execution continues unit `await` statement is encountered.
    - At each `await`, it suspends, allowing other coroutines to run.
- Completion:
    - The coroutine finishes execution and produces a return value (or raises an exception).

---

Example:

```py
async def main():
    coro = fetch_data()        # Created
    result = await coro        # Scheduled → Running → Completed
```

<br>

#### `async def` and `await`

`async def`:
- Declares a coroutine function.
- Always returns a coroutine object when called.
- Must be awaited or scheduled, otherwise it never runs.

---

`await`:
- Suspends the coroutine until the awaited I/O operation completes.
- Allows the event loop to run other tasks. Ensures non-blocking concurrency.
- We can only use `await` inside an `async def` function.

<br>

#### When NOT to use coroutines?

- Coroutines are ideal for:
    - Network I/O (HTTP requests)
    - File I/O
    - Timers
    - Waiting for OS events
- Coroutines are NOT ideal for CPU-bound tasks, such as:
    - Large computations (e.g., hashing, image processing)
    - Compression
    - Machine learning inference
    - Heavy data processing

---

Why not?
- CPU-bound code does NOT `await` while doing computation. Hence, the control is never yield to the event loop.
- It blocks the event loop, blocking all other coroutines.
- Use `ProcessPoolExecutor` for CPU-bound work.

---

Example handling CPU-bound tasks correctly:

```py
import asyncio
from concurrent.futures import ProcessPoolExecutor

def cpu_heavy(n):
    return sum(i*i for i in range(n))

async def main():
    loop = asyncio.get_running_loop()
    result = await loop.run_in_executor(
        ProcessPoolExecutor(), cpu_heavy, 10_000_000
    )
    print(result)
```

<br>
<br>
<br>

### Task

#### What is a Task?

- A Task is a wrapper around a coroutine that allows it to run **concurrently** within the event loop.
- While a coroutine defines *what* to do, a task defines *when* and *how* it is executed by scheduling it on the event loop.
- Tasks enable **cooperative multitasking**: multiple coroutines appear to run simultaneously by yielding control (via `await`).

Key point:
- A coroutine does not run until it is wrapped in a Task or awaited directly.
- A Task runs automatically once scheduled.

<br>

#### How tasks are scheduled?

- When we create a task (e.g., using `asyncio.create_task()`), it is registered with the event loop.
- The event loop places the task into its internal **ready queue**.
- The loop runs tasks until they hit an `await` on an I/O operation, after which the loop switches to another ready task.
- Task switching is **cooperative**, not preemptive. That means switching happens explicitly using `await`.

<br>

#### Creating a Task

- We create a task when we want to run coroutines concurrently and not sequentially.
- We can create a new task using `asyncio.create_task(coro)`. This will create and immediately schedule a new task.

```py
import asyncio

async def work(name: str):
    print(f"work {name} started")
    await asyncio.sleep(1)
    print(f"work {name} finished")

async def main():
    resA = asyncio.create_task(work("A"))
    resB = asyncio.create_task(work("B"))

    # Wait for both tasks to finish
    await resA
    await resB

asyncio.run(main())

# Output:
# work A started
# work B started
# work A finished
# work B finished
```

---

- If we do not create tasks and use `await` instead, the coroutines will run sequentially.

```py
import asyncio

async def work(name: str):
    print(f"work {name} started")
    await asyncio.sleep(1)
    print(f"work {name} finished")

async def main():
    # Await coroutines directly
    await work("A")
    await work("B")

asyncio.run(main())

# Output:
# work A started
# work A finished
# work B started
# work B finished
```

<br>

#### Task cancellation

- Use `task.cancel()` to request cancellation of a task.
- The next time the task hits an `await`, a `CancelledError` is raised inside the task.
- The task can:
    - handle the cancellation internally (catch and cleanup), or
    - allow it to propagate (common case).

Key point:
- Cancellation is **cooperative**. If the coroutine never awaits, it cannot be cancelled.

```py
async def work(name: str):
    try:
        logger.info(f"work {name} started")
        await asyncio.sleep(5)
        logger.info(f"work {name} finished")
    except asyncio.CancelledError:
        logger.info(f"work {name} cancelled")

async def main():
    task = asyncio.create_task(work("A"))
    
    # Cancel task
    await asyncio.sleep(1)
    task.cancel()

    # Wait for task to finish
    await task

asyncio.run(main())

# Output:
# 00:00:01 - work A started
# 00:00:02 - work A cancelled
```

<br>

#### Task states

- A task will always be in one of these 3 states:
    - **Pending**: Task is scheduled but not yet started OR waiting on an await.
    - **Running**: Currently being executed by the event loop.
    - **Done**: Finished normally, raised exception, or was cancelled.

---

- `asyncio` also provides some helpful methods to check a task's state:
    - `task.done()`: Returns True if completed (success, exception, or cancelled).
    - `task.cancelled()`: Returns True if task is cancelled.
    - `task.result()`:
        - Returns task result if task is completed successfully.
        - Raises exception if an exception occured during task execution.
        - Raises `InvalidStateError` if task's result is not available yet.
        - Raises `CancelledError` if task has been cancelled.
    - `task.exception()`:
        - Returns exception if an exception occured during task execution.
        - Returns `None` is no exception occured.
        - Raises `InvalidStateError` if task's result is not available yet.
        - Raises `CancelledError` if task has been cancelled.

```py
task = asyncio.create_task(some_coro())

print(task.done())      # False
await task
print(task.done())      # True
```

<br>

#### Exceptions inside a Task

- If a task raises an exception, that exception is captured by the Task object.
- The exception will be raised only when we:
    - `await` the Task, or
    - inside the done callback
- If we never await that task, Python prints a warning: `Task exception was never retrieved`.

```py
async def work():
    await asyncio.sleep(1)
    raise ValueError("boom!")

async def main():
    task = asyncio.create_task(work())

    try:
        await task
    except ValueError as e:
        logger.info(f"Caught error: {str(e)}")

asyncio.run(main())

# Output:
# 00:00:01 - Caught error: boom!
```

---

Key points:
- A Coroutines raise exceptions immediately when awaited.
- A Task capture the exceptions. It do not crash the event loop.
- Add proper exception handling when awaiting the task or inside the done callback.

<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
<br>

# Async IO Package

<br>
<br>
<br>
<br>
<br>

## Index

- [Event Loop](#event-loop)
- [Coroutine](#coroutine)

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
- Unlike normal functions, calling a coroutine does not execute it immediately. It returns a **coroutine object** which must be awaited or scheduled on the event loop.
- They support **cooperative multitasking**, meaning they voluntarily yield control.

---

Example:

```py
async def fetch_data():
    await asyncio.sleep(1)
    return "done"
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
<br>
<br>

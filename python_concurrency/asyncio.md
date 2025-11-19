# Async IO Package

<br>
<br>
<br>
<br>
<br>

## Index

- [Event Loop](#event-loop)

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

### 

<br>
<br>
<br>
<br>
<br>

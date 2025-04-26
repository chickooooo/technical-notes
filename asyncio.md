Multiprocessing
- True parallelism
- Preferred for CPU bound tasks

Multitherading
- Uses context switching, due to GIL
- Preferred for I/O bound tasks

Asyncio
- It is a single thread, single process design
- It uses cooperative multitasking

Coroutine
- A coroutine is a specialized version of a Python generator function.
- A coroutine is a function that can suspend its execution before reaching return.
- It can indirectly pass control to another coroutine for some time.

Simple asyncio program

<br>

```python
import time
import asyncio

async def count():
    print("one")
    await asyncio.sleep(1)
    print("two")

async def main():
    await asyncio.gather(count(), count(), count())

if __name__ == "__main__":
    start = time.perf_counter()
    asyncio.run(main())
    time_taken = time.perf_counter() - start

    print(f"time taken {time_taken:.2f} seconds")

# Output:
# one
# one
# one
# two
# two
# two
# time taken 1.00 seconds
```

<br>

When executing an asynchronous function, when python encounter an await statement, it passes control back to the event loop. and till the response is ready for the current function, event loop processes some other task.

<br>

- `def` creates a normal function, while `async def` creates a coroutine or an asynchronous generator.
- The keyword `await` passes function control back to the event loop.

<br>

```python
async def g():
    # Pause here and come back to g() when f() is ready
    r = await f()
    return r
```

<br>

- Using `yield` in an `async def` block creates an asynchronous generator, which you iterate over with `async for`.

<br>

```python
async def count_till_3():
    i = 1
    while i <= 3:
        await asyncio.sleep(0.5)
        yield i
        i += 1

async def loop():
    async for i in count_till_3():
        print(i)

async def main():
    await asyncio.gather(loop(), loop())

# Output:
# 1
# 1
# 2
# 2
# 3
# 3
# time taken 1.50 seconds
```

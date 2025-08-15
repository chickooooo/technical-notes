## Index
- [setTimeout Function](#settimeout-function)
- [setInterval Function](#setinterval-function)
- [setTimeout & setInterval Common Questions](#settimeout--setinterval-common-questions)
- [What is an Event Loop?](#what-is-an-event-loop)
- [Key components of Event Loop](#key-components-of-event-loop)
- [How Event Loop works?](#how-event-loop-works)
- [Closure](#closure)
- [Lexical scope](#lexical-scope)

<br>
<br>
<br>

### setTimeout Function

- Executes a function once after a specified delay (in milliseconds).

<br>

**Syntax**

```javascript
const timeoutID = setTimeout(callback, delay, ...args);
```

<br>

**Example**

```javascript
const timeoutID = setTimeout(() => {
  console.log("Runs after 2 seconds");
}, 2000);
```

<br>

**Stop a Timeout**

```javascript
clearTimeout(timeoutID);
```

<br>
<br>
<br>

### setInterval Function

- Executes a function repeatedly at specified intervals.

<br>

**Syntax**

```javascript
const intervalID = setInterval(callback, delay, ...args);
```

<br>

**Example**

```javascript
const intervalID = setInterval(() => {
  console.log("Runs every 1 second");
}, 1000);
```

<br>

**Stop an Interval**

```javascript
clearInterval(intervalID);
```

<br>
<br>
<br>

### setTimeout & setInterval Common Questions

- Both return an ID used to cancel them.
- Delay is in milliseconds (ms).
- Callbacks are asynchronous but not exact in timing (event loop delays possible).

<br>

**Can you simulate setInterval() using setTimeout()?**

- Yes, by recursively calling `setTimeout()`.

```javascript
function repeat() {
  console.log("Repeated task");
  setTimeout(repeat, 1000);
}
setTimeout(repeat, 1000);
```

<br>

**What happens if the callback in setInterval() takes longer than the interval?**

- The next invocation doesn't wait for the previous one to finish.
- This can cause overlapping executions, leading to performance issues or race conditions.

<br>

**Is the delay in setTimeout() or setInterval() guaranteed?**

- No. The delay is minimum time before execution.
- Actual execution may be delayed due to the event loop or call stack.

<br>

**What does setTimeout(fn, 0) do?**

- It delays execution of fn until the call stack is clear.
- It runs after current synchronous code.

<br>

**How do you pass arguments to the callback function in setTimeout()?**

- Use extra arguments after the delay.

```js
setTimeout((msg) => console.log(msg), 1000, "Hello")
```

<br>

**How can you pause and resume setInterval()?**

- You must manually implement pause/resume using `clearInterval()` and `setInterval()` again.

<br>
<br>
<br>

### What is an Event Loop?

- JavaScript runs in a single thread, so only one task can execute at a time.
- If we have a long-running task (e.g., network requests, timers), it would block the main thread and freeze the UI.
- So, event loop is used to coordinates the execution of synchronous code and asynchronous tasks.
- Event loop is the mechanism that keeps JavaScript's single-threaded runtime non-blocking.

### Key components of Event Loop

**Call Stack**

- Executes synchronous code.
- Function calls are added to the stack when invoked.
- And removed from the stack when completed. Removal is LIFO.

<br>

**Web APIs (Provided by the Browser)**

- Handles async tasks like: `setTimeout`, `fetch`, `geolocation`, etc.
- Executes in background, outside the call stack.

<br>

**Task Queue (Callback Queue)**

- Stores callbacks from Web APIs once async tasks complete.
- Example: setTimeout callback.
- These callbacks are handled after the call stack is empty.
- The operation works in FIFO fassion.

<br>

**Microtask Queue**

- Stores callbacks from: `Promises (then, catch, finally)`, `async/await`, `queueMicrotask`.
- Has higher priority than the task queue.
- The event loop will fully drained the microtask queue before moving to the task queue.

<br>
<br>
<br>

### How Event Loop works?

- Start executing synchronous code in the call stack.
- If an async function (like `fetch` or `setTimeout`) is called, it is passed to the Web API.
- When the async function is done, its callback is queued (task queue or microtask queue).
- Once the call stack is empty, the event loop, theevent loop first checks the microtask queue and runs all microtasks.
- Once microtask queue is empty, it then picks one task from the task queue and runs it.
- This cycle continuously operates in a loop.

<br>

```js
console.log(1);

setTimeout(() => console.log(2), 0);

Promise.resolve().then(() => console.log(3));

console.log(4);

// Output: 1 → 4 → 3 → 2
```

<br>

- The delay in `setTimeout` is the time to enqueue the task, not to execute it.
- Microtasks can delay tasks indefinitely if they keep queuing more microtasks.

<br>

<img src="assets/event_loop.png" alt="Event Loop" width="400"/>

<br>
<br>
<br>

### Closure

- A closure is a function that is defined inside of another function.
- The inner function has access to the variables and scope of the outer function.
- Using closures allow us to:
  - Simulate private variables.
  - Persist state across function calls.

<br>

```js
function createCounter() {
    let counter = 0;

    function increament() {
        counter++;
    }

    function getValue() {
        return counter
    }

    return { increament, getValue }
}

let counter = createCounter();

counter.increament();  // 1
counter.increament();  // 2

console.log(counter.getValue());  // Output: 2
```

<br>
<br>
<br>

### Lexical scope

- Lexical scope means: "Where you write your code determines what variables you can access."
- Scope is decided when you write the code, not when you run it.
- A function can see variables:
  - Inside itself
  - In the place where it was defined (not where it’s called)

<br>

```js
function outer() {
  let name = "Alex";

  function inner() {
    console.log(name); // can access 'name'
  }

  inner();
}
outer();
```

- `inner()` can access `name` because it was written inside `outer()`.

<br>
<br>
<br>

### 

<br>
<br>
<br>

- var vs let vs const
- == vs ===

- promises
- event delegation
- fetch vs axios

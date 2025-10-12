# Concurrency

<br>
<br>
<br>
<br>
<br>

## Goroutine

- A goroutine is a lightweight thread managed by Go runtime.
- It is basically a function or a method, running concurrently with other goroutines.
- Goroutines are cheaper than OS threads in terms of memory and scheduling overhead. That means we can have thousands or even millions of them running concurrently.

<br>
<br>

### Create a goroutine

```go
go someFunction()
```

- We can create a goroutine by using the `go` keyword before a function call.
- This starts the execution of `someFunction()` immediately in a separate lightweight thread.
- This goroutine will run concurrently with the main program and other goroutines.
- The main program continues it's execution and doesn't wait for `someFunction()` to finish.

<br>
<br>

### Main goroutine

- The `main` function itself runs in the main goroutine.
- If the main goroutine exits, all other goroutines are terminated immediately.

<br>
<br>
<br>

## Channel

<br>
<br>
<br>

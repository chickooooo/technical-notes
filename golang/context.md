# `context` Package

<br>
<br>
<br>
<br>
<br>

## Index

- [What is context?](#what-is-context)
- [`context.Background()`](#contextbackground)
- [`context.TODO()`](#contexttodo)
- [`context.WithCancel()`](#contextwithcancel)
- [`context.WithTimeout()`](#contextwithtimeout)
- [`context.WithDeadline()`](#contextwithdeadline)
- [`context.WithValue()`](#contextwithvalue)
- [`ctx.Err()`](#ctxerr)

<br>
<br>
<br>
<br>
<br>

## What is context?

- A `context` is used to carry **deadlines**, **cancellation signals** and **request scoped values** across goroutines.
- It is particularly useful for:
    - Managing timeouts.
    - Handling cancellations.
    - Passing values (request ID, user, etc.) across function calls.
- It is used to prevent goroutine leaks.

<br>
<br>
<br>

## `context.Background()`

- It is a base context in Go.
- It returns an empty context that has no values, no deadline and is never cancelled.
- Typically used as the root context inside `main()` and tests.
- Other contexts like timeouts, cancellations, etc. make use of this base context.

```go
cxt := context.Background()
```

<br>
<br>
<br>

## `context.TODO()`

- It is a placeholder context in Go.
- It is used when we aren't sure which context to use or if the context is not yet available in that scope.
- It should be replaced later with the actual context value.

```go
ctx := context.TODO()
```

<br>
<br>
<br>

## `context.WithCancel()`

- It creates a new context that can be cancelled manually.
- It returns:
    - A derived context that inherits from the parent context.
    - A `cancel` function to manually close the context.
- When `cancel()` is called, a signal is pushed to `ctx.Done()` channel.

```go
baseCtx := context.Background()
ctx, cancel := context.WithCancel(baseCtx)

go func() {
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Done")
    case <-ctx.Done():
        fmt.Println("Cancelled")
    }
}()

time.Sleep(2 * time.Second)
cancel()

time.Sleep(100 * time.Millisecond)

// Output:
// Cancelled
```

<br>
<br>
<br>

## `context.WithTimeout()`

- It creates a new context that automatically cancels after the specified timeout.
- It returns:
    - A derived context that inherits from the parent context.
    - A `cancel` function, which should be deffered to avoid context leak.
- Useful for setting time limit on operations like database calls, HTTP requests, etc.

```go
baseCtx := context.Background()
ctx, cancel := context.WithTimeout(baseCtx, 2*time.Second)
defer cancel() // clean up the context

go func() {
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Done")
    case <-ctx.Done():
        fmt.Println("Cancelled")
    }
}()

time.Sleep(2500 * time.Millisecond)

// Output:
// Cancelled
```

<br>
<br>
<br>

## `context.WithDeadline()`

- Similar to `context.WithTimeout()`, but instead of cancelling after a timeout duration, it cancels at the specified date & time.
- Uses `time.Time` instead of `time.Duration`.

```go
baseCtx := context.Background()
ctx, cancel := context.WithDeadline(baseCtx, time.Now().Add(5*time.Second))
defer cancel() // clean up the context

go func() {
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Done")
    case <-ctx.Done():
        fmt.Println("Cancelled")
    }
}()

time.Sleep(3500 * time.Millisecond)

// Output:
// Done
```

<br>
<br>
<br>

## `context.WithValue()`

- It returns a copy of the parent context and adds key-value pair to it.
- Used to pass **request-scoped data** like user ID, auth token, etc.

<br>

Key points:
- Prefer using a custom type for keys to avoid collisions.
- Avoid overusing it for passing optional params. It is not meant to replace function arguments.

```go
type MyKey string

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, MyKey("abc"), 123)

	fmt.Println(ctx.Value(MyKey("abc")))  // 123
	fmt.Println(ctx.Value(MyKey("abcd"))) // <nil>

	fmt.Println(baseCtx.Value(MyKey("abc"))) // <nil>
}
```

<br>
<br>
<br>

## `ctx.Err()`

- Returns the reason why the context was cancelled or done.
- It returns:
    - `nil`: If context is still active.
    - `context.Canceled`: If canceled manually (e.g., via `cancel()`).
    - `context.DeadlineExceeded`: If the deadline or timeout was hit.

```go
func doSomething(ctx context.Context) (string, error) {
    select {
    case <-time.After(3 * time.Second):
        return "done", nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}
```

<br>
<br>
<br>

## 

<br>
<br>
<br>

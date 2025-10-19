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

- In Go, a channel is a medium that allows communication between goroutines.
- It acts as a pipe through which goroutines can send and receive values **synchronously**.
- Channels help **coordinate execution** and **share data** without using shared memory directly.

<br>
<br>

### Create a channel

#### Declaring a channel

```go
var myChannel chan int

fmt.Println(myChannel == nil) // true
```

- This declares a channel of type int, but it is not initialised yet.
- Channels are typed, so `myChannel` can only transport `int` values.

<br>

#### Initialising a channel

```go
myChannel = make(chan int)
```

- Use `make()` to allocate and initialise a channel.
- If no second argument is passed to `make()`, it creates an **Unbuffered Channel**.

<br>
<br>

### Send and receive data

```go
ch := make(chan int)

// Sending a value to channel
ch <- 42

// Receiving a value from channel
value := <-ch
```

- Use `ch <- {val}` syntax to send a value to a channel.
- Use `{val} <-ch` syntax to receive a value from a channel.

<br>

- For an unbuffered channel or a buffered channel with full capacity, both send and receive operations are blocking in nature.
- That means, the sender waits till there is someone to receive the data at the other end.
- Similarly, the receiver waits till there is a value to receive.

<br>
<br>

### Pass channel to a goroutine

#### Read + write access

```go
package main

import (
	"fmt"
)

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine!"
}

func main() {
	ch := make(chan string)
	go sendMessage(ch)

	msg := <-ch
	fmt.Println(msg) // Hello from goroutine!
}
```

- In the above example, the `sendMessage()` function can read as well as write to `ch` channel.

<br>

#### Write only

```go
func sendOnly(ch chan<- int) {
	ch <- 10
}
```

- Use `chan<- type` syntax to give write only acces to the goroutine.

<br>

#### Read only

```go
func receiveOnly(ch <-chan int) {
	fmt.Println(<-ch)
}
```

- Use `<-chan type` syntax to give read only acces to the goroutine.
- This improves clarity and safety in concurrent programs.

<br>
<br>

### Buffered and Unbuffered channel

#### Buffered channel

```go
ch := make(chan int, 2) // buffered channel with capacity 2
```

- Buffered channels have a fixed capacity and can hold upto `n` items.
- While sending a new item to the channel, if the channel is full, the execution is blocked till a slot is emptied.
- While receiving from a channel, if the channel is empty, the execution is blocked till a new item is received.

<br>

#### Unbuffered channel

```go
ch := make(chan int) // unbuffered channel
```

- Unbuffered channels holds no data.
- In order for data to flow, both sender and receiver should be connected to the channel.
- If either one is not connected, the execution is blocked.
- Unbuffered channels facilitate synchronous data transfer.

<br>
<br>

### Close a channel

```go
close(ch)
```

- Use built-in `close()` function to close an open channel.
- On the receiver end, we can check if the channel is open before receiving the data.

```go
data, ok := <-ch

if !ok {
	fmt.Println("channel is closed")
}
```

- If `ok == true` the channel is open, otherwise the channel is closed. 

<br>

#### Closed channel behaviour

- Only sender should be allowed to close a channel.
- Sending data on a closed channel causes a panic.
- Receiving data from a closed channel will return zero value of that channel's type.

<br>
<br>

### Loop over a channel

```go
for item := range ch {
	fmt.Println(item)
}
```

- This will keep receiving data until the channel is closed.

<br>
<br>
<br>

## Select statement

- A `select` statement is similar to a `switch` statement, but for channels.
- It allows a goroutine to manage multiple channel operations.

<br>

Key features:
- It blocks until one of its cases can proceed.
- If multiple cases are ready, one is chosen at random (pseudo random).
- It also supports `default` and timeouts.
- Only 1 case of `select` statement gets executed. No multiple cases execution.

<br>
<br>

### Syntax

```go
select {
	// Waiting to receive data from channel ch1
	case val := <-ch1:
		fmt.Println("Received data:", val)

	// Waiting to send data to channel ch2
	case ch2 <- "data":
		fmt.Println("Sent data to ch2")
	
	// If neither ch1 & ch2 are ready
	default:
		fmt.Println("No communication. Moving on.")
}
```

- Case 1: `val := <-ch1`. This case gets triggered when we receive data on `ch1`.
- Case 2: `ch2 <- "data"`. This case gets triggered if the channel ch2 is ready to receive data immediately. When selected, the data "data" is sent into ch2, and the associated code block executes.
- Case 3: `default`. This case gets triggered when neither `ch1` nor `ch2` are ready for operation. In this case, the message is logged and the program continues its execution.

<br>
<br>

### Basic select

```go
select {
case msg1 := <-ch1:
	fmt.Println(msg1)
case msg2 := <-ch2:
	fmt.Println(msg2)
}
```

- Here we are waiting to receive data either on `ch1` or `ch2`.
- The channel from which the data is received first, that case gets selected. Other cases is ignored.
- If both the channels have data on them, any one is selected at random.
- Note: Only 1 case is chosen for execution. There is no such scenario where both cases will be chosen.

<br>
<br>

### Select with default

```go
select {
case val := <-ch:
	fmt.Println("Received:", val)
default:
	fmt.Println("No value received")
}
```

- When the execution of `select` statement starts, if there is no data on `ch` to be received, the `default` case is executed.
- The statement inside the `default` case is logged and the program execution continues post `select`.

<br>
<br>

### Select with timeout

```go
select {
case msg := <-ch:
	fmt.Println("Received:", msg)
case <-time.After(2 * time.Second):
	fmt.Println("Timeout!")
}
```

- If we don't receive data on channel `ch` within 2 seconds, the timeout case is executed.
- In general, if no other case of `select` statement is selected before the timeout, the timeout case is executed.

<br>
<br>
<br>

## Goroutine leak

- A goroutine leak occurs when a goroutine is started but never exits.
- This usually happens when a goroutine is blocked indefinitely either over a channel or an external resource.
- These blocked goroutines consume memory and other system resources, leading to performance degradation over time.

<br>
<br>

### Common goroutine leaks

#### Blocked channel

- A goroutine is waiting to read/write to a channel that never sends/receives data.

```go
func leaky() {
    ch := make(chan int)
    go func() {
        <-ch // blocks forever
    }()
}
```

<br>

#### Infinite select

- Not exiting from an infinite `for` loop + `select` combo.

```go
func leaky2() {
    ch := make(chan int)
    go func() {
        for {
            select {
            case data := <-ch:
                fmt.Println(data)
            }
        }
    }()
}
```

<br>

#### Cancelled work

- When a task that a goroutine is performing is cancelled by the user.

```go
func leaky3(ch chan<- int) {
	// doing work
	data := doWork()

	// send data to the client
	// but the client is not there to receive
	ch <- data
}
```

<br>
<br>

### Solutions for goroutine leaks 

#### `done` channel

- Use a `done` channel inside `select` statement to indicate that data transfer has completed.

```go
func safe() {
    ch := make(chan int)
	done := make(chan struct{})

	go func() {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-done:
			fmt.Println("Stopping goroutine")
			return
		}
	}()

	fmt.Println(runtime.NumGoroutine()) // 2
	close(done)

	time.Sleep(100 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine()) // 1
}
```

<br>

#### Context package

- Use features provided by `context` package for safely exiting goroutine.
- It can handle cancellations, timeouts, etc.

```go
func worker(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		case data := <-ch:
			fmt.Println("Got:", data)
		}
	}
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, ch)

	ch <- 1
	ch <- 2

	fmt.Println(runtime.NumGoroutine()) // 2
	cancel()                            // this will stop the goroutine

	time.Sleep(100 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine()) // 1
}
```

<br>

Other solutions include using timeouts (`Time.After()`) and `default` case.

<br>
<br>

### Detect goroutine leaks

- Get number of active goroutine using `runtime.NumGoroutine()`.
- Use `pprof` or `runtime/pprof` to profile goroutines.

<br>
<br>
<br>

## 

<br>
<br>
<br>

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

## 

<br>
<br>
<br>

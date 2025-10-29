# `sync` Package

<br>
<br>
<br>
<br>
<br>

## Index

- [`sync` Package](#sync-package-1)
- [Mutex](#mutex)
- [Read/Write Mutex](#readwrite-mutex)

<br>
<br>
<br>
<br>
<br>

### `sync` Package

- In Go, `sync` package provides fundamental primitives for concurrent programming.
- It provides tools for safe access to shared data and coordination between goroutines.

<br>
<br>
<br>

### Mutex

- Mutex is short form for mutual exclusion.
- A mutex `sync.Mutex` prevents multiple goroutines from accessing the shared resource at a time.
- It prevents **race condition** by ensuring only 1 goroutine can enter the critical section at a time.

```go
var (
    counter int
    mut      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()

    mut.Lock()   // Lock before accessing shared data
    counter++
    mut.Unlock() // Unlock after finishing
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
```

- `mut.Lock()` blocks other goroutines from entering the critical section.
- `mut.Unlock()` releases the lock, allowing others to proceed.
- If we forgot to unlock a resource, it will cause a deadlock.
- A good practice is to use `defer mut.Unlock()`, right after locking the resource.

<br>
<br>
<br>

### Read/Write Mutex

- A Read/Write Mutex (`sync.RWMutex`) is a special kind of lock that distinguishes between readers and writers.
    - While reading, multiple readers can access the shared resource, but not the writer.
    - While writing, only a single writer can access the shared resource and no other reader or writer.
- This makes `RWMutex` ideal for scenarios where reads are much more frequent than writes.

```go
var shared = make(map[string]int)
var rwMut sync.RWMutex

func write(wg *sync.WaitGroup) {
	defer wg.Done()

	rwMut.Lock()

	fmt.Println("Writing")
	time.Sleep(2 * time.Second)
	shared["apple"] = 1
	shared["banana"] = 2
	fmt.Println("Done writing")

	rwMut.Unlock()
}

func read(wg *sync.WaitGroup, key string) {
	defer wg.Done()

	rwMut.RLock()

	fmt.Println("Reading")
	time.Sleep(1 * time.Second)
	fmt.Printf("Value of %s is: %d\n", key, shared[key])
	fmt.Println("Done reading")

	rwMut.RUnlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go write(&wg)                     // Write to shared resource
	time.Sleep(time.Millisecond * 50) // Some delay (for locking to take effect)
	go read(&wg, "apple")             // Read from shared resource
	go read(&wg, "banana")            // Read from shared resource

	wg.Wait()
	fmt.Println("Done!")
}
```

- In the above example, while writing (`rwMut.Lock()`), the reads will be blocked.
- But when reading (`rwMut.RLock()`), both reads will happen simultaneously.
- If we use `.Lock()` instead of `.RLock()` in the `read()` function, the reads will happen one after another and not simultaneously.

<br>
<br>
<br>

### WaitGroup

- A `WaitGroup` is used to block the main goroutine, until a set of goroutines finish their execution.
- It makes use of `.Add()`, `.Done()` and `.Wait()` methods to manage synchronization.

<br>

#### Working

- A `WaitGroup` maintains a counter of how many goroutines are currently running.
- We `.Add(1)` to the counter when we start a new goroutine.
- Each goroutine calls `.Done()` when it finishes.
- The main goroutine calls `.Wait()`, which blocks until the counter drops back to zero.

<br>

#### Example

```go
func worker(id int, wg *sync.WaitGroup) { // Accept pointer
	defer wg.Done() // Decrement counter

	fmt.Println("Started job:", id)
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	fmt.Println("Finished job:", id)
}

func main() {
	var wg sync.WaitGroup // Create waitgroup

	for i := range 5 {
		wg.Add(1) // Increment counter
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Done")
}
```

<br>

#### Important notes

- Always accept pointer (`*WaitGroup`) in the goroutine. If value is passed, a copy of that waitgroup is created.
- Always call `.Add(1)` outside of the goroutine and before starting the goroutine. This avoids race condition.
- `WaitGroup` is not reusable for overlapping goroutine sets. Always ensure all goroutines are done before reusing it.

<br>
<br>
<br>

### Once

<br>
<br>
<br>

### Map

<br>
<br>
<br>

### 

<br>
<br>
<br>

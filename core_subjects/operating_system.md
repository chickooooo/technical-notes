## Operating System




What is an operating system?
- A software that acts as an interface between user applications and computer hardware.
- The primary tasks performed by OS are:
    - Process management
    - Memory management
    - File system management
    - Device management, etc.

- Linux and MacOS are developed on top of Unix OS.
- Windows, iOS, Android, etc.
- RTOS (real time operating system) is used for systems that need realtime response. Example: embedded systems, IoT devices, missiles, satellites, etc.

<br>
<br>
<br>

### Process Management

Programs:
- A program is a set of instructions written in a programming language to perform a certain task.
- Example: A python program to convert HTML to PDF.

Process:
- A process is an instance of a program in execution.
- When a program is loaded in memory and executed, it becomes a process.
- A process is an individual entity with it's own memory space and a set of resources.
- Processes provide isolation and protection between different instances of a program.
- Processes are managed by the OS and can communicate with each other through IPC.

Thread:
- A thread is a lightweight unit of execution within a process.
- It contains a sequence of instructions that can be executed independently.
- Thread share the same memory space and resources within a process.
- They are faster to create and switch between than processes.
- A process will always have a main thread and can spawn more threads if required. 


Multiprogramming:
- Multiprogramming is a technique where multiple programs are loaded into memory simultaneously.
- CPU switches between these programs to execute instructions.
- This maximizes CPU utilization and keep the CPU busy by quickly switching between different programs when one is waiting for I/O.

Multitasking:
- The ability of an OS to run multiple tasks (which can be processes or threads) seemingly at the same time.
- Achieved by rapidly switching the CPU's attention between different tasks.
- On a single core, it is acheived using concurrency. On a multi-core system, it can involve parallelism.
- It is the broader concept that enables both multiprocessing and multithreading.

Multiprocessing:
- Multiprocessing refers to the execution of multiple processes simultaneously.
- Multiprocessing increases system throughput by running each process on a separate core.

Multithreading:
- Multithreading involves executing multiple threads within a single process.
- These threads share the same resources within a process.
- Tey can run parallely each on a separate core.

Summary:
- Multitasking is the overall goal of doing many things at once. Multiprocessing and multithreading are two different ways (using multiple processes or multiple threads within a process, respectively) to achieve multitasking, especially on modern multi-core processors.

FAQs
- A quad-core processor indicates that the CPU has 4 independent processing cores.
- A single process can run on multiple cores if it is designed to do so:
    - If the process has multiple threads spawned, the OS will run each thread on a separate CPU core (it that core is free).
    - If the process has only 1 main thread, then only 1 CPU core is used by that process.
- A python process with 1 main thread + 3 spawned threads will utilise only 1 CPU core and use context switching between threads. Due to GIL.


Context switching:
- Context switching is the process where the CPU switches from one process / thread to another.
- This increases CPU utilization by switching processes that are waiting for IO.
- This allows multiple processes to share a single CPU core.
- It looks like multiple processes are running simultaneously.

Concurrency:
- Concurrency is the ability of the system to handle multiple tasks at the same time.
- It uses context switching to switch between tasks.
- When 2 tasks are running concurrently, that means only 1 is getting executed at a time.
- The OS just switch between them so fast that it seems like both are running at the same time.
- Often used for IO bound tasks. Reduces CPU ideal time.

Parallelism:
- Parallelism is the ability of the system to **execute** multiple tasks at the same time.
- This involves executing tasks parallely on a multi-core system. 
- Often used for CPU bound tasks. Reduces computation time by dividing the workload.


Various process states:
- New: The process is being created.
- Ready: The process is ready to run but is waiting for CPU scheduling.
- Running: The process is currently being executed by the CPU.
- Waiting / Bocked: The process is waiting for some event (like I/O) to complete.
- Terminated / Exit: The process has finished execution (either successfully or due to an error).
- Suspended (Ready/Blocked): The process is temporarily removed from memory and stored on disk (swapped out), either in the ready or blocked state. This happens when the system is low on RAM and needs to free up memory for other processes.

```
New -> Ready -> Running -> Terminated
        |         |
        ------- Waiting
        |         |
        ------- Suspended
```


CPU scheduling algorithms:

Preemptive & non-preemptive approach:
- The CPU can be taken away from a process by the scheduler, even if the process is not finished.
- Once a process gets the CPU, it keeps it until it finishes or blocks itself.

First come, first served (FCFS):
- The process that arrives first is executed first.
- This algorithm is non-preemptive in nature.
- If a long process is running, several short processes have to wait for long.
- No support for process prioritization.

Shortest job next / shortest job first (SJN / SJF):
- The process with the shortest burst time (execution time) is executed first.
- Non-preemptive by default.
- Can be made preemptive, called **Shortest remaining time first (SRTF)**.
- It minimizes average waiting time.
- Processes with longer burst time have to wait a lot.

Shortest remaining time first (SRTF):
- SRTF is the preemptive version of Shortest Job First (SJF).
- The CPU always runs the process with the shortest remaining time.
- If a new process arrives with a shorter remaining time, it preempts the current process.
- Advantages & drawbacks are the same.

Round robin (RR):
- Preemptive scheduling algorithm.
- It assigns a fixed time (e.g. 10 miliseconds) to each process in a circular manner.
- Once a process exhausts its time, it is moved to the back of the ready queue, allowing the next process in line to execute.
- Provides fair execution time to all processes.
- Suffers from high context-switching overhead.
- Not efficient for long-running processes.

Priority scheduling:
- Assigns priority to each process.
- CPU is allocated to the process with the highest priority.
- It can preemptive or non-preemptive in nature.
    - In preemptive approach, if a higher priority process arrives, current process will be preempted.
    - In non-preemptive approach, the current process continues executing until it completes or voluntarily gives up the CPU.
- Processes with lower priority have to wait a lot.

Multilevel queue scheduling:
- It divides the ready queue into multiple priority queues.
- Each queue will have it's own scheduling algorithm.
- The CPU is allocated to processes in the highest non-empty queue.
- Processes do not move between queues.
- Simple to implement, but inflexible due to static queue assignment.
- Suffers from starvation in lower-priority queues if higher-priority queues are always busy.

Multilevel feedback queue scheduling:
- Advanced version of multilevel queue scheduling where processes can move between queues.
- Initially, all processes enter the highest-priority queue.
- If a process uses too much CPU time, it is moved to a lower-priority queue (i.e., longer jobs are "demoted").
- Short jobs tend to stay in higher-priority queues, getting faster service.
- Reduces starvation by giving lower-priority processes a chance over time.
- Often used in real-world operating systems like Windows and Linux.


Inter Process Communication (IPC)
- Two processes communicate with each other using IPC.
- IPC allows processes to exchange data, signals and messages either in a same system or over a network.
- Common IPC mechanisms:
    - Shared memory
    - Message passing
    - Pipes
    - Sockets

Shared memory:
- Two or more processes share a region of memory.
- One process writes to the memory, the other reads from it.
- Fastest method sicnce it avoid kernel calls.
- Requires synchronisation mechanism to avoid conflicts.
- Pros:
    - Very fast.
    - Efficient for large amounts of data.
- Cons:
    - Complex synchronization.
    - Only works between processes on the same machine.

Message passing:
- Processes send and receive messages using OS provided interfaces.
- Can be synchronous (blocking) or asynchronous (non-blocking).
- Suitable for communication between local or remote processes.
- Example:
    - Local: pipes, signals
    - Remote: TCP, UDP, RPC
- Pros:
    - Simple to use as managed by OS.
    - Works for local as well as remote processes.
- Cons:
    - Slower than shared memory due to kernel overhead.

Pipes:
- It is a unidirectional data channel.
- Data is written at one end (by writer process) and read at the other end (by reader process).
- Pros:
    - Preferred for parent-child process communication.
- Cons:
    - Unidirectional unless two pipes are used.

Sockets:
- Allows communication between processes over a network (or locally).
- Mostly used in web services.
- Supports TCP (reliable) and UDP (fast, no guarantee).
- Pros:
    - Local and remote communication support.
    - Flexible: provides protocols like TCP & UDP.
- Cons:
    - Resource intensive.
    - Complex implementation.


Critical section problem:
- In concurrent programming, when multiple processes or threads access shared resources (e.g. variable, file, memory, etc.) there is a risk of race condition.
- To handle this, we define a critical section

Critical section:
- It is a part of the code where a shared resource is accessed.
- Only one process/thread should execute in this critical section at a time to avoid data inconsistency.

Process synchronization:
- It refers to techniques & mechanisms used to coordinate the execution of processes/threads, so they can work together harmoniously.
- It prevents issues like race conditions, data inconsistency and deadlocks.
- Each process synchronization mechanism must fulfill these 3 requirements:
    - Mutual Exclusion:
        - Only one process/thread should be able to access shared resource or enter critical section at a time.
        - This prevents data inconsistency.
    - Progress:
        - It should allow to make progress by making sure at least 1 process/thread can enter the critical section.
        - It avoids situation where all processes/threads are blocked leading to deadlock.
    - Bounded Waiting:
        - It should make sure that a process/thread waiting to enter the critical section will eventually be allowed to do so.
        - It prevents process/thread from waiting indefinitely.

Process synchronisation:

- Lock / Mutex (mutual exclusion):
    - It is a binary lock.
    - It allows a process/thread to lock the shared resource while in use.
    - While a shared resource is locked, other processes/threads cannot access it.
    - Others trying to acquire the lock must wait (block).

- Semaphore:
    - It can allow more than 1 process/thread to access a shared resource.
    - It has 2 types:
        - Binary semaphore: Like mutex, only 0 or 1.
        - Counting semaphore: Allows upto N threads to access the shared resource.
    - Useful when managing limited sources (e.g. 3 database connections).

- Read-write lock:
    - It allows multiple readers to access the shared resource simultaneously.
    - But only 1 writer can write to the resource at a time.
    - More performant that mutex (which allows only 1 reader/write at a time).

Deadlock:
- A deadlock is a situation where a set of processes are blocked because each process is waiting for a resource held by another, forming a cycle of dependencies.
- "Process A is waiting for Process B, Process B is waiting for Process C, ... and eventually one is waiting for Process A again. None can proceed."

Necessary conditions for a deadlock (Coffman's conditions):
- Mutual exclusion: Resources must be non-shareable (only once process can use it at a time).
- Hold and wait: A process holding a resource is waiting for another.
- No preemption: Resources cannot be forcefully taken away. They must be released voluntarily.
- Circular wait: The waiting processes form a circular chain. 

Deadlock handling techniques:

- Deadlock prevention:
    - Aims to eliminate the necessary conditions for a deadlock.
    - By ensuring that one or more of these conditions are not satisfied, deadlocks can be prevented from happening in the first place.

- Deadlock avoidance:
    - Deadlock avoidance techniques use resource allocation algorithms to avoid situations that may lead to deadlock.
    - These techniques involve resource allocation graph, **Bankers algorithm**, etc.
    - The idea is to have a system that can predict whether granting a resource request will lead to a potential deadlock.
    - If a request might cause a deadlock, it is delayed until granting it will not cause any issues.

- Deadlock detection:
    - It involves periodically checking the resource allocation state to determine if a deadlock has occured.
    - It uses algorithms such as resource allocation graph, **Bankers algorithm**, etc.
    - When a deadlock is detected, deadlock recovery techniques are used.

- Deadlock recovery:
    - Deadlock recovery techniques are used to recover from a deadlock once it has been detected.
    - It involves:
        - Terminating one or more processes to break the circular chain.
        - Release the resources held by these processes.
        - Rolling back the system to a safe state.

- Deadlock ignorance:
    - This involves ignoring the deadlock entirely.
    - This technique assumes that deadlocks will occur rarely.
    - This saves the system resources that would have been used for deadlock handling techniques. Making system more performant.


Banker's Algorithm:
- It is a deadlock avoidance algorithm.
- It checks whether allocating a resource to a processs will lead the system to be in a safe state or not.
    - If the does, it will allocate the resource.
    - Otherwise, it will reject the current request and try again later.
- It’s called the **Banker's** algorithm because it's similar to how a banker makes decisions about loaning money: only loans that are guaranteed to be repaid eventually are approved.
 
<br>
<br>
<br>

### Memory Management

What is memory management in OS?
- Memory management is the functionality of the OS to manage primary memory (RAM).
- The included tasks are:
    - Memory allocation and deallocation.
    - Tracking which parts of memory is in use.
    - Ensuring isolation between processes.
    - Optimising performance using techniques like **paging**, **segmentation** and **caching**.
- There are 2 broad categories used for memory management:
    - Fixed partitioning
    - Dynamic partitioning


Fixed Partitioning:
- Here the memory is divided into fixed sized partitions or blocks (e.g. 100 bytes).
- Each partition can hold exactly one process.
- Suppose a process requires 1 KB of memory, 10 partitions are assigned to it.
- Advantages:
    - Relatively simple to implement.
    - Provides fast memory allocation.
- Disadvantages:
    - Prone to inefficient memory utilization due to **Internal Fragmentation**.

---

Fixed Partitioning
```
+-----------+-----------+-----------+-----------+-----------+
| Partition | Partition | Partition | Partition | Partition |
|   100 B   |   100 B   |   100 B   |   100 B   |   100 B   |
+-----------+-----------+-----------+-----------+-----------+

```

---

Internal Fragmentation
```
Memory Partition (100 bytes):  [======================..........]
                                ↑                     ↑
                               70B                 30B wasted
                          (Used by process)    (Internal Fragmentation)
```

<br>

Dynamic Partitioning:
- Also known as variable partitioning.
- Here memory is not pre-divided into partitions.
- Allows memory to be allocated and deallocated dynamically based on the requirements of the process.
- Advantages:
    - No internal fragmentation.
    - It provides better memory utilization as compared to fixed partitioning.
- Disadvantages:
    - More complex management than fixed partitioning.
    - Suffers from **External Fragmentation**.

---

Dynamic Partitioning:
```
+---------+-------------------+-------------+-----------------+
|   P1    |        P2         |     P3      |   Free (150B)   |
|  (70B)  |      (180B)       |   (100B)    |                 |
+---------+-------------------+-------------+-----------------+
```

---

External Fragmentation
```
+---------+-------------------+-------------+-----------------+
|   P1    |   Free (180B)     |     P3      |   Free (150B)   |
+---------+-------------------+-------------+-----------------+
```
- A new process P4 needs 200B, but even though there’s 330B free total, P4 cannot be allocated because there’s no single 200B block available.

<br>

Memory allocation algorithms in dynamic partitioning

First Fit:
- The first-fit algorithm allocates the first available memory block that is large enough to accomodate the process.
- This algorithm is simple and provides fast allocation.
- It may result in larger leftover fragments.

Best Fit:
- The best-fit algorithm searches for the smallest available memory block that is large enough to accomodate the process.
- It aims to minimise the leftover fragments.
- It leads to overall better memory allocation, but takes more time than First Fit.

Worst Fit:
- The worst-fit algorithm allocates the largest available memory block to the process.
- It assumes that it will leave a large enough hole to be reused later.
- It aims to reduce memory fragmentation caused by small processes.














<br>
<br>
<br>
<br>
<br>

- what are signals in OS (e.g. SIGKILL, SIGTERM, etc.)
- what is core
- what is cpu
- what is vcpu

- File management
- Most asked OS interview questions

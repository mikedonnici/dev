# Concurrency in Go

_Concurrency_ and _parallelism_ or _parallel execution_ are closely related, but not the same.

_Parallel execution_ is two programs running at exactly the same time - this requires two processors / cores.

Parallel execution does not speed up an individual task but gives better overall throughput.

The amount of benefit depends on the types of tasks being performed. Using a dishwashing metaphor: you cannot both wash and dry a dish at the same time.

Conversley, multi-core systems are only exploited when the programs running on them are designed for parallel execution.

Improvements in hardware speed are slowing due to physical limitations.

So, _concurrent_ programming is used to exploit multi-core systems and achieve improved performance.

Concurrent code tells the computer that it is ok to run certain instructions in parallel, if it is able to do so.

**Concurrent vs Parallel**

Concurrent execution **may** run in parallel (hardware permitting), but not necessarily. In concurrent execution the times that multiple tasks are actually running overlaps. However, it may be such that one task is _paused_ while another starts up. The processor time is divided up amongst the processes until they are all done. In effect, they are all _in progress_ concurrently but not necessarily _running_ concurrently.

By contrast, parallel execution means that multiple tasks are _actually running_ at the same time. There is no pause in one so another can have processor time. Hence, this requires seperate physical hardware.

The process of deciding which core particular tasks will run on, and if they will be parallel or concurrent, is not controlled by the programmer. It is controlled by the operating system and hardware and the Go runtime scheduler.

The programmer just decides what _can_ be done in parallel.

**Concurency without Parallelism**

Can still get improved performance with concurrent code on a single core. This is because the latency of other operations can be hidden. For example, memory operation, screen output, keyboard input, network activity and so on.

So, while waiting on other operations, the program can move on to the next task. In simple terms, it is just using waiting time constructively.

## Processes

A _process_ is an instance of a running program.

A process has:

- Memory

  - Virtual address space
  - Code and access to shared libraries
  - Stack, heap

- Registers (small, single word pieces of memory)

  - program counter, data registers, stack pointers

An _operating system_ essentially allows a lot of processes to execute concurrently. It manages the processes to ensure they don't intefer with eachother, and get fair use of system resources. This process is called _**scheduling**_.

For example, in Linux systems each process is generally allowed 20ms of processor time before the next processes gets a turn.

On a single-core system he processes are not _actually_ running in parallel, however they appear as they though they are because of scheduling.

There are many types of scheduling algorithms that allocate processor time based on _priorities_.

A **context switch** occurs when the OS switches from one process to another. During a context switch the _state_ (ie the context) is maintainedin memory for when the process is restarted. During a context switch it is the OS kernel code that is running, until the context has changed and the new process can run for its scheduled time.

Context-switching takes time - writing the context to memory and then reading the next context from memory into registers. Memory access can be slow.

## Threads and Goroutines

Threads are light-weight processes that share some common context. With this architecture changing threads still requires some context change, however the share part of the context does not have to be changed out for each thread. Thus, the _amount_ of context change is reduced so it is faster.

**Goroutines** are like threads. Multiple goroutines can be run within a main operating system thread. From the OS point of view it is executing a single thread, but within that thread Go is scheduling its own (sub)threads - goroutines.

This is manged by the **Go Runtime Scheduler**. This is like a mini OS and uses its own _logical processor_. By default, it uses a single logical processor and because it is all running on a single OS thread it is doing concurrency (not parallelism). However, multiple _logical processors_ can be used and then different goroutines can be run on different logical processors. In turn, these logical processors can be mapped to a different OS thread and if the system has multiple cores, these thread _may_ run in parallel.

**Interleaving** refers to the unknown order in which concurrent process instructions may be scheduled. This can make debugging difficult as it is harder to reason about the state of concurrent processes than it is about sequential processes.

## Race conditions

As interleaving is non-deterministic this can lead to a race condition - a situation where the outcome of a process will _depend_ on the interleaving. That is, the outcome of the process is _non-deterministic_ and can vary depending on the order in which the tasks are performed - effectively, a race.

The output of a process should almost definately be _deterministic_ so this is broken software.

A race condition can only occur when more than one task has access to a common variable - that is, the tasks _communicate_ via a common variable.

For example:

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    var x int

    go func() {
        time.Sleep(1 * time.Millisecond)
        x = 1
    }()

    go func() {
        time.Sleep(1 * time.Millisecond)
        x = 2
    }()

    time.Sleep(1 * time.Millisecond)
    fmt.Println(x)
}
```

## Working with goroutines

Standard func call will _block_ until the function returns:

```go
func main() {
  a := 1
  foo()
  a = 2
}
```

A goroutine will not block so main will not wait for `foo()`:

```go
func main() {
  a := 1
  go foo()
  a = 2
}
```

**Note:** when `main()` completes all goroutines exit - it does not wait for `foo()` to finish.

## Syncronisation

Synchronisation is used to prevent undesirable interleaving by creating _global_ events whose execution is monitored by all threads, simultaneously.

Synchronisation is necessary but it is, effectively, the opposite of concurrency and reduces performance.

The **`sync`** package contains functions to synchronise goroutines.

`sync.WaitGroup` forces a goroutine to wait for a specified number of _other_ goroutines before it executes.

- `Add()` increments the counter
- `Done()` decrements the counter
- `Wait()` blocks until counter = 0

Example:

```go
func main() {

    x := 1

    var wg sync.WaitGroup
    wg.Add(1)         // wait for one goroutine

    go func() {
        x = 2
        wg.Done()     // signal to the waitgroup
  }()
    wg.Wait()         // wait until all done

    fmt.Print(x)         // should be 2
}
```

<https://play.golang.org/p/wF-kByBoGDu>

## Communication

Goroutines usually work together to perform a larger task. This requires communication _between_ the goroutines.

Data is transferred between goroutines using _channels_.

Channels are _typed_ and created using `make()`, and data is sent and received over a channel using the `<-` operator:

```go
ch := make(chan int)
ch <- 2    // send 2 over chan
x := <- ch // receive next value from chan
```

Example:

```go
func main() {

    a, b, c, d := 1, 2, 3, 4

    ch := make(chan int)
    go sum(a, b, ch)
    go sum(c, d, ch)

    e, f := <-ch, <-ch
    fmt.Println(e * f) // 21
}

func sum(a, b int, ch chan int) {
    ch <- a + b
}
```

<https://play.golang.org/p/8R2hWY82qvR>

### Unbuffered channels

A channel is unbuffered by default and cannot hold data in transit.

This means that only a single value can be moved through the channel at a time and that both the send and receive operation must complete before another value can move through the channel. That is, bother the send or receive operation will block until the opposite side of the channel has performed its operation.

```go
// routine #1 and #2 are synchronous because of the unbuffered channel
ch := make(chan int)

// routine #1 - blocks until routine #2 receives the value
ch <- 1

// routine #2 - blocks until routine #1 sends a value
x := <- ch
```

Unbuffered channel communication is _synchronous_ and can be used in the same way as a `WaitGroup`:

```go
// routine #1
ch <- 1
// routine #2 - result disgarded but signals completion of thread #1
<- ch
```

### Buffered channels

Buffered channels can hold a specified number of values in transit.

They only block when the buffer is _full_ at the sending side or _empty_ at the receiving side.

```go
// routine #1
ch := make(chan int, 3)
ch <- 1 // not blocked
ch <- 2 // not blocked
ch <- 3 // will block until 1 value is received

// routine #2
a := <- ch // blocks until ch has one value
b := <- ch
c := <- ch // ch is empty
```

In general terms it is desirable to minimise blocking, for better performance.

### Channel iteration

Can iteratively read from a channel:

```go
for i := range ch {
    fmt.Println(i)
}
```

In this case the loop is exited when the sender explicitly closes the channel with `close(ch)`.

### Reading from multiple channels

In the case where data from all channels is required, this is straight forward:

```go
a := <- ch1
b := <- ch2
c = a * b
```

When data can be inbound on multiple channels, but _only one_ is required (ie. first come, first served) a `select` statement is used:

```go
select {
    case a = <- ch1:
        fmt.Println(a)
    case b <- ch2
        fmt.Println(b)
}
```

A `select` statement is used to read the _first_ data from a set of channels without blocking on the channels that did not deliver the data first, or at all.

`select` can also be used on outbound channels:

```go
select {
    case a = <- chIn:
        fmt.Println("Received", a)
    case chOut <- b:
        fmt.Print("Sent", b)
}
```

In both examples the entire `select` statement blocks until one of the cases _unblocks_ and execution can resume.

#### Using select with an _abort_ channel

Can use a signal from a different channel to quite waiting on data from one or more other channels. For example:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    ch := make(chan string)
    quit := make(chan int)
    var rec string
    var abort bool

    for {
        if abort {
            break
        }

        go userInput(ch, quit)

        select {
        case rec = <-ch:
            fmt.Println("Received on ch: ", rec)
        case <-quit:
            abort = true
        }
    }

    fmt.Println("Received a signal on the quit channel")
}

func userInput(ch chan string, quit chan int) {
    fmt.Print("Enter string, 'q' to quit: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    s := scanner.Text()
    if s == "q" {
        quit <- 1
        return
    }

    ch <- s
}
```

# Concurrency in Go

_Concurrency_ and _parallelism_ or _parallel execution_ are closely related, but not the same.

_Parallel execution_ is two programs running at exactly the same time - this requires two processors / cores.

Parallel execution does not speed up an individual task but gives better overall throughput.

The amount of benefit depends on the types of tasks being performed. Using a dishwashing metaphor: you cannot 
both wash and dry a dish at the same time.

Conversley, multi-core systems are only exploited when the programs running on them are designed for parallel execution.

Improvements in hardware speed are slowing due to physical limitations. 

So, _concurrent_ programming is used to exploit multi-core systems and achieve improved performance.

Concurrent code tells the computer that it is ok to run certain instructions in parallel, if it is able to do so.

**Concurrent vs Parallel**

Concurrent execution _*may*_ run in parallel (hardware permitting), but not necessarily. In concurrent execution the 
times that multiple tasks are actually running overlaps. However, it may be such that one task is _paused_ while another 
starts up. The processor time is divided up amongst the processes until they are all done. In effect, they are all _in progress_ 
concurrently but not necessarily _running_ concurrently.

By contrast, parallel execution means that multiple tasks are _actually running_ at the same time. There is no pause in one 
so another can have processor time. Hence, this requires seperate physical hardware.

The process of deciding which core particular tasks will run on, and if they will be parallel or concurrent, is not controlled 
by the programmer. It is controlled by the operating system and hardware and the Go runtime scheduler.

The programmer just decides what _can_ be done in parallel.

**Concurency without Parallelism**

Can still get improved performance with concurrent code on a single core. This is because the latency of other operations can 
be hidden. For example, memory operation, screen output, keyboard input, network activity and so on.

So, while waiting on other operations, the program can move on to the next task. In simple terms, it is just using waiting 
time constructively.




# Event-driven design

*Notes from ThreeDotes [Go event-driven course](https://academy.threedots.tech/trainings/go-event-driven)*



## Overview



Event-driven patterns are an *asynchronous* approach to building systems.

Traditional synchronous approaches involve processes that block flow while they wait for result:

```go
func TaskHandler() {
  subTaskOne()   // block, then 
  subTaskTwo()   // block, then
  subTaskThree() // block until done
}
```



The design of an event-driven system is such that tasks can be setup such that they run independently and do not hold each other up.

```go
func TaskHandler() {
  go subTaskOne()   
  go subTaskTwo()   
  go subTaskThree()
}
```



Event-driven systems are composed of processes that communicate asynchronously using *messages*. These are sent via a *message broker* which is an intermediary between message *publishers* and *subscribers*.

Published messages are appended to a *topic* and, generally, subscribers to the topic will receive the messages on a  *first-in, first-out (FIFO)* fashion.

In most systems, publishing a message will be a one-time process. Subscribers, however, will generally be an async worker process that starts and waits for new messages on a particular topic. 

Messages are delivered to a subscriber one-at-a-time so the subscriber lets the message broker know that a message has been correctly processed with a *message acknowledgement* (*ack*). 

If the subscriber fails to process the message it can sends a *negative acknowledgement (nack)* to the broker and the message is returned to the queue for delivery at a later stage. 

Often, there are multiple instances of a subscriber such as in distributed / containerised systems. To avoid messages being processed by multiple subscribers the concept of a *consumer group* is used (aka *subscription*, *queue*). Subscribers are allocated to a group and each message is delivered to a single subscriber within the group, in a round-robin fashion. In other words, messages are delivered to the group.



## Events

In event-driven architecture and *event* is a just a message, however it represents something that has already happened, ie an immutable fact. This is an important consideration in the way systems are designed. 

For example, a synchronous approach may couple all relevant processes together:

```go
func PlaceOrder(order Order) {
	SaveOrder(order)
	NotifyUser(order)
	NotifySales(order)
	NotifyWarehouse(order)
	GenerateInvoice(order)
	ChargeCustomer(order)
}
```

By publishing an event, the scope of the func is reduced:

```go
func PlaceOrder(order Order) {
	SaveOrder(order)          // <-- persist
	PublishOrderPlaced(order) // <-- publish event and let other processes respond
}
```

This is still a form of coupling, but allows for more flexibility in the way the processes are designed.



> [!NOTE]
>
> **An event should be a verb in past tense stating that something happened**. When designing the event, think about what happened, not what needs to happen after. Otherwise, you may fall into the "passive-aggressive events" trap, where the publisher knows what happens after the event is published.



## Marshalling

Event payloads are published as raw bytes so data structures must be serialised before they can be published. 

JSON is a common message format but can also use [Protocol Buffers](https://protobuf.dev) or [Avro](https://avro.apache.org). 

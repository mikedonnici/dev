# gRPC

<https://grpc.io>

## Overview

gRPC is an RPC framework that can be run in many different environments.

It can use protocol buffers as its _Interface Definition Languages_ (IDL) as well as its
message exchange format. It can use others (such as JSON) but protobuf is the default and most common.

Using gRPC, a client is able to call methods on a server application as though they were local methods. This is because the client has a _stub_ which provides a set of methods identical to those on the server.

When the client invokes a local method the server method is called with gRPC handling all of the underlying communication between server and client.

gRPC uses http2 which supports:

- multiplexing - simultaneous send / rec. over the same TCP channel, reduced latency
- server push - server can send multiple responses to a single request, less chatter
- header compression to reduce payload size
- transmission of binary data
- high level of security

http2 also enables four _modes_ of client-server communication:

1. **Unary**: standard single request, single response
2. **Server Streaming**: Single request from client, multiple response messages from server
3. **Client Streaming**: Multiple client request messages, single response from server
4. **Bi-directional Streaming**: Multiple, asyncronous client-server messages

In the protocol buffers these four types of contract are easily defined:

eg:

```proto
syntax = "proto3";

service GreetService {

  // unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};

  // server streaming
  rpc GreetOneMany(GreetOneManyRequest) returns (stream GreetOneManyResponse) {};

  // client streaming
  rpc GreetManyOne(stream GreetManyOneRequest) returns (GreetManyOneResponse) {};

  // bi-directional streaming
  rpc GreetManyMany(stream GreetManyManyRequest) returns (stream GreetManyManyResponse) {};
}
```
## Status and Error Codes

gRPC includes a set of [standard error and status codes](https://grpc.io/docs/guides/error/).

If protocol buffer is in use, there is an [extended error implementation](https://cloud.google.com/apis/design/errors#error_model) available.

A deadline can be set that creates a timeout, after which a `DEADLINE_EXCEEDED` error is returned.

Setting a deadline is generally recommended. The server should check if the deadline has been exceeded, and then cancel 
the work it is doing.

Note that deadlines are propagated form one gRPC server to another. So calls to multiple gRPC services should be 
considered when setting deadlines. 

Ref: <https://grpc.io/blog/deadlines/>

## Encryption and auth support

To generate tls certs see commands in `/ssl/openssl.sh`.

ref: 
- <https://grpc.io/docs/guides/auth/>
- <https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md>

## gRPC Reflection and Evans CLI

Can use the above to interrogate and interact with a gRPc service.

## gRPC Gateway

Use a a reverse-proxy client for implementing a RESTful API between clients and a gRPC server.

- <https://github.com/grpc-ecosystem/grpc-gateway>


Course ref: <https://github.com/simplesteph/grpc-go-course>



# backend types

This repo I'll implement various types of backend communication services.

To implement:

- [ ] WebSocket
- [ ] Server Sent Event
- [ ] Polling
- [ ] Long Polling

Brief explanation

## WebSocket

Uses the WebSocket protocol, which is a standardized protocol for real-time, bidirectional communication over the internet. Typically used to perform `push`. This strategy is best used when real-time is needed.

### WS Pros

- Bi-directional communication channel, which means both the client and server can send data to each other in real-time. It's also efficient in terms of resource usage because the connection is kept open and can be reused for multiple requests.

### WS Cons

- WebSockets require a dedicated server and a more complex implementation on both the client and server side.
- It also has limited support in some older web browsers.
- The server does not care about the client, so it may crashes the client side with overflow data amount or data which the client cannot handle.
- It also can send messages to the clients when they are not online anymore.

### WS Applicability

It's suitable for applications that require real-time updates and bidirectional communication between the client and server, such as online gaming and chat applications.

## Server-Sent Events (SSE)

Creates a connection between the client and keep streaming this connection. To summarize, it is a never ending stream of data. To help the client understand its data, it follows the following contract: each data needs to init with `data:`, while it ends with two blank-lines `\n\n`.

### SSE Pros

- Low latency because the server can send updates to the client without waiting for a request. It's also simple to implement on both the client and server side.
- It supports automatic reconnection in case of connection loss.

### SSE Cons

- SSE is a one-way communication channel, which means the client can't send data back to the server using SSE. In other words, it does not have bidirectional communication.
- As it locks the client connection, some browsers can lock the whole applications (HTTP/1.1 only allows 6 parallel connections).

### SSE Applicability

It's suitable for applications that require real-time updates, especially for applications that only need to send data from the server to the client.

## Polling

When the server receives a request, it instantly replies with an request Id, which the client can ask about its status. Therefore, the client needs to constantly ask to server if the request is done.

### Polling Pros

- Simple to implement, compatible with most web browsers and servers.

### Polling Cons

- Wastes resources (bandwidth and CPU) because the client has to constantly send requests to the server even when there are no updates.
- Can overload the server depending of the large number of request, while 90% or more will be false.
- May have a delay between the request is done and when the client asks again for it.

### Pooling Applicability

It's suitable for applications that don't require real-time updates, and the update frequency is low, such as weather prediction consultation.

## Long Polling

On the other hand, the long polling blocks the client connection when it asks if the request is done, then it replies only when the request is done.

### LP Pros

-Lower latency than regular polling because the server can hold the request until there's new data. It's also compatible with most web browsers and servers.

### LP Cons

- It can still waste resources if the server is holding a lot of requests for a long time. It also requires more complex implementation than regular polling.
- Can suffer from same problem as polling if the connection do not timeout and be hold for too long.

### LP Applicability

It's suitable for applications that require real-time updates, but the update frequency is low.

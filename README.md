# GoChat

GoChat is a simple group chat server built using the Go programming language. It uses TCP sockets to establish connections with clients and allows multiple clients to connect at the same time. The server broadcasts messages to all connected clients.

## Running the server

You can run the server using Docker or locally.

### Docker

```bash
docker compose up
```
To run the server in the background, use the `-d` or `--daemon` flag:

### Locally

```bash
go run main.go
```

## Connecting to the server

### Go client
    
```bash
go run client/client.go
```

### Python client

```bash
python3 client/client.py
```

### Arguments

```bash
-s, --server    Server address (default: localhost)
-p, --port      Server port (default: 6969)
```

## Documentation

## ChatServer

The `ChatServer` struct represents a chat server that can handle multiple clients. It has two fields:

- `clients`: a map of connected clients, where the keys are `net.Conn` objects and the values are empty structs.
- `mutex`: a read-write mutex used to synchronize access to the `clients` map to avoid multiple threads accessing it at the same time.

## Start

The `Start` function is used to start the chat server. It takes an `address` string as its only argument, which specifies the IP address and port number to listen on. It creates a new `ChatServer` instance and listens for incoming connections on the specified address. When a new connection is accepted, it spawns a new goroutine to handle the client.

## NewChatServer

The `NewChatServer` function creates a new `ChatServer` instance with an empty `clients` map.

## HandleClient

The `HandleClient` method is used to handle a single client connection. It takes a `net.Conn` object as its only argument, which represents the client connection. It adds the connection to the `clients` map, reads messages from the client, and broadcasts them to all connected clients. When the client disconnects, it removes the connection from the `clients` map.

## Broadcast

The `Broadcast` method is used to broadcast a message to all connected clients. It takes a `msg` string as its only argument, which represents the message to broadcast. It iterates over all connected clients and sends the message to each one. If an error occurs while sending the message, it removes the client from the `clients` map.

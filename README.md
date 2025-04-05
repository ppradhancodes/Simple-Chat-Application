# Simple Chat Application

This repository contains two implementations of a simple chat application - one in Rust and one in Go. Both implementations provide similar functionality with their respective language-specific features.

## Features

- User registration and management
- Send and receive messages between users
- View message history
- Search messages by keyword
- List registered users
- Concurrent message handling
- Timestamp-based message tracking

## Rust Implementation

The Rust implementation showcases:
- Memory safety through ownership and borrowing
- Concurrent message handling with async/await
- Structured error handling with Result types
- Thread-safe storage with Mutex

### Running the Rust Application

```bash
cd rust
cargo build
cargo run
```

## Go Implementation

The Go implementation demonstrates:
- Efficient message passing
- Simple error handling

### Running the Go Application

```bash
cd go
go mod tidy
go build
./chat-app
```

## Usage

Both applications provide a similar command-line interface:

1. Start the application
2. Enter a username to register/login
3. Use the following commands:
   - Send message (1)
   - View messages (2)
   - Search messages (3)
   - List users (4)
   - Delete message (5)
   - Logout (6)

## Design Considerations

- In-memory storage for simplicity
- Thread-safe operations for concurrent access
- Efficient message filtering and search
- User-friendly command-line interface
- Timestamp-based message tracking

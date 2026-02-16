# Go NATS Notification Service

A robust notification service built with Go and NATS messaging system. This service demonstrates a pub-sub pattern for handling various types of notifications including Email, SMS, Push notifications, and Alerts.

## Features

- 🚀 **Fast & Scalable**: Built on NATS for high-performance message delivery
- 📨 **Multiple Notification Types**: Support for Email, SMS, Push notifications, and Alerts
- 🎯 **Priority-based**: Notifications can be categorized by priority (Low, Medium, High, Critical)
- 🔄 **Pub/Sub Pattern**: Decoupled publisher and subscriber architecture
- ⚙️ **Configurable**: Easy configuration via environment variables
- 📦 **Metadata Support**: Attach custom metadata to notifications

## Architecture

The service consists of two main components:

1. **Publisher**: Sends notifications to NATS subjects
2. **Subscriber**: Listens to NATS subjects and processes notifications

```
Publisher → NATS Server → Subscriber(s)
```

## Prerequisites

- Go 1.21 or higher
- NATS Server (for local development)

## Installation

### Install NATS Server

**On macOS:**
```bash
brew install nats-server
```

**On Linux:**
```bash
curl -L https://github.com/nats-io/nats-server/releases/latest/download/nats-server-linux-amd64.tar.gz | tar xz
sudo mv nats-server /usr/local/bin/
```

**Using Docker:**
```bash
docker run -p 4222:4222 nats:latest
```

### Install Dependencies

```bash
go mod download
```

## Configuration

The service can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `NATS_URL` | NATS server URL | `nats://localhost:4222` |
| `NATS_SUBJECT` | NATS subject for notifications | `notifications` |

## Usage

### 1. Start NATS Server

```bash
nats-server
```

### 2. Run the Subscriber

In one terminal, start the subscriber to listen for notifications:

```bash
go run cmd/subscriber/main.go
```

### 3. Run the Publisher

In another terminal, publish notifications:

```bash
go run cmd/publisher/main.go
```

### Building Binaries

Build the publisher:
```bash
go build -o bin/publisher cmd/publisher/main.go
```

Build the subscriber:
```bash
go build -o bin/subscriber cmd/subscriber/main.go
```

Run the binaries:
```bash
./bin/subscriber &
./bin/publisher
```

## Project Structure

```
.
├── cmd/
│   ├── publisher/          # Publisher application
│   │   └── main.go
│   └── subscriber/         # Subscriber application
│       └── main.go
├── internal/
│   └── notification/       # Core notification logic
│       ├── types.go        # Notification types and models
│       ├── publisher.go    # Publisher implementation
│       └── subscriber.go   # Subscriber implementation
├── pkg/
│   └── config/            # Configuration management
│       └── config.go
├── go.mod
├── go.sum
└── README.md
```

## Notification Types

The service supports the following notification types:

- **Email**: Email notifications
- **SMS**: Text message notifications
- **Push**: Mobile/Web push notifications
- **Alert**: System alerts and warnings

## Priority Levels

Notifications can have different priority levels:

- **Low**: Non-urgent notifications
- **Medium**: Standard notifications
- **High**: Important notifications requiring attention
- **Critical**: Urgent notifications requiring immediate action

## Example Usage

### Publishing a Notification

```go
package main

import (
    "time"
    "github.com/kennedymwaniki/Go-NATS-PUBSUB-/internal/notification"
    "github.com/nats-io/nuid"
)

func main() {
    publisher, _ := notification.NewPublisher("nats://localhost:4222", "notifications")
    defer publisher.Close()

    notif := &notification.Notification{
        ID:        nuid.Next(),
        Type:      notification.Email,
        Priority:  notification.High,
        Recipient: "user@example.com",
        Subject:   "Welcome!",
        Body:      "Welcome to our service",
        Metadata: map[string]string{
            "template": "welcome",
        },
        CreatedAt: time.Now(),
    }

    publisher.Publish(notif)
}
```

### Subscribing to Notifications

```go
package main

import (
    "fmt"
    "github.com/kennedymwaniki/Go-NATS-PUBSUB-/internal/notification"
)

func main() {
    subscriber, _ := notification.NewSubscriber("nats://localhost:4222", "notifications")
    defer subscriber.Close()

    handler := func(notif *notification.Notification) error {
        fmt.Printf("Received: %s - %s\n", notif.Type, notif.Subject)
        return nil
    }

    subscriber.Subscribe(handler)
    
    // Keep running...
    select {}
}
```

## Advanced Features

### Custom NATS Configuration

```bash
export NATS_URL="nats://custom-host:4222"
export NATS_SUBJECT="custom.notifications"
go run cmd/subscriber/main.go
```

### Multiple Subscribers

You can run multiple subscriber instances for load balancing and high availability:

```bash
# Terminal 1
go run cmd/subscriber/main.go

# Terminal 2
go run cmd/subscriber/main.go

# Terminal 3
go run cmd/subscriber/main.go
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Linting

```bash
go vet ./...
```

## Production Considerations

1. **Error Handling**: Implement retry logic for failed notifications
2. **Monitoring**: Add metrics and monitoring for notification delivery
3. **Logging**: Use structured logging for better observability
4. **Security**: Use TLS for NATS connections in production
5. **Persistence**: Consider using NATS Streaming or JetStream for message persistence

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.

## Resources

- [NATS Documentation](https://docs.nats.io/)
- [NATS Go Client](https://github.com/nats-io/nats.go)
- [Go Documentation](https://golang.org/doc/)
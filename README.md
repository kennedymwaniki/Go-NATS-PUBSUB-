# Rental Management Notification Service

A scalable notification service for a rental management application built with Go and NATS messaging system. This service enables landlords to receive automatic notifications for payments, complaints, and repair requests from tenants.

## Features

- **Payment Notifications**: Automatic alerts when tenants make payments
- **Complaint Notifications**: Instant alerts when tenants lodge complaints
- **Repair Request Notifications**: Real-time notifications for maintenance and repair requests
- **Scalable Architecture**: Built on NATS for high-performance pub/sub messaging
- **Easy Integration**: Simple API for publishing and subscribing to notifications

## Architecture

The service uses a publisher-subscriber pattern with NATS as the message broker:

```
Publisher (Tenant App) --> NATS Server --> Subscriber (Landlord App)
```

## Prerequisites

- Go 1.20 or higher
- Docker and Docker Compose (for running NATS)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/kennedymwaniki/Go-NATS-PUBSUB-.git
cd Go-NATS-PUBSUB-
```

2. Install dependencies:
```bash
go mod download
```

3. Start NATS server using Docker Compose:
```bash
docker-compose up -d
```

## Usage

### Starting the Subscriber (Landlord Side)

The subscriber listens for all notifications and processes them:

```bash
go run cmd/subscriber/main.go
```

### Starting the Publisher (Tenant Side)

The publisher allows you to send different types of notifications:

```bash
go run cmd/publisher/main.go
```

Follow the interactive prompts to send:
1. Payment notifications
2. Complaint notifications
3. Repair request notifications

### Example: Publishing a Payment Notification

```bash
$ go run cmd/publisher/main.go
Notification Publisher Started
================================
Choose notification type:
1. Payment Notification
2. Complaint Notification
3. Repair Request Notification
4. Exit

Enter choice (1-4): 1
Enter Tenant ID: T001
Enter Property ID: P123
Enter Amount: 1500.00
Enter Payment Method (e.g., credit_card, bank_transfer): credit_card
Enter Status (e.g., completed, pending): completed
✓ Payment notification published successfully
```

## Notification Types

### 1. Payment Notification
Sent when a tenant makes a payment.

Fields:
- Tenant ID
- Property ID
- Amount
- Payment Method
- Status
- Timestamp

### 2. Complaint Notification
Sent when a tenant lodges a complaint.

Fields:
- Tenant ID
- Property ID
- Complaint Type
- Priority (low, medium, high)
- Message
- Timestamp

### 3. Repair Request Notification
Sent when a tenant requests repairs.

Fields:
- Tenant ID
- Property ID
- Request Type
- Urgency (low, medium, high)
- Message
- Timestamp

## Configuration

The service can be configured using environment variables:

- `NATS_URL`: NATS server URL (default: `nats://localhost:4222`)

Example:
```bash
export NATS_URL=nats://localhost:4222
```

## Project Structure

```
.
├── cmd/
│   ├── publisher/          # Publisher application
│   │   └── main.go
│   └── subscriber/         # Subscriber application
│       └── main.go
├── pkg/
│   ├── models/            # Data models
│   │   └── notification.go
│   └── notification/      # Notification service logic
│       └── service.go
├── config/                # Configuration
│   └── config.go
├── docker-compose.yml     # Docker compose for NATS
├── go.mod
└── README.md
```

## Development

### Building the Applications

Build the publisher:
```bash
go build -o bin/publisher cmd/publisher/main.go
```

Build the subscriber:
```bash
go build -o bin/subscriber cmd/subscriber/main.go
```

### Running Tests

```bash
go test ./...
```

## NATS Monitoring

NATS provides a monitoring interface accessible at `http://localhost:8222` when running via Docker Compose.

## Future Enhancements

- [ ] Email notification integration
- [ ] SMS notification support
- [ ] Push notification for mobile apps
- [ ] Notification history and analytics
- [ ] Web dashboard for viewing notifications
- [ ] Database persistence for notifications
- [ ] Authentication and authorization
- [ ] Rate limiting and throttling
- [ ] Notification templates
- [ ] Multi-tenant support

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.

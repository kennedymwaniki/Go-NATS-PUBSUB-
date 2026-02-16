package notification

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// Publisher handles publishing notifications to NATS
type Publisher struct {
	conn    *nats.Conn
	subject string
}

// NewPublisher creates a new notification publisher
func NewPublisher(natsURL, subject string) (*Publisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	log.Printf("Connected to NATS server at %s", natsURL)

	return &Publisher{
		conn:    nc,
		subject: subject,
	}, nil
}

// Publish sends a notification to NATS
func (p *Publisher) Publish(notification *Notification) error {
	data, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	if err := p.conn.Publish(p.subject, data); err != nil {
		return fmt.Errorf("failed to publish notification: %w", err)
	}

	log.Printf("Published notification: ID=%s, Type=%s, Priority=%s",
		notification.ID, notification.Type, notification.Priority)

	return nil
}

// Close closes the NATS connection
func (p *Publisher) Close() {
	if p.conn != nil {
		p.conn.Close()
		log.Println("Publisher connection closed")
	}
}

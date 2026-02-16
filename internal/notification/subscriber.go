package notification

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// MessageHandler is a function type that processes received notifications
type MessageHandler func(*Notification) error

// Subscriber handles subscribing to notifications from NATS
type Subscriber struct {
	conn         *nats.Conn
	subject      string
	subscription *nats.Subscription
}

// NewSubscriber creates a new notification subscriber
func NewSubscriber(natsURL, subject string) (*Subscriber, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	log.Printf("Subscriber connected to NATS server at %s", natsURL)

	return &Subscriber{
		conn:    nc,
		subject: subject,
	}, nil
}

// Subscribe starts listening for notifications
func (s *Subscriber) Subscribe(handler MessageHandler) error {
	sub, err := s.conn.Subscribe(s.subject, func(msg *nats.Msg) {
		var notification Notification
		if err := json.Unmarshal(msg.Data, &notification); err != nil {
			log.Printf("Error unmarshaling notification: %v", err)
			return
		}

		log.Printf("Received notification: ID=%s, Type=%s, Priority=%s",
			notification.ID, notification.Type, notification.Priority)

		if err := handler(&notification); err != nil {
			log.Printf("Error handling notification: %v", err)
		}
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	s.subscription = sub
	log.Printf("Subscribed to subject: %s", s.subject)

	return nil
}

// Close closes the subscription and NATS connection
func (s *Subscriber) Close() {
	if s.subscription != nil {
		s.subscription.Unsubscribe()
	}
	if s.conn != nil {
		s.conn.Close()
		log.Println("Subscriber connection closed")
	}
}

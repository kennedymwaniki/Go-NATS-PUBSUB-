package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/internal/notification"
	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/config"
	"github.com/nats-io/nuid"
)

func main() {
	log.Println("Starting Notification Publisher...")

	// Load configuration
	cfg := config.Load()
	log.Printf("Configuration: NATS URL=%s, Subject=%s", cfg.NATSUrl, cfg.Subject)

	// Create publisher
	publisher, err := notification.NewPublisher(cfg.NATSUrl, cfg.Subject)
	if err != nil {
		log.Fatalf("Failed to create publisher: %v", err)
	}
	defer publisher.Close()

	// Publish sample notifications
	notifications := []notification.Notification{
		{
			ID:        nuid.Next(),
			Type:      notification.Email,
			Priority:  notification.High,
			Recipient: "user@example.com",
			Subject:   "Welcome to our service!",
			Body:      "Thank you for signing up. We're excited to have you on board.",
			Metadata: map[string]string{
				"template": "welcome",
				"language": "en",
			},
			CreatedAt: time.Now(),
		},
		{
			ID:        nuid.Next(),
			Type:      notification.SMS,
			Priority:  notification.Critical,
			Recipient: "+1234567890",
			Subject:   "Security Alert",
			Body:      "Your account was accessed from a new device.",
			Metadata: map[string]string{
				"source": "security-system",
			},
			CreatedAt: time.Now(),
		},
		{
			ID:        nuid.Next(),
			Type:      notification.Push,
			Priority:  notification.Medium,
			Recipient: "device-token-123",
			Subject:   "New Message",
			Body:      "You have a new message from John Doe.",
			Metadata: map[string]string{
				"sender": "john.doe",
				"app":    "messaging",
			},
			CreatedAt: time.Now(),
		},
		{
			ID:        nuid.Next(),
			Type:      notification.Alert,
			Priority:  notification.Low,
			Recipient: "admin@example.com",
			Subject:   "System Update",
			Body:      "System maintenance scheduled for tomorrow at 2 AM.",
			Metadata: map[string]string{
				"category": "maintenance",
			},
			CreatedAt: time.Now(),
		},
	}

	// Publish notifications with delays
	for i, notif := range notifications {
		if err := publisher.Publish(&notif); err != nil {
			log.Printf("Error publishing notification %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Published notification %d/%d\n", i+1, len(notifications))
		time.Sleep(1 * time.Second)
	}

	log.Println("All notifications published successfully!")
}

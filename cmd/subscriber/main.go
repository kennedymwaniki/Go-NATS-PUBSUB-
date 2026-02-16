package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/internal/notification"
	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/config"
)

func main() {
	log.Println("Starting Notification Subscriber...")

	// Load configuration
	cfg := config.Load()
	log.Printf("Configuration: NATS URL=%s, Subject=%s", cfg.NATSUrl, cfg.Subject)

	// Create subscriber
	subscriber, err := notification.NewSubscriber(cfg.NATSUrl, cfg.Subject)
	if err != nil {
		log.Fatalf("Failed to create subscriber: %v", err)
	}
	defer subscriber.Close()

	// Define message handler
	handler := func(notif *notification.Notification) error {
		separator := strings.Repeat("=", 60)
		fmt.Println("\n" + separator)
		fmt.Printf("📬 NEW NOTIFICATION\n")
		fmt.Println(separator)
		fmt.Printf("ID:        %s\n", notif.ID)
		fmt.Printf("Type:      %s\n", notif.Type)
		fmt.Printf("Priority:  %s\n", notif.Priority)
		fmt.Printf("Recipient: %s\n", notif.Recipient)
		fmt.Printf("Subject:   %s\n", notif.Subject)
		fmt.Printf("Body:      %s\n", notif.Body)
		fmt.Printf("Created:   %s\n", notif.CreatedAt.Format("2006-01-02 15:04:05"))

		if len(notif.Metadata) > 0 {
			fmt.Println("Metadata:")
			for k, v := range notif.Metadata {
				fmt.Printf("  - %s: %s\n", k, v)
			}
		}
		fmt.Println(separator + "\n")

		// Simulate processing based on notification type
		switch notif.Type {
		case notification.Email:
			log.Printf("Processing email notification to %s", notif.Recipient)
		case notification.SMS:
			log.Printf("Processing SMS notification to %s", notif.Recipient)
		case notification.Push:
			log.Printf("Processing push notification to %s", notif.Recipient)
		case notification.Alert:
			log.Printf("Processing alert notification to %s", notif.Recipient)
		}

		return nil
	}

	// Subscribe to notifications
	if err := subscriber.Subscribe(handler); err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	log.Println("✓ Subscriber is running and listening for notifications...")
	log.Println("Press Ctrl+C to stop")

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Println("\nShutting down subscriber...")
}

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
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/notification"
	"github.com/nats-io/nats.go"
)

func main() {
	// Get NATS URL from environment or use default
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	// Create notification service
	service, err := notification.NewService(natsURL)
	if err != nil {
		log.Fatalf("Failed to create notification service: %v", err)
	}
	defer service.Close()

	log.Println("Notification Publisher Started")
	log.Println("================================")
	log.Println("Choose notification type:")
	log.Println("1. Payment Notification")
	log.Println("2. Complaint Notification")
	log.Println("3. Repair Request Notification")
	log.Println("4. Exit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter choice (1-4): ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			log.Println("Invalid choice. Please enter a number between 1 and 4.")
			continue
		}

		switch choice {
		case 1:
			publishPaymentNotification(service, reader)
		case 2:
			publishComplaintNotification(service, reader)
		case 3:
			publishRepairRequestNotification(service, reader)
		case 4:
			log.Println("Exiting...")
			return
		default:
			log.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func publishPaymentNotification(service *notification.Service, reader *bufio.Reader) {
	fmt.Print("Enter Tenant ID: ")
	tenantID, _ := reader.ReadString('\n')
	tenantID = strings.TrimSpace(tenantID)

	fmt.Print("Enter Property ID: ")
	propertyID, _ := reader.ReadString('\n')
	propertyID = strings.TrimSpace(propertyID)

	fmt.Print("Enter Amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Printf("Invalid amount: %v", err)
		return
	}

	fmt.Print("Enter Payment Method (e.g., credit_card, bank_transfer): ")
	paymentMethod, _ := reader.ReadString('\n')
	paymentMethod = strings.TrimSpace(paymentMethod)

	fmt.Print("Enter Status (e.g., completed, pending): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	if err := service.PublishPaymentNotification(tenantID, propertyID, amount, paymentMethod, status); err != nil {
		log.Printf("Failed to publish payment notification: %v", err)
		return
	}

	log.Println("✓ Payment notification published successfully")
}

func publishComplaintNotification(service *notification.Service, reader *bufio.Reader) {
	fmt.Print("Enter Tenant ID: ")
	tenantID, _ := reader.ReadString('\n')
	tenantID = strings.TrimSpace(tenantID)

	fmt.Print("Enter Property ID: ")
	propertyID, _ := reader.ReadString('\n')
	propertyID = strings.TrimSpace(propertyID)

	fmt.Print("Enter Complaint Type (e.g., noise, maintenance): ")
	complaintType, _ := reader.ReadString('\n')
	complaintType = strings.TrimSpace(complaintType)

	fmt.Print("Enter Priority (e.g., low, medium, high): ")
	priority, _ := reader.ReadString('\n')
	priority = strings.TrimSpace(priority)

	fmt.Print("Enter Message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	if err := service.PublishComplaintNotification(tenantID, propertyID, complaintType, priority, message); err != nil {
		log.Printf("Failed to publish complaint notification: %v", err)
		return
	}

	log.Println("✓ Complaint notification published successfully")
}

func publishRepairRequestNotification(service *notification.Service, reader *bufio.Reader) {
	fmt.Print("Enter Tenant ID: ")
	tenantID, _ := reader.ReadString('\n')
	tenantID = strings.TrimSpace(tenantID)

	fmt.Print("Enter Property ID: ")
	propertyID, _ := reader.ReadString('\n')
	propertyID = strings.TrimSpace(propertyID)

	fmt.Print("Enter Request Type (e.g., plumbing, electrical): ")
	requestType, _ := reader.ReadString('\n')
	requestType = strings.TrimSpace(requestType)

	fmt.Print("Enter Urgency (e.g., low, medium, high): ")
	urgency, _ := reader.ReadString('\n')
	urgency = strings.TrimSpace(urgency)

	fmt.Print("Enter Message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	if err := service.PublishRepairRequestNotification(tenantID, propertyID, requestType, urgency, message); err != nil {
		log.Printf("Failed to publish repair request notification: %v", err)
		return
	}

	log.Println("✓ Repair request notification published successfully")
}

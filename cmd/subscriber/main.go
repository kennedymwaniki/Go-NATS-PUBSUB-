package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/models"
	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/notification"
)

func main() {
	// Get NATS URL from environment or use default
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	// Create notification service
	service, err := notification.NewService(natsURL)
	if err != nil {
		log.Fatalf("Failed to create notification service: %v", err)
	}
	defer service.Close()

	log.Println("Notification Subscriber Started")
	log.Println("================================")
	log.Println("Listening for notifications...")

	// Subscribe to all notifications
	err = service.Subscribe(handleNotification)
	if err != nil {
		log.Fatalf("Failed to subscribe to notifications: %v", err)
	}

	// Wait for interrupt signal to gracefully shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("\nShutting down subscriber...")
}

func handleNotification(subject string, notification interface{}) {
	log.Println("\n" + strings.Repeat("=", 50))
	log.Printf("📬 New Notification Received on: %s", subject)
	log.Println(strings.Repeat("=", 50))

	switch n := notification.(type) {
	case *models.PaymentNotification:
		handlePaymentNotification(n)
	case *models.ComplaintNotification:
		handleComplaintNotification(n)
	case *models.RepairRequestNotification:
		handleRepairRequestNotification(n)
	default:
		log.Printf("Unknown notification type: %T", notification)
	}

	log.Println(strings.Repeat("=", 50))
}

func handlePaymentNotification(n *models.PaymentNotification) {
	log.Println("💰 PAYMENT NOTIFICATION")
	log.Printf("ID: %s", n.ID)
	log.Printf("Tenant ID: %s", n.TenantID)
	log.Printf("Property ID: %s", n.PropertyID)
	log.Printf("Amount: $%.2f", n.Amount)
	log.Printf("Payment Method: %s", n.PaymentMethod)
	log.Printf("Status: %s", n.Status)
	log.Printf("Message: %s", n.Message)
	log.Printf("Timestamp: %s", n.Timestamp.Format("2006-01-02 15:04:05"))

	// Here you can add logic to:
	// - Send email to landlord
	// - Update payment records in database
	// - Generate receipt
	// - Send SMS notification
	log.Println("\n✓ Payment notification processed")
}

func handleComplaintNotification(n *models.ComplaintNotification) {
	log.Println("⚠️  COMPLAINT NOTIFICATION")
	log.Printf("ID: %s", n.ID)
	log.Printf("Tenant ID: %s", n.TenantID)
	log.Printf("Property ID: %s", n.PropertyID)
	log.Printf("Complaint Type: %s", n.ComplaintType)
	log.Printf("Priority: %s", n.Priority)
	log.Printf("Status: %s", n.Status)
	log.Printf("Message: %s", n.Message)
	log.Printf("Timestamp: %s", n.Timestamp.Format("2006-01-02 15:04:05"))

	// Here you can add logic to:
	// - Alert landlord based on priority
	// - Create ticket in support system
	// - Assign to maintenance team
	// - Send acknowledgment to tenant
	log.Println("\n✓ Complaint notification processed")
}

func handleRepairRequestNotification(n *models.RepairRequestNotification) {
	log.Println("🔧 REPAIR REQUEST NOTIFICATION")
	log.Printf("ID: %s", n.ID)
	log.Printf("Tenant ID: %s", n.TenantID)
	log.Printf("Property ID: %s", n.PropertyID)
	log.Printf("Request Type: %s", n.RequestType)
	log.Printf("Urgency: %s", n.Urgency)
	log.Printf("Status: %s", n.Status)
	log.Printf("Message: %s", n.Message)
	log.Printf("Timestamp: %s", n.Timestamp.Format("2006-01-02 15:04:05"))

	// Here you can add logic to:
	// - Schedule repair based on urgency
	// - Assign to maintenance personnel
	// - Send cost estimate to tenant
	// - Update property maintenance records
	log.Println("\n✓ Repair request notification processed")
}

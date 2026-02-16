package main

import (
	"log"
	"time"

	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/notification"
)

func main() {
	// Connect to NATS
	service, err := notification.NewService("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer service.Close()

	log.Println("Sending test notifications...")
	
	// Test 1: Payment Notification
	log.Println("\n1. Sending Payment Notification...")
	err = service.PublishPaymentNotification("T001", "P123", 1500.00, "credit_card", "completed")
	if err != nil {
		log.Printf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Test 2: Complaint Notification
	log.Println("\n2. Sending Complaint Notification...")
	err = service.PublishComplaintNotification("T002", "P124", "noise", "high", "Loud music from neighbors")
	if err != nil {
		log.Printf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Test 3: Repair Request Notification
	log.Println("\n3. Sending Repair Request Notification...")
	err = service.PublishRepairRequestNotification("T003", "P125", "plumbing", "high", "Leaking faucet in kitchen")
	if err != nil {
		log.Printf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	log.Println("\n✓ All test notifications sent successfully!")
}

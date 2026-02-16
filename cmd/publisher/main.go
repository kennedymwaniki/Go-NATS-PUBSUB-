package main

import (
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

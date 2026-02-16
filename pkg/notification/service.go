package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kennedymwaniki/Go-NATS-PUBSUB-/pkg/models"
	"github.com/nats-io/nats.go"
)

// Service handles notification operations
type Service struct {
	nc *nats.Conn
}

// NewService creates a new notification service
func NewService(natsURL string) (*Service, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	log.Printf("Connected to NATS at %s", natsURL)
	return &Service{nc: nc}, nil
}

// Close closes the NATS connection
func (s *Service) Close() {
	if s.nc != nil {
		s.nc.Close()
	}
}

// PublishPaymentNotification publishes a payment notification
func (s *Service) PublishPaymentNotification(tenantID, propertyID string, amount float64, paymentMethod, status string) error {
	notification := &models.PaymentNotification{
		Notification: models.Notification{
			ID:         generateID(),
			Type:       models.PaymentType,
			TenantID:   tenantID,
			PropertyID: propertyID,
			Message:    fmt.Sprintf("Payment of $%.2f received via %s", amount, paymentMethod),
			Timestamp:  time.Now(),
		},
		Amount:        amount,
		PaymentMethod: paymentMethod,
		Status:        status,
	}

	return s.publish("notifications.payment", notification)
}

// PublishComplaintNotification publishes a complaint notification
func (s *Service) PublishComplaintNotification(tenantID, propertyID, complaintType, priority, message string) error {
	notification := &models.ComplaintNotification{
		Notification: models.Notification{
			ID:         generateID(),
			Type:       models.ComplaintType,
			TenantID:   tenantID,
			PropertyID: propertyID,
			Message:    message,
			Timestamp:  time.Now(),
		},
		ComplaintType: complaintType,
		Priority:      priority,
		Status:        "pending",
	}

	return s.publish("notifications.complaint", notification)
}

// PublishRepairRequestNotification publishes a repair request notification
func (s *Service) PublishRepairRequestNotification(tenantID, propertyID, requestType, urgency, message string) error {
	notification := &models.RepairRequestNotification{
		Notification: models.Notification{
			ID:         generateID(),
			Type:       models.RepairRequestType,
			TenantID:   tenantID,
			PropertyID: propertyID,
			Message:    message,
			Timestamp:  time.Now(),
		},
		RequestType: requestType,
		Urgency:     urgency,
		Status:      "pending",
	}

	return s.publish("notifications.repair", notification)
}

// publish publishes a notification to a NATS subject
func (s *Service) publish(subject string, notification interface{}) error {
	data, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	if err := s.nc.Publish(subject, data); err != nil {
		return fmt.Errorf("failed to publish notification: %w", err)
	}

	log.Printf("Published notification to subject: %s", subject)
	return nil
}

// Subscribe subscribes to all notification subjects
func (s *Service) Subscribe(handler func(subject string, notification interface{})) error {
	// Subscribe to payment notifications
	_, err := s.nc.Subscribe("notifications.payment", func(m *nats.Msg) {
		var notification models.PaymentNotification
		if err := json.Unmarshal(m.Data, &notification); err != nil {
			log.Printf("Error unmarshaling payment notification: %v", err)
			return
		}
		handler("notifications.payment", &notification)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to payment notifications: %w", err)
	}

	// Subscribe to complaint notifications
	_, err = s.nc.Subscribe("notifications.complaint", func(m *nats.Msg) {
		var notification models.ComplaintNotification
		if err := json.Unmarshal(m.Data, &notification); err != nil {
			log.Printf("Error unmarshaling complaint notification: %v", err)
			return
		}
		handler("notifications.complaint", &notification)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to complaint notifications: %w", err)
	}

	// Subscribe to repair request notifications
	_, err = s.nc.Subscribe("notifications.repair", func(m *nats.Msg) {
		var notification models.RepairRequestNotification
		if err := json.Unmarshal(m.Data, &notification); err != nil {
			log.Printf("Error unmarshaling repair notification: %v", err)
			return
		}
		handler("notifications.repair", &notification)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to repair notifications: %w", err)
	}

	log.Println("Subscribed to all notification subjects")
	return nil
}

// generateID generates a unique ID for notifications
func generateID() string {
	return uuid.New().String()
}

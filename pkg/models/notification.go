package models

import "time"

// NotificationType represents the type of notification
type NotificationType string

const (
	PaymentType       NotificationType = "payment"
	ComplaintType     NotificationType = "complaint"
	RepairRequestType NotificationType = "repair_request"
)

// Notification represents a general notification structure
type Notification struct {
	ID         string                 `json:"id"`
	Type       NotificationType       `json:"type"`
	TenantID   string                 `json:"tenant_id"`
	PropertyID string                 `json:"property_id"`
	Message    string                 `json:"message"`
	Timestamp  time.Time              `json:"timestamp"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

// PaymentNotification represents payment-related notifications
type PaymentNotification struct {
	Notification
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
}

// ComplaintNotification represents complaint-related notifications
type ComplaintNotification struct {
	Notification
	ComplaintType string `json:"complaint_type"`
	Priority      string `json:"priority"`
	Status        string `json:"status"`
}

// RepairRequestNotification represents repair request notifications
type RepairRequestNotification struct {
	Notification
	RequestType string `json:"request_type"`
	Urgency     string `json:"urgency"`
	Status      string `json:"status"`
}

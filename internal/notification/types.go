package notification

import "time"

// NotificationType represents the type of notification
type NotificationType string

const (
	Email NotificationType = "email"
	SMS   NotificationType = "sms"
	Push  NotificationType = "push"
	Alert NotificationType = "alert"
)

// Priority represents the notification priority
type Priority string

const (
	Low      Priority = "low"
	Medium   Priority = "medium"
	High     Priority = "high"
	Critical Priority = "critical"
)

// Notification represents a notification message
type Notification struct {
	ID        string            `json:"id"`
	Type      NotificationType  `json:"type"`
	Priority  Priority          `json:"priority"`
	Recipient string            `json:"recipient"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}

package types

import (
	"time"
)

// KafkaMessage represents the structure of messages sent through Kafka for email notifications.
type KafkaMessage struct {
	Email  EmailDetails `json:"email"`
	System System       `json:"system"`
}

// EmailDetails encapsulates the information required to send an email.
type EmailDetails struct {
	// Type indicates the type of the email, e.g., WELCOME, PASSWORD_RESET, etc.
	Type string `json:"type"`
	// To is the recipient's email address.
	To string `json:"to"`
	// Username is the recipient's username.
	Username string `json:"username"`
}

type System struct {
	// MessageID is a unique identifier for the message.
	MessageID string `json:"message_id"`
	// Timestamp denotes when the message was produced.
	Timestamp time.Time `json:"timestamp"`
	// Source indicates which service or system produced the message.
	Source string `json:"source"`
	// Metadata can contain any additional information relevant to the message context.
	Metadata map[string]string `json:"metadata"`
}

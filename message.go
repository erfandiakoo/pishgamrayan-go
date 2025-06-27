package pishgamrayan

import (
	"time"
)

// Message ...
type Message struct {
	SenderNumber     string   `json:"senderNumber"`
	MessageBodies    []string `json:"messageBodies"`
	RecipientNumbers []string `json:"recipientNumbers"`
	SendDateTime     string   `json:"sendDateTime"`
}

// MessageSendParam ...
type MessageSendParam struct {
	Date    time.Time
	LocalID []string
}

// MessageResult ...
type MessageResult struct {
	StatusCode     int64          `json:"statusCode"`
	MessageID      int64          `json:"messageId"`
	BlackListCount int64          `json:"blackListCount"`
	InvalidInputs  []InvalidInput `json:"invalidInputs"`
}

type InvalidInput struct {
	RecipientNumber string `json:"recipientNumber"`
	MessageBody     string `json:"messageBody"`
}

// MessageService ...
type MessageService struct {
	client *Client
}

// NewMessageService ...
func NewMessageService(client *Client) *MessageService {
	m := &MessageService{client: client}
	return m
}

type MessageStatusType int

const (
	Type_MessageStatus_Queued       MessageStatusType = 1
	Type_MessageStatus_Schulded     MessageStatusType = 2
	Type_MessageStatus_SentToCenter MessageStatusType = 4
	Type_MessageStatus_Sent         MessageStatusType = 5
	Type_MessageStatus_Failed       MessageStatusType = 6
	Type_MessageStatus_Delivered    MessageStatusType = 10
	Type_MessageStatus_Undelivered  MessageStatusType = 11
	Type_MessageStatus_Canceled     MessageStatusType = 13
	Type_MessageStatus_Filtered     MessageStatusType = 14
	Type_MessageStatus_Received     MessageStatusType = 50
	Type_MessageStatus_Incorrect    MessageStatusType = 100
)

var MessageStatusMap = map[MessageStatusType]string{
	Type_MessageStatus_Queued:       "Queued",
	Type_MessageStatus_Schulded:     "Schulded",
	Type_MessageStatus_SentToCenter: "SentToCenter",
	Type_MessageStatus_Sent:         "Sent",
	Type_MessageStatus_Failed:       "Failed",
	Type_MessageStatus_Delivered:    "Delivered",
	Type_MessageStatus_Undelivered:  "Undelivered",
	Type_MessageStatus_Canceled:     "Canceled",
	Type_MessageStatus_Filtered:     "Filtered",
	Type_MessageStatus_Received:     "Received",
	Type_MessageStatus_Incorrect:    "Incorrect",
}

func (t MessageStatusType) String() string {
	return MessageStatusMap[t]
}

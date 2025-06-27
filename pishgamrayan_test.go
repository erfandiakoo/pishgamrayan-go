package pishgamrayan

import (
	"testing"
)

var (
	api      *Pishgaman
	apiKey   string
	receptor string
)

func setup() {
	apiKey = "2yZLIdfYEJJYnI121MNh8aQcorKX4pkLO4ICxsCUoQs="
	receptor = "09128575183"
}

func TestMessageSend(t *testing.T) {
	setup()
	api := New(apiKey)
	sender := "500036637609"
	message := "Hello Go!"

	msg := Message{
		RecipientNumbers: []string{"09128575183"},
		SenderNumber:     sender,
		MessageBodies:    []string{message},
	}

	if res, err := api.Message.Send(msg); err != nil {
		t.Errorf("MessageSend failed: %v", err)
	} else if res.StatusCode >= 0 {
		t.Errorf("MessageSend failed: %+v", res)
	}
}

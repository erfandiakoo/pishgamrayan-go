package pishgamrayan

import (
	"net/url"
)

// Send ...
func (m *MessageService) Send(sender string, receptor []string, message string) ([]Message, error) {
	v := url.Values{}
	v.Set("sender", sender)
	v.Set("receptor", ToString(receptor))
	v.Set("message", message)

	return m.CreateSend(v)
}

// CreateSend ...
func (m *MessageService) CreateSend(v url.Values) ([]Message, error) {
	u := m.client.EndPoint("sms", "send")
	res := new(MessageResult)
	err := m.client.Execute(u.String(), v, res)
	return res.Entries, err
}

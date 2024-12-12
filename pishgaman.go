package pishgamrayan

// Kavenegar ...
type Pishgaman struct {
	Message *MessageService
}

// New ...
func New(apikey string) *Pishgaman {
	client := NewClient(apikey)
	return NewWithClient(client)
}

// NewWithClient ...
func NewWithClient(client *Client) *Pishgaman {
	p := &Pishgaman{}
	p.Message = NewMessageService(client)
	return p
}

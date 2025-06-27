package pishgamrayan

// Send ...
func (m *MessageService) Send(req Message) (MessageResult, error) {
	return m.CreateSend(req)
}

// CreateSend ...
func (m *MessageService) CreateSend(req Message) (MessageResult, error) {
	u := m.client.EndPoint("Messages", "Send")
	res := new(MessageResult)
	err := m.client.Execute(u.String(), req, res)
	return *res, err
}

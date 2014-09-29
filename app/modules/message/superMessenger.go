package message

type SuperMessenger struct {
	message string
}

func (m *SuperMessenger) GetMessage() string {
	return "super " + m.message
}

func (m *SuperMessenger) SetMessage(message string) {
	m.message = message
}

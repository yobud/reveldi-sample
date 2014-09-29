package message

type Messenger struct {
	message string
}

func (m *Messenger) GetMessage() string {
	return m.message
}

func (m *Messenger) SetMessage(message string) {
	m.message = message
}

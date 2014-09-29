package message

type PrefixMessenger struct {
	prefix string
	message string
}

func (m *PrefixMessenger) GetMessage() string {
	return m.prefix + " " + m.message
}

func (m *PrefixMessenger) SetMessage(message string) {
	m.message = message
}

func (m *PrefixMessenger) SetPrefix(prefix string) {
	m.prefix = prefix
}

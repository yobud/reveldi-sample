package message

type MessengerInterface interface {
	GetMessage() string
	SetMessage(message string)
}

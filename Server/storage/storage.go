package storage

import (
	"chatex/message"
)

type Storage interface {
	AddMessage(message *message.Message) error

	GetAllMessages() ([]*message.Message, error)
	GetMessage(index uint) (*message.Message, error)
}

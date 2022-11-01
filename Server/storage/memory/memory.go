package memory

import (
	"chatex/message"
	"sync"
)

type memory struct {
	messages []*message.Message
	lock     sync.Mutex
}

func (storage *memory) AddMessage(message *message.Message) error {
	storage.lock.Lock()
	storage.messages = append(storage.messages, message)
	storage.lock.Unlock()

	return nil
}

func (storage *memory) GetAllMessages() ([]*message.Message, error) {
	return storage.messages, nil
}

func (storage *memory) GetMessage(index uint) (*message.Message, error) {
	return storage.messages[index], nil
}

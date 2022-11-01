package memory

import (
	"chatex/message"
	"chatex/storage"
)

var instance *memory

func Instance() storage.Storage {
	if instance == nil {
		instance = &memory{
			messages: make([]*message.Message, 0),
		}
	}

	return instance
}

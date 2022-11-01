package server

import "chatex/message"

type Handler interface {
	HandleIncomingData()
	HandleOutgoingData()

	TransmitMessage(message message.Message)
}

func (server *Server) AddClient(client Handler) {
	server.clients = append(server.clients, client)
}

func (server *Server) RemoveClient(client Handler) {
	for i, s := range server.clients {
		if s == client {
			server.clients = append(server.clients[:i], server.clients[i+1:]...)
			break
		}
	}
}

func (server *Server) TransmitMessage(newMessage message.Message) {
	for _, handler := range server.clients {
		handler.TransmitMessage(newMessage)
	}
}

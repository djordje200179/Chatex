package client

import (
	"chatex/message"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Client struct {
	nickname      string
	connection    net.Conn
	messagesQueue chan message.Message
}

func New(connection net.Conn) *Client {
	return &Client{
		"anonymous",
		connection,
		make(chan message.Message),
	}
}

func (client *Client) Addr() net.Addr {
	return client.connection.RemoteAddr()
}

func (client *Client) HandleIncomingData() {
	fmt.Println(client.Addr(), "started receiving messages")

	decoder := json.NewDecoder(client.connection)

	for {
		fmt.Println(client.Addr(), "waiting for a message")

		var command newMessageCommand
		err := decoder.Decode(&command)
		if err != nil {
			fmt.Println(client.Addr(), "error decoding the message")
			panic(err)
		}

		newMessage := message.Message{
			Text:      command.message,
			Sender:    client.Addr(),
			Timestamp: time.Now(),
		}

		fmt.Println(newMessage)
		//client.server.CreateMessage(newMessage)
	}
}

func (client *Client) HandleOutgoingData() {
	fmt.Println(client.Addr(), "started transmitting messages")

	encoder := json.NewEncoder(client.connection)

	for newMessage := range client.messagesQueue {
		err := encoder.Encode(newMessage)
		if err != nil {
			fmt.Println(client.Addr(), "error encoding the message")
			panic(err)
		}
	}
}

func (client *Client) TransmitMessage(message message.Message) {
	if message.Sender == client.Addr() {
		return
	}

	client.messagesQueue <- message
}

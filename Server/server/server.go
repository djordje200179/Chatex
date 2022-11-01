package server

import (
	"chatex/server/client"
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	clients  []Handler
}

func New(port uint16) *Server {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	return &Server{
		listener,
		[]Handler{},
	}
}

func (server *Server) Start() {
	fmt.Println("Listening at", server.listener.Addr())

	for {
		fmt.Println("Waiting for a connection...")

		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting the connection:", err)
			continue
		}

		fmt.Println("Accepted a connection from", conn.RemoteAddr())

		newClient := client.New(conn)
		server.AddClient(newClient)

		go newClient.HandleIncomingData()
		go newClient.HandleOutgoingData()
	}
}

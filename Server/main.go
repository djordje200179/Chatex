package main

import "chatex/server"

func main() {
	server.New(server.Config{
		Port: 555,
	}).Start()
}

package main

import (
	"terminal_chatroom/server"
)

func main() {
	server := server.New(&server.Config{
		Host: "localhost",
		Port: "3333",
	})
	server.Run()
}

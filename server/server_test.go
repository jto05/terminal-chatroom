package server_test

import (
	"testing"

	"terminal_chatroom/server"
)

func TestServer(t *testing.T) {
	server := server.New(&server.Config{
		Host: "localhost",
		Port: "3333",
	})
	server.Run()
}

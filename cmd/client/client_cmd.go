package main

import (
	"os"

	"terminal_chatroom/client"
)

func main() {
	c := client.New(&client.Config{
		Host:   "localhost",
		Port:   "3333",
		Input:  os.Stdin,
		Output: os.Stdout,
	})

	c.Run()
}

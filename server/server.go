package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	host    string
	port    string
	clients []Client
}

type Client struct {
	conn net.Conn
}

type Config struct {
	Host string
	Port string
}

func New(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

// TODO: add a server function that allows the server to broadcast a message to all clients
/*
Start contains
*/
func (s *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on %s:%s...\n", s.host, s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
		}

		go client.handleConnections()
	}
}

func (s *Server) broadcast(msg string) {
	fmt.Printf("Broadcasting: %s", msg)
}

func (client *Client) handleConnections() {
	reader := bufio.NewReader(client.conn)
	for { // loops until connection on Client's until error received or connection closed
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			client.conn.Close()
			return
		}

		fmt.Printf("Message incoming: %s", string(message))
		to_write := "Message received. \n"
		_, err = client.conn.Write([]byte(to_write))
		if err != nil {
			fmt.Println(err)
		}
	}
}

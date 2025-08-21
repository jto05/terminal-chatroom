package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	address string
}

func New(address string) *Server {
	return &Server{
		address,
	}
}

/*
Start contains
*/
func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on " + s.address + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.handleConnections(conn)

	}
}

func (s *Server) handleConnections(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Receiving: %f")
}

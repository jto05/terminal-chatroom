package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	host string
	port string

	username string
	input    io.Reader
	output   io.Writer
}

type Config struct {
	Host   string
	Port   string
	Input  io.Reader
	Output io.Writer
}

func New(config *Config) (client *Client) {
	return &Client{
		host:     config.Host,
		port:     config.Port,
		username: "",
		input:    config.Input,
		output:   config.Output,
	}
}

func (c *Client) Run() {
	conn, err := net.Dial("tcp", c.host+":"+c.port)
	if err != nil {
		fmt.Fprintln(c.output, err)
		log.Fatal(err)
	}
	defer conn.Close()

	sc := bufio.NewScanner(c.input)

	// change username if wasn't initialized in config
	if c.username == "" {
		fmt.Fprintf(c.output, "Enter username: ")
		if sc.Scan() {
			c.username = sc.Text()
		}
	}

	go c.receiveServerData(conn) // start getting data from server

	//
	for sc.Scan() {
		fmt.Fprintf(c.output, "%s : ", c.username)

		text := sc.Text()

		if text == "/exit" {
			fmt.Fprintln(c.output, "goodbye!")
			break
		}

		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *Client) receiveServerData(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(c.output, "Disconnected from server.")
		}

		// fmt.Fprintf(c.output, "From server: %s", message)
	}
}

package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	host   string
	port   string
	input  io.Reader
	output io.Writer
}

type Config struct {
	Host   string
	Port   string
	Input  io.Reader
	Output io.Writer
}

func New(config *Config) (client *Client) {
	return &Client{
		host:   config.Host,
		port:   config.Port,
		input:  config.Input,
		output: config.Output,
	}
}

// TODO: allow client to send as many requests to the server as they want
func (c *Client) Run() {
	sc := bufio.NewScanner(c.input)

	conn, err := net.Dial("tcp", c.host+":"+c.port)
	if err != nil {
		fmt.Fprintln(c.output, err)
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintln(c.output, "Enter text:")

	for sc.Scan() {
		text := sc.Text()
		fmt.Fprintf(c.output, "You wrote: '%s'\n", text)

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

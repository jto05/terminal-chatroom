package client

import (
	"bufio"
	"fmt"
	"io"
	// "log"
	// "net"
)

type Client struct {
	host   string
	port   string
	reader io.Reader
	writer io.Writer
}

type Config struct {
	Host   string
	Port   string
	Reader io.Reader
	Writer io.Writer
}

func New(config *Config) (client *Client) {
	return &Client{
		host:   config.Host,
		port:   config.Port,
		reader: config.Reader,
		writer: config.Writer,
	}
}

// TODO: allow client to send as many requests to the server as they want
func (c *Client) Run() {
	sc := bufio.NewScanner(c.reader)
	// conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.host, c.port))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// a scanner?
	fmt.Fprintln(c.writer, "Enter text:")
	for sc.Scan() {
		line := sc.Text()
		fmt.Fprintln(c.writer, "You wrote: "+line)

		if line == "exit" {
			fmt.Fprintln(c.writer, "goodbye!")
			break
		}
		// conn.Write([]byte(input))
	}
}

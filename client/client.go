package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Client struct {
	host string
	port string

	username          string
	last_printed_line string
	input             io.Reader
	output            io.Writer
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

	fmt.Fprintf(c.output, "Hello! You're dialing %s:%s\n", c.host, c.port)
	fmt.Fprintln(c.output, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// change username if wasn't initialized in config
	if c.username == "" {
		fmt.Fprintln(c.output, "Enter username:")
		if sc.Scan() {
			c.username = strings.TrimSpace(sc.Text())
		}
		fmt.Fprintln(c.output, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}

	go c.receiveServerData(conn) // start getting data from server

	for {
		fmt.Fprintf(c.output, "%s : ", c.username)

		if !sc.Scan() {
			break
		}

		line := strings.TrimSpace(sc.Text())

		if line == "/exit" {
			fmt.Fprintln(c.output, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Fprintln(c.output, "Goodbye!")
			break
		}

		msg := c.username + " : " + line

		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *Client) receiveServerData(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(c.output, "Disconnected from server.")
		}
		msg = strings.TrimSpace(msg)

		split_msg := strings.Split(msg, " : ") // split message into username and content

		if msg == "" {
			continue
		}

		if split_msg[0] != c.username { // only print if the message written was NOT made from user
			fmt.Fprint(c.output, "\r\033[K")           // clear line
			fmt.Fprintln(c.output, msg)                // print the incoming message
			fmt.Fprintf(c.output, "%s : ", c.username) // print the prompt
		}
	}
}

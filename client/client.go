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

	fmt.Fprintln(c.output, "Hello!")

	// change username if wasn't initialized in config

	go c.receiveServerData(conn) // start getting data from server

	for sc.Scan() {
		if c.username == "" {
			fmt.Fprintf(c.output, "Enter username: ")
			c.username = sc.Text()
		}
		if c.last_printed_line != (c.username + " : ") {
			fmt.Fprintf(c.output, "%s : ", c.username)
			c.last_printed_line = c.username + " : "
		}
		line := sc.Text()

		line = strings.TrimSpace(line)

		if line == "/exit" {
			fmt.Fprintln(c.output, "goodbye!")
			break
		}

		msg := c.username + " : " + line
		c.last_printed_line = msg

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

		split_msg := strings.Split(msg, " : ")

		if split_msg[0] != c.username { // only print if the message written was NOT made from user
			// i need to make it so this message is overrites if prev line was "username : "
			fmt.Fprint(c.output, "\r"+msg+"\n")
			fmt.Fprint(c.output, c.username+" : ")
			c.last_printed_line = c.username + " : "
		}
	}
}

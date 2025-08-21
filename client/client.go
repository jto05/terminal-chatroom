package client

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	host string
	port string
}

type Config struct {
	Host string
	Port string
}

func New(config *Config) (client *Client) {
	return &Client{
		host: config.Host,
		port: config.Port,
	}
}

func (client *Client) Run() {
}

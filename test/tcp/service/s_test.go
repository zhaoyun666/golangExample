package service

import (
	"log"
	"net"
	"testing"
	"time"
)

func buildTestServer() *server {
	return New("127.0.0.1:8989")
}

func Test_accept_new_client_callback(t *testing.T) {
	server := buildTestServer()

	var messageReceived bool
	var messageText string
	var newClient bool
	var connectionClosed bool

	server.OnNewClient(func(c *Client) {
		newClient = true
	})

	server.OnNewMessage(func(c *Client, message string) {
		messageReceived = true
		messageText = message
	})

	server.OnClientConnectionClosed(func(c *Client, err error) {
		connectionClosed = true
	})
	go server.Listen()

	// Wait for server
	// If test fails - increase this value
	time.Sleep(10 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	if err != nil {
		log.Fatalf("error:%s", err.Error())
	}
	conn.Write([]byte("Test send one message\n"))
	conn.Close()
}

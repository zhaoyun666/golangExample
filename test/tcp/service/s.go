package service

import (
	"bufio"
	"log"
	"net"
	"time"
)

type Client struct {
	conn   net.Conn
	Server *server
}

// TCP server
type server struct {
	address                  string
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message string)
}

// Read client data from channel
func (c *Client) listen() {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		c.Server.onNewMessage(c, message)
	}
}

// send text message to client
func (c *Client) Send(message string) error {
	_, err := c.conn.Write([]byte(message))
	return err
}

// send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, message string)) {
	s.onNewMessage = callback
}

// Start Network Server
func (s *server) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("err:%s", err.Error())
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
		go recreateConnection(conn)
		s.onNewClientCallback(client)
	}
}

// Recreate connection tcp
func recreateConnection(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	request := make([]byte, 1024)
	defer conn.Close()
	for {
		recv_len, err := conn.Read(request)
		if err != nil {
			log.Printf("err:%s", err.Error())
			break
		}
		if recv_len == 0 {
			break
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime + "\n"))
		request = make([]byte, 1024)
	}
}

// Create new tcp server instance
func New(address string) *server {
	log.Println("creating server with address", address)
	server := &server{
		address: address,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message string) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})
	return server
}

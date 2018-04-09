package main

import (
	"learning-golang-process/test/tcp/service"
	"fmt"
)

func main() {
	server := service.New("127.0.0.1:8989")
	server.OnNewClient(func(c *service.Client) {
		c.Send("hello")
	})
	server.OnNewMessage(func(c *service.Client, message string) {
		fmt.Println(message)
	})

	server.OnClientConnectionClosed(func(c *service.Client, err error) {
		fmt.Println(c, err)
	})

	server.Listen()
}

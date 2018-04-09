package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"os"
	"io"
	"learning-golang-process/test/tcp/common"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	if err != nil {
		log.Fatalf("err:%s", err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("err:%s", err.Error())
	}
	defer conn.Close()
	go printOutput(conn)
	writeInput(conn)
}

func writeInput(conn *net.TCPConn) {
	fmt.Print("Enter username: ")
	// Read from stdin.
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = username[:len(username)-1]
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter text: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		err = common.WriteMsg(conn, username+": "+text)
		if err != nil {
			log.Println(err)
		}
	}
}

func printOutput(conn *net.TCPConn) {
	for {

		msg, err := common.ReadMsg(conn)
		// Receiving EOF means that the connection has been closed
		if err == io.EOF {
			// Close conn and exit
			conn.Close()
			fmt.Println("Connection Closed. Bye bye.")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}
}

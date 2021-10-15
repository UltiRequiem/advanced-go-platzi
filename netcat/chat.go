package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incommingClients = make(chan Client)
	leavingClients   = make(chan Client)
	messages         = make(chan string)

	host = flag.String("host", "localhost", "The Host")
	port = flag.Int("port", 3000, "The Port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	message := make(chan string)

	go MessageWriter(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server %s\n", clientName)

	messages <- fmt.Sprintf("New client is here: %s\n", clientName)

	incommingClients <- message

	inputMessage := bufio.NewScanner(conn)

	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message

	messages <- fmt.Sprintf("%s said goodbye!\n", clientName)

}

func MessageWriter(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)

	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		}
	}
}

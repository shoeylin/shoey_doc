package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7"
	CONN_TYPE = "tcp"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:7")
	go send(conn)
	wg.Add(10)
	wg.Wait()
}
func send(conn net.Conn) {
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		wg.Done()
	}
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting the server: ", err)
		os.Exit(1)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connections")
			continue
		}

		// a thread to handle each connection
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}

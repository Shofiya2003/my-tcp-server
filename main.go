package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
	for {
		// Accept() is a blocking call
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
	fmt.Println("New Client Connected", conn.RemoteAddr().String())
	// buffer := make([]byte, 1024)

	// n, err := conn.Read(buffer)
	// if err != nil {
	// 	log.Println("Error reading from the ", err)
	// }
	// log.Println(buffer[:n])
	time.Sleep(2 * time.Second)
	_, err := conn.Write([]byte("Welcome!\n"))
	if err != nil {
		log.Println("Error writing to the client", err)
	}

}

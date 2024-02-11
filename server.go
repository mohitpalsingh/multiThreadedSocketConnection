package main

import (
	"fmt"
	"net"
)

func main() {
	runServerMultiThreaded()
}

func runServerMultiThreaded() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Recieved:", string(buffer))

	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
}

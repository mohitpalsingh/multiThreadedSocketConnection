package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	runClientSingleThreaded()
}

func runClientSingleThreaded() {
	var wg sync.WaitGroup

	numberOfClients := 5

	wg.Add(numberOfClients)

	for i := 0; i < numberOfClients; i++ {
		go func(clientId int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println("Error connecting:", err.Error())
				return
			}

			defer conn.Close()

			message := "Hello Server!"
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending:", err.Error())
				return
			}

			buffer := make([]byte, 1024)
			_, err = conn.Read(buffer)
			if err != nil {
				fmt.Println("Error receiving message:", err.Error())
				return
			}
			fmt.Println("Server says:", string(buffer))
		}(i + 1)
	}
	wg.Wait()

}

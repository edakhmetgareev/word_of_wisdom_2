package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aed86/proof_of_work/internal/server"
)

func main() {
	// Start the tcp server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Server listening on :8080")
	for {
		fmt.Println("Ready to accept connections...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Accepted connection from:", conn.RemoteAddr())

		// Handle the connection
		go func() {
			fmt.Println("Handling connection from:", conn.RemoteAddr())
			err := server.Handle(conn)
			if err != nil {
				fmt.Println("Error handling connection:", err)
			}
		}()
	}
}

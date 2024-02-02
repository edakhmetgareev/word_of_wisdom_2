package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aed86/proof_of_work/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	// Start the tcp server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		fmt.Println("Error listening:", err)
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Printf("Server listening on port: %s \n", serverPort)
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
			conn.Close()
		}()
	}
}

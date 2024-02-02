package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aed86/word_of_wisdom_2/internal/server"
	"github.com/aed86/word_of_wisdom_2/pkg"
)

func main() {
	err := pkg.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	serverPort, err := pkg.GetEnv("SERVER_PORT")
	if err != nil {
		log.Fatalf("Error getting SERVER_PORT: %s", err)
	}

	err = startServer(serverPort)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func startServer(port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("error listening: %w", err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on port: %s \n", port)
	for {
		fmt.Println("Ready to accept connections...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Accepted connection from:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection from:", conn.RemoteAddr())
	err := server.Handle(conn)
	if err != nil {
		fmt.Println("Error handling connection:", err)
	}
	conn.Close()
}

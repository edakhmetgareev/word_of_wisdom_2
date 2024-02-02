package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aed86/proof_of_work/internal/client"
)

const (
	serverAddr = "localhost:8080"
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	if err := client.Handle(conn); err != nil {
		log.Fatal(err)
	}
}

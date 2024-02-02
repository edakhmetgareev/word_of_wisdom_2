package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/aed86/proof_of_work/internal/client"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	fmt.Println("Connecting to server:", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		fmt.Println("Error setting deadline:", err)
		return
	}

	if err := client.Handle(conn); err != nil {
		log.Fatal(err)
	}
}

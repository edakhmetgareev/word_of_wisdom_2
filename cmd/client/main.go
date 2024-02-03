package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aed86/word_of_wisdom_2/internal/client"
	"github.com/aed86/word_of_wisdom_2/pkg"
	"github.com/aed86/word_of_wisdom_2/pkg/tcpconn"
)

const clientsCount = 500

func main() {
	err := pkg.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var (
		wg           sync.WaitGroup
		successCount atomic.Int32
	)
	wg.Add(clientsCount)
	for i := 0; i < clientsCount; i++ {
		go run(&wg, &successCount)
	}

	wg.Wait()

	fmt.Printf("Success count: %d\n", successCount.Load())
}

func run(wg *sync.WaitGroup, successCount *atomic.Int32) {
	defer wg.Done()
	conn, err := getConn()
	if err != nil {
		log.Fatalf("Error getting connection: %s", err)
	}
	defer conn.Close()

	c := client.NewClient(tcpconn.NewTCPConn(conn))
	if err := c.Handle(); err != nil {
		fmt.Println("Error handling client:", err)
		return
	}
	successCount.Add(1)
}

func getConn() (net.Conn, error) {
	address, err := getServerAddr()
	if err != nil {
		return nil, fmt.Errorf("error getting server address: %w", err)
	}

	fmt.Println("Connecting to server:", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %w", err)
	}

	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		fmt.Println("Error setting deadline:", err)
		return nil, fmt.Errorf("error setting deadline: %w", err)
	}

	return conn, nil
}

func getServerAddr() (string, error) {
	serverHost, err := pkg.GetEnv("SERVER_HOST")
	if err != nil {
		return "", fmt.Errorf("error getting SERVER_HOST: %w", err)
	}

	serverPort, err := pkg.GetEnv("SERVER_PORT")
	if err != nil {
		return "", fmt.Errorf("error getting SERVER_PORT: %w", err)
	}

	return fmt.Sprintf("%s:%s", serverHost, serverPort), nil
}

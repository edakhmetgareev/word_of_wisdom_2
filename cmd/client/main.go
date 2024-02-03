package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/aed86/word_of_wisdom_2/internal/client"
	"github.com/aed86/word_of_wisdom_2/pkg"
	"github.com/aed86/word_of_wisdom_2/pkg/tcpconn"
)

type ClientConfig struct {
	Address string
}

func main() {
	err := pkg.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var (
		wg           sync.WaitGroup
		successCount atomic.Int32
	)

	address, err := getServerAddr()
	if err != nil {
		log.Fatalf("Error getting server address: %s", err)
	}

	c := client.NewClient()
	maxClientsCount := getMaxFileDescriptorsLimit() / 100 // keep some space for other processes
	log.Printf("Max clients count: %d\n", maxClientsCount)

	sem := make(chan struct{}, 130)

	config := &ClientConfig{Address: address}
	for i := 0; i < maxClientsCount; i++ {
		sem <- struct{}{} // acquire the semaphore

		wg.Add(1)
		go func(i int) {
			run(c, config, &wg, &successCount, i)
			<-sem // release the semaphore
		}(i)
	}

	wg.Wait()

	log.Printf("Success count: %d\n", successCount.Load())
}

func run(c *client.Client, config *ClientConfig, wg *sync.WaitGroup, successCount *atomic.Int32, i int) {
	defer wg.Done()

	conn, err := net.Dial("tcp", config.Address)
	if err != nil {
		log.Printf("error connecting to server: %s, i: %d \n", err, i)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	if err := c.Handle(tcpconn.NewTCPConn(conn), i); err != nil {
		log.Printf("Error handling client: err: %s, i: %d \n", err, i)
		return
	}

	successCount.Add(1)
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

func getMaxFileDescriptorsLimit() int {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	return int(rLimit.Cur) // soft limit
}

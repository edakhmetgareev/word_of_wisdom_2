package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/aed86/word_of_wisdom_2/internal/server"
	"github.com/aed86/word_of_wisdom_2/pkg"
	"github.com/aed86/word_of_wisdom_2/pkg/tcpconn"
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

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	go startTCPServer(ctx, serverPort, &wg)
	waitForShutdownSignal(cancel)

	wg.Wait()
	fmt.Println("Server shutdown complete.")
}

func startTCPServer(ctx context.Context, port string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println("error listening: %w", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on port: %s \n", port)
	for {
		select {
		case <-ctx.Done():
			// Stop server if context is cancelled
			return
		default:
		}

		// Receive incoming connections in a non-blocking way
		connChan := make(chan net.Conn)
		errChan := make(chan error)

		go func() {
			fmt.Println("Ready to accept connections...")
			conn, err := listener.Accept()
			if err != nil {
				errChan <- err
				return
			}
			connChan <- conn
		}()

		select {
		case <-ctx.Done():
			fmt.Println("Shutting down server...")
			return
		case conn := <-connChan:
			fmt.Println("Accepted connection from:", conn.RemoteAddr())
			go handleConnection(conn)
		case err := <-errChan:
			fmt.Println("Error accepting connection:", err)
		}
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection from:", conn.RemoteAddr())
	s := server.NewServer(tcpconn.NewTCPConn(conn))
	if err := s.Handle(conn); err != nil {
		fmt.Println("Error handling connection:", err)
	}

	conn.Close()
}

func waitForShutdownSignal(cancel context.CancelFunc) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-signalChannel
	fmt.Printf("Received signal: %s\n", sig)

	// Cancel the context
	cancel()
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}

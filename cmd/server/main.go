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

type ServerConfig struct {
	Port string
}

func main() {
	err := pkg.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	serverPort, err := pkg.GetEnv("SERVER_PORT")
	if err != nil {
		log.Fatalf("Error getting SERVER_PORT: %s", err)
	}

	if err := setMaxFileDescriptorsLimit(); err != nil {
		log.Fatalf("Error setting max file descriptors limit: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	config := &ServerConfig{Port: serverPort}
	go startTCPServer(ctx, config, &wg)
	waitForShutdownSignal(cancel)

	wg.Wait()
	log.Println("Server shutdown complete.")
}

func startTCPServer(ctx context.Context, config *ServerConfig, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Printf("error listening: %v", err)
		return
	}
	defer listener.Close()

	log.Printf("Server listening on port: %s \n", config.Port)

	s := server.NewServer()
	// Receive incoming connections in a non-blocking way
	connChan := make(chan net.Conn)
	errChan := make(chan error)

	for {
		select {
		case <-ctx.Done():
			// Stop server if context is cancelled
			return
		default:
		}

		go func() {
			conn, err := listener.Accept()
			if err != nil {
				errChan <- err
				return
			}

			connChan <- conn
		}()

		select {
		case <-ctx.Done():
			log.Println("Shutting down server...")
			return
		case conn := <-connChan:
			log.Println("Accepted connection from:", conn.RemoteAddr())
			go handleConnection(conn, s)
		case err := <-errChan:
			log.Println("Error accepting connection:", err)
		}
	}
}

func handleConnection(conn net.Conn, s *server.Server) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	if err := s.Handle(tcpconn.NewTCPConn(conn)); err != nil {
		log.Println("Error handling connection:", err)
	}
}

func waitForShutdownSignal(cancel context.CancelFunc) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-signalChannel
	log.Printf("Received signal: %s\n", sig)

	// Cancel the context
	cancel()
}

// setMaxFileDescriptorsLimit sets the maximum number of open file descriptors
func setMaxFileDescriptorsLimit() error {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}

	log.Printf("Set description limit to %d \n", rLimit.Cur)
	return nil
}

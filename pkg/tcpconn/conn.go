package tcpconn

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

var ErrTooManyOpenFiles = errors.New("too many open files")
var ErrConnectionReset = errors.New("connection reset by peer")

type Validatable interface {
	IsValid() bool
}

type TCPConn struct {
	conn net.Conn
}

func NewTCPConn(conn net.Conn) *TCPConn {
	return &TCPConn{conn: conn}
}

func (t *TCPConn) Send(data any) error {
	message, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}

	if err := t.writeConn(t.conn, message); err != nil {
		return fmt.Errorf("error writing to connection: %w", err)
	}

	return nil
}

func (t *TCPConn) Read(target Validatable) error {
	resp, err := t.scanConnection(t.conn)
	if err != nil {
		return fmt.Errorf("error scanning connection: %w", err)
	}

	err = json.Unmarshal(resp, target)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	if !target.IsValid() {
		return fmt.Errorf("invalid response: %v", target)
	}

	return nil
}

func (t *TCPConn) writeConn(conn net.Conn, message []byte) error {
	// check if message has a newline character
	if message[len(message)-1] != '\n' {
		message = append(message, '\n')
	}

	_, err := conn.Write(message)
	return err
}

func (t *TCPConn) scanConnection(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, handleTCPConnErr(err)
	}

	return buffer[:n], nil
}

func handleTCPConnErr(err error) error {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if opErr.Err.Error() == "too many open files" {
			return ErrTooManyOpenFiles
		}

		if opErr.Err.Error() == "read: connection reset by peer" {
			return ErrConnectionReset
		}
	}
	return err
}

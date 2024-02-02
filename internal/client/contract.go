package client

import "github.com/aed86/word_of_wisdom_2/pkg/tcpconn"

type Conn interface {
	Read(target tcpconn.Validatable) error
	Send(data any) error
}

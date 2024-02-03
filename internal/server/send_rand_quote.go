package server

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (s *Server) sendRandQuote(tcpConn Conn) error {
	quote := getRandomQuote()

	if err := tcpConn.Send(dto.QuoteResp{Quote: quote}); err != nil {
		return fmt.Errorf("error writing quote response: %w", err)
	}

	fmt.Println("Quote was sent to the client")

	return nil
}

func (s *Server) sendQuoteErr(tcpConn Conn) error {
	fmt.Println("Invalid proof of work. Sending error message to client.")

	if err := tcpConn.Send(dto.QuoteResp{ErrorMessage: "Invalid proof of work."}); err != nil {
		return fmt.Errorf("error sending quote response: %w", err)
	}

	fmt.Println("Error message was sent to the client.")

	return nil
}

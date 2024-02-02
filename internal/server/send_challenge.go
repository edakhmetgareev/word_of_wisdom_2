package server

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (s *Server) sendChallenge(data dto.ChallengeResp) error {
	if err := s.tcpConn.Send(data); err != nil {
		return fmt.Errorf("error writing challenge response: %w", err)
	}

	fmt.Println("Challenge was sent to the client")

	return nil
}

package server

import (
	"fmt"
	"net"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

const difficulty = 5 // Number of leading zeros in the hash of the challenge

func (s *Server) Handle(conn net.Conn) error {
	challenge, err := generateChallenge()
	if err != nil {
		return fmt.Errorf("error generating challenge: %w", err)
	}

	r := dto.ChallengeResp{
		Challenge:    challenge,
		LeadingZeros: difficulty,
	}

	if err := s.sendChallenge(r); err != nil {
		return fmt.Errorf("error sending challenge: %w", err)
	}

	// Receive proof from client
	proof, err := s.readClientProof()
	if err != nil {
		return fmt.Errorf("error getting client proof: %w", err)
	}

	// Check if the proof of work is valid
	if !isValidProof(challenge, proof.Proof) {
		return s.sendQuoteErr()
	}

	// Send random quote to client
	if err := s.sendRandQuote(); err != nil {
		return fmt.Errorf("error sending quote: %w", err)
	}

	return nil
}

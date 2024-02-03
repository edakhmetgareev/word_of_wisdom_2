package server

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

const difficulty = 5 // Number of leading zeros in the hash of the challenge

func (s *Server) Handle(conn Conn) error {
	challenge, err := generateChallenge()
	if err != nil {
		return err
	}

	r := dto.ChallengeResp{
		Challenge:    challenge,
		LeadingZeros: difficulty,
	}

	err = s.sendChallenge(conn, r)
	if err != nil {
		return err
	}

	// Receive proof from client
	proof, err := s.readClientProof(conn)
	if err != nil {
		return err
	}

	// Check if the proof of work is valid
	if !isValidProof(challenge, proof.Proof) {
		return s.sendQuoteErr(conn)
	}

	// Send random quote to client
	if err := s.sendRandQuote(conn); err != nil {
		return fmt.Errorf("error sending quote: %w", err)
	}

	return nil
}

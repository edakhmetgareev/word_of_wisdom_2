package server

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (s *Server) readClientProof() (*dto.ProofResp, error) {
	var data dto.ProofResp
	if err := s.tcpConn.Read(&data); err != nil {
		return nil, fmt.Errorf("error reading challenge response: %w", err)
	}

	fmt.Printf("Response received from client: %+v\n", data)

	return &data, nil
}

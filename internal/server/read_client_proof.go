package server

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (s *Server) readClientProof(tcpConn Conn) (*dto.ProofResp, error) {
	var data dto.ProofResp
	if err := tcpConn.Read(&data); err != nil {
		return nil, fmt.Errorf("error reading proof response: %w", err)
	}

	return &data, nil
}

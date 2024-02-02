package client

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (c *Client) sendProof(proof string) error {
	if err := c.tcpConn.Send(dto.ProofResp{Proof: proof}); err != nil {
		return fmt.Errorf("error writing proof response: %w", err)
	}

	return nil
}

package client

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (c *Client) readChallenge() (*dto.ChallengeResp, error) {
	var resp dto.ChallengeResp
	if err := c.tcpConn.Read(&resp); err != nil {
		return nil, fmt.Errorf("error reading challenge response: %w", err)
	}

	return &resp, nil
}

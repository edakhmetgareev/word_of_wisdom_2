package client

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (c *Client) readQuote() (*dto.QuoteResp, error) {
	var resp dto.QuoteResp
	if err := c.tcpConn.Read(&resp); err != nil {
		return nil, fmt.Errorf("error reading quote response: %w", err)
	}

	if resp.ErrorMessage != "" {
		return nil, fmt.Errorf("error reading quote response: %s", resp.ErrorMessage)
	}

	return &resp, nil
}

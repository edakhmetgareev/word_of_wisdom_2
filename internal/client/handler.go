package client

import (
	"fmt"

	"github.com/aed86/word_of_wisdom_2/internal/dto"
)

func (c *Client) Handle(conn Conn, i int) error {
	data, err := c.readChallenge(conn)
	if err != nil {
		return fmt.Errorf("error reading challenge: %w", err)
	}

	proof, err := findProof(data.Challenge, data.LeadingZeros)
	if err != nil {
		return fmt.Errorf("error finding proof: %w", err)
	}

	if err := conn.Send(dto.ProofResp{Proof: proof}); err != nil {
		return fmt.Errorf("error writing proof response: %w", err)
	}

	var quote dto.QuoteResp
	if err := conn.Read(&quote); err != nil {
		return fmt.Errorf("error reading quote response: %w", err)
	}

	if quote.ErrorMessage != "" {
		return fmt.Errorf("error reading quote response: %s", quote.ErrorMessage)
	}

	fmt.Printf("I have received a quote: \"%s\" \n", quote.Quote)

	return nil
}

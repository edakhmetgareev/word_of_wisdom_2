package client

import (
	"fmt"
)

func (c *Client) Handle() error {
	data, err := c.readChallenge()
	if err != nil {
		return fmt.Errorf("error reading challenge: %w", err)
	}

	proof, err := findProof(data.Challenge, data.LeadingZeros)
	if err != nil {
		return fmt.Errorf("error finding proof: %w", err)
	}

	if err := c.sendProof(proof); err != nil {
		return fmt.Errorf("error sending proof: %w", err)
	}

	quote, err := c.readQuote()
	if err != nil {
		return fmt.Errorf("error reading quote: %w", err)
	}

	fmt.Printf("I have received a quote: \"%s\" \n", quote.Quote)

	return nil
}

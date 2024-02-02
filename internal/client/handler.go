package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/aed86/proof_of_work/internal/dto"
)

func Handle(conn net.Conn) error {
	data, err := readChallenge(conn)
	if err != nil {
		return fmt.Errorf("error reading challenge: %w", err)
	}

	err = findAndWriteProof(conn, data.Challenge, data.LeadingZeros)
	if err != nil {
		return fmt.Errorf("error finding and writing proof: %w", err)
	}

	quote, err := readQuote(conn)
	if err != nil {
		return fmt.Errorf("error reading quote: %w", err)
	}

	fmt.Println(quote)

	return nil
}

func readChallenge(conn net.Conn) (*dto.ChallengeResp, error) {
	resp, err := scanConnection(conn)
	if err != nil {
		return nil, fmt.Errorf("error reading challenge: %w", err)
	}

	result, err := unmarshalChallengeResp(resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling challenge response: %w", err)
	}

	if !result.IsValid() {
		return nil, fmt.Errorf("invalid challenge response: %v", result)
	}

	return &result, nil
}

func readQuote(conn net.Conn) (string, error) {
	resp, err := scanConnection(conn)
	if err != nil {
		return "", fmt.Errorf("error scanning connection: %w", err)
	}

	resp = strings.TrimPrefix(resp, "Quote:")
	result, err := unmarshalQuoteResp(resp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling quote response: %w", err)
	}

	if result.ErrorMessage != "" {
		return "", fmt.Errorf("error message from server: %s", result.ErrorMessage)
	}

	return result.Quote, nil
}

func scanConnection(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return scanner.Text(), nil
}

func unmarshalQuoteResp(resp string) (dto.QuoteResp, error) {
	var result dto.QuoteResp
	err := json.Unmarshal([]byte(resp), &result)
	return result, err
}

func unmarshalChallengeResp(resp string) (dto.ChallengeResp, error) {
	var data dto.ChallengeResp
	err := json.Unmarshal([]byte(resp), &data)
	return data, err
}

func findAndWriteProof(conn net.Conn, challenge string, leadingZeros int) error {
	proof, err := findProof(challenge, leadingZeros)
	if err != nil {
		return fmt.Errorf("error finding proof: %w", err)
	}

	_, err = fmt.Fprintf(conn, "%s\n", proof)
	return err
}

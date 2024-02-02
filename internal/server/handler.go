package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/aed86/proof_of_work/internal/dto"
)

const difficulty = 4 // Number of leading zeros in the hash of the challenge

func Handle(conn net.Conn) error {
	defer conn.Close()

	challenge, err := generateChallenge()
	if err != nil {
		return fmt.Errorf("error generating challenge: %w", err)
	}

	fmt.Println("Challenge generated:", challenge)

	jsonResp, err := prepareChallengeResp(challenge, difficulty)
	if err != nil {
		return err
	}

	jsonResp = append(jsonResp, []byte("\n")...)
	_, err = conn.Write(jsonResp)
	if err != nil {
		return fmt.Errorf("error sending challenge: %w", err)
	}
	fmt.Println("Challenge sent to client")

	// Read the response from the client
	response, err := readClientResponse(conn)
	if err != nil {
		return fmt.Errorf("error reading client response: %w", err)
	}

	fmt.Println("Response received from client:", response)

	// Check if the proof of work is valid and write the response
	err = validateProofAndWriteResponse(conn, challenge, response)
	if err != nil {
		return fmt.Errorf("error validating proof and writing response: %w", err)
	}

	return nil
}

func verifyPoW(challenge, response string) bool {
	// Сравнение ответа клиента с ожидаемым значением PoW
	return strings.HasPrefix(response, challenge) && len(response) == len(challenge)+difficulty
}

func prepareChallengeResp(challenge string, difficulty int) ([]byte, error) {
	resp := dto.ChallengeResp{
		Challenge:    challenge,
		LeadingZeros: difficulty,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("error marshalling challenge response: %w", err)
	}

	return jsonResp, nil
}

func readClientResponse(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	return scanner.Text(), scanner.Err()
}

func validateProofAndWriteResponse(conn net.Conn, challenge string, response string) error {
	if isValidProof(challenge, response) {
		return sendQuote(conn)
	}

	return sendQuoteErr(conn)
}

func sendQuote(conn net.Conn) error {
	fmt.Println("Valid proof of work. Sending quote to client.")
	quote := getRandomQuote()
	jsonResponse, err := json.Marshal(dto.QuoteResp{
		Quote: quote,
	})
	if err != nil {
		return fmt.Errorf("error marshalling quote response: %w", err)
	}

	_, err = conn.Write(jsonResponse)
	if err != nil {
		return fmt.Errorf("error sending quote response: %w", err)
	}

	return nil
}

func sendQuoteErr(conn net.Conn) error {
	fmt.Println("Invalid proof of work. Sending error message to client.")

	jsonResponse, err := json.Marshal(dto.QuoteResp{
		ErrorMessage: "Invalid proof of work.",
	})
	if err != nil {
		return fmt.Errorf("error marshalling quote response: %w", err)
	}

	_, err = conn.Write(jsonResponse)
	if err != nil {
		return fmt.Errorf("error sending quote response: %w", err)
	}

	return err
}

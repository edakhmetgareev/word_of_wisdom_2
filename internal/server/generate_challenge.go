package server

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateChallenge() (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("error generating random bytes: %w", err)
	}
	return base64.StdEncoding.EncodeToString(randomBytes), nil
}

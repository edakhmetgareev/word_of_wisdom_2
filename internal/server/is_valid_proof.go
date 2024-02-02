package server

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

var leadingZeros = strings.Repeat("0", difficulty)

func isValidProof(challenge string, response string) bool {
	data := challenge + response
	hash := sha256.Sum256([]byte(data))
	hashString := hex.EncodeToString(hash[:])

	// Check if the hash has the required number of leading zeros
	return strings.HasPrefix(hashString, leadingZeros)
}

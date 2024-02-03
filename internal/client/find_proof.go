package client

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const maxAttempts = 1000000

func findProof(challenge string, difficulty int) (string, error) {
	leadingZeros := strings.Repeat("0", difficulty)

	// easiest way to solve the problem
	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		attempt := challenge + strconv.Itoa(nonce)
		// calculate the hash of the attempt
		hash := calculateHash(attempt)

		// check if the hash has the required number of leading zeros
		if strings.HasPrefix(hash, leadingZeros) {
			return attempt, nil
		}
	}

	return "", fmt.Errorf("could not find a valid proof after %d attempts", maxAttempts)
}

func calculateHash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

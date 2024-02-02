package client

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

const maxAttempts = 1000000

func findProof(challenge string, difficulty int) (string, error) {
	var nonce int64
	leadingZeros := strings.Repeat("0", difficulty)
	startedAt := time.Now()

	defer func() {
		fmt.Println("Total time elapsed to find proof:", time.Since(startedAt))
	}()

	for i := 0; i < maxAttempts; i++ {
		rnd, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<63-1))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %w", err)
		}

		nonce = rnd.Int64()

		proof := challenge + strconv.FormatInt(nonce, 10)
		hash := sha256.Sum256([]byte(proof))
		hashString := hex.EncodeToString(hash[:])

		// Check if the hash has the required number of leading zeros
		if strings.HasPrefix(hashString, leadingZeros) {
			return strconv.FormatInt(nonce, 10), nil
		}
	}

	return "", fmt.Errorf("could not find a valid proof after %d attempts", maxAttempts)
}

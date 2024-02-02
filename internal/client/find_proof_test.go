package client

import (
	"testing"
)

func BenchmarkFindProof(b *testing.B) {
	challenge := "challenge"
	difficulty := 5

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = findProof(challenge, difficulty)
	}
}

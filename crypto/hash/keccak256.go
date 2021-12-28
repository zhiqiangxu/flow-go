package hash

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

// keccak_256Algo, embeds commonHasher
type keccak_256Algo struct {
	hash.Hash
}

// NewKECCAK_256 returns a new instance of KECCAK-256 hasher
func NewKECCAK_256() Hasher {
	return &keccak_256Algo{
		Hash: sha3.NewLegacyKeccak256()}
}

func (s *keccak_256Algo) Algorithm() HashingAlgorithm {
	return KECCAK_256
}

// ComputeHash calculates and returns the KECCAK-256 digest of the input.
// The function updates the state (and therefore not thread-safe)
// but does not reset the state to allow further writing.
func (s *keccak_256Algo) ComputeHash(data []byte) Hash {
	s.Reset()
	// `Write` delegates this call to sha256.digest's `Write` which does not return an error.
	_, _ = s.Write(data)
	return s.Sum(nil)
}

// SumHash returns the KECCAK-256 output.
// It does not reset the state to allow further writing.
func (s *keccak_256Algo) SumHash() Hash {
	return s.Sum(nil)
}

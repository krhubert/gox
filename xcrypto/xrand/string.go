package xrand

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
)

// GenerateKey returns random hex key of given size.
func GenerateKey(size int) ([]byte, error) {
	key, err := bufio.NewReader(rand.Reader).Peek(size)
	if err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(key))[size:], nil
}

package xrand

import "testing"

func TestGenerateKeyLength(t *testing.T) {
	var tests = []struct {
		size int
	}{{1}, {2}, {5}, {60}, {1000}}

	for _, tt := range tests {
		key, err := GenerateKey(tt.size)
		if err != nil {
			t.Fatal(err)
		}
		if len(key) != tt.size {
			t.Fatalf("GenerateKey(%d) invalid key length - got: %d, epxected: %d", tt.size, len(key), tt.size)
		}
	}
}

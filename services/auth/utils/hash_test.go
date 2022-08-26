package utils

import (
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	plaintext := "password"
	hash, _ := HashPassword(plaintext)

	if plaintext == hash {
		t.Errorf("expected plaintext and hash to differ but both equal %s", plaintext)
	}

	if !CheckPasswordHash(plaintext, hash) {
		t.Errorf("expected hash %s to resolve to plaintext %s", hash, plaintext)
	}
}

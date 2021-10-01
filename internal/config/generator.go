package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	// EncryptionKeyByteLen is the length (in bytes) of the generated encryption key. The
	// value of 32 indicates that AES-256 encryption should be used.
	EncryptionKeyByteLen = 32

	// SessionSecretByteLen is the length (in bytes) of the generates session secret. We
	// follow the recommended length of 64 bytes from github.com/gorilla/sessions, the
	// underlying session package used by Response.
	SessionSecretByteLen = 64
)

func GenerateEncryptionKey() (string, error) {
	rawKey := make([]byte, 32)

	_, err := rand.Read(rawKey)
	if err != nil {
		return "", fmt.Errorf("generate key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(rawKey), nil
}

func GenerateSessionSecret() (string, error) {
	rawKey := make([]byte, 64)

	_, err := rand.Read(rawKey)
	if err != nil {
		return "", fmt.Errorf("generate key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(rawKey), nil
}

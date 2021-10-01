package response

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func (c *Core) Encrypt(plainText []byte) ([]byte, error) {
	return encryptBytes(c.config.EncryptionKey, plainText)
}

func (c *Core) EncryptString(plainText string) ([]byte, error) {
	return encryptBytes(c.config.EncryptionKey, []byte(plainText))
}

func encryptBytes(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to encrypt bytes: aes.NewCipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to encrypt bytes: cipher.NewGCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, fmt.Errorf("unable to encrypt bytes: invalid nonce size: %w", err)
	}

	return gcm.Seal(nonce, nonce, plainText, nil), nil
}

func (c *Core) Decrypt(plainText []byte) ([]byte, error) {
	return decryptBytes(c.config.EncryptionKey, plainText)
}

func (c *Core) DecryptString(cipherText string) ([]byte, error) {
	return decryptBytes(c.config.EncryptionKey, []byte(cipherText))
}

func decryptBytes(key []byte, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to decrypt bytes: aes.NewCipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to decrypt bytes: cipher.NewGCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return []byte{}, fmt.Errorf("unable to decrypt bytes: invalid nonce size: %w", err)
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to decrypt bytes: gcm.open: %w", err)
	}

	return plainText, nil
}

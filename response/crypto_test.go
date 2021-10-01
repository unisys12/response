package response

import (
	"testing"

	"github.com/matryer/is"
)

var testEncryptionKey = "WQYLN3gBAgYaTe5b7RtZgKw+FpnGAlmlxyQxiLK6YWo="

func Test_EncryptDecrypt(t *testing.T) {
	is := is.New(t)

	plainText := []byte("plain text")
	encrypted, err := encryptBytes(testEncryptionKey, plainText)
	is.NoErr(err)

	decryptedText, err := decryptBytes(testEncryptionKey, encrypted)
	is.NoErr(err)

	is.Equal(plainText, decryptedText)
}

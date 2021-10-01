package response

import (
	"encoding/base64"
	"testing"

	"github.com/matryer/is"
)

var testEncryptionKey = "WQYLN3gBAgYaTe5b7RtZgKw+FpnGAlmlxyQxiLK6YWo="

func Test_EncryptDecrypt(t *testing.T) {
	is := is.New(t)

	encryptionKey, err := base64.StdEncoding.DecodeString(testEncryptionKey)
	is.NoErr(err)

	plainText := []byte("plain text")
	encrypted, err := encryptBytes(encryptionKey, plainText)
	is.NoErr(err)

	decryptedText, err := decryptBytes(encryptionKey, encrypted)
	is.NoErr(err)

	is.Equal(plainText, decryptedText)
}

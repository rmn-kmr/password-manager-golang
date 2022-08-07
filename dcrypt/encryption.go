package dcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(text, mySecret string) (string, error) {
	// Create the AES cipher
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	// Byte array of the string
	plainText := []byte(text)
	// Return an encrypted stream
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	// Encrypt bytes from plaintext to ciphertext
	cfb.XORKeyStream(cipherText, plainText)
	return encode(cipherText), nil
}

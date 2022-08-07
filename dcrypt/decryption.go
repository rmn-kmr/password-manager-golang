package dcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

func decode(s string) ([]byte, error) {
	// DecodeString returns the bytes represented by the base64 string.
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return data, nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	// Create the AES cipher
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText, err := decode(text)
	if err != nil {
		return "", err
	}
	// Return a decrypted stream
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	// Decrypt bytes from ciphertext
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

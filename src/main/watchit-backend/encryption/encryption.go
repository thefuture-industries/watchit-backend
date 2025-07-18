package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

var key []byte
var iv []byte

// Шифрование данных
func Encrypt(plaintext string) (string, error) {
	key, _ = base64.StdEncoding.DecodeString(os.Getenv("SUPER_SECRET_KEY"))
	iv, _ = base64.StdEncoding.DecodeString(os.Getenv("IV"))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error when creating an encryption block")
	}

	gsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error when creating the block cipher mode")
	}

	ciphertext := gsm.Seal(nil, []byte(iv), []byte(plaintext), nil)

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	return encodedCiphertext, nil
}

// Де-шифрование данных
func Decrypt(encodedCiphertext string) (string, error) {
	key, _ = base64.StdEncoding.DecodeString(os.Getenv("SUPER_SECRET_KEY"))
	iv, _ = base64.StdEncoding.DecodeString(os.Getenv("IV"))

	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", fmt.Errorf("error in decrypting encrypted data")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error when creating an encryption block")
	}

	gsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error when creating the block cipher mode")
	}

	plaintext, err := gsm.Open(nil, []byte(iv), ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("data decryption error")
	}

	return string(plaintext), nil
}

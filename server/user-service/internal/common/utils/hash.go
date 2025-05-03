// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

var key string
var iv string

func Encrypt(plaintext string) (string, error) {
	key = os.Getenv("SUPER_SECRET_KEY")
	iv = os.Getenv("IV")

	// Создание блока шифрования AES
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error creating the encryption block")
	}

	// Создание блочного режима шифрования CBC
	gsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating block encryption mode")
	}

	// Шифрование данных
	ciphertext := gsm.Seal(nil, []byte(iv), []byte(plaintext), nil)

	// Кодирование зашифрованных данных в base64 для удобства
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	return encodedCiphertext, nil
}

func Decrypt(encodedCiphertext string) (string, error) {
	key = os.Getenv("SUPER_SECRET_KEY")
	iv = os.Getenv("IV")

	// Декодирование зашифрованных данных из base64
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", fmt.Errorf("error decoding encrypted data")
	}

	// Создание блока шифрования AES
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error creating the encryption block")
	}

	// Создание блочного режима шифрования CBC
	gsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating block encryption mode")
	}

	// Расшифровка данных
	plaintext, err := gsm.Open(nil, []byte(iv), ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("data decryption error")
	}

	return string(plaintext), nil
}

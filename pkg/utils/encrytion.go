package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"

	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
)

func AESEncryption(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, &entity.CError{
			Err:        err,
			StatusCode: 500,
		}
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, &entity.CError{
			Err:        err,
			StatusCode: 500,
		}
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return ciphertext, nil
}

func ToByte32(keyBytes []byte) ([]byte, error) {
	if len(keyBytes) < 32 {
		padding := make([]byte, 32-len(keyBytes))
		keyBytes = append(keyBytes, padding...)
		return keyBytes, nil
	}
	return nil, &entity.CError{
		Err:        errors.New("key length must be 32 bytes"),
		StatusCode: 500,
	}

}

func AESDecryption(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, &entity.CError{
			Err:        err,
			StatusCode: 500,
		}
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func AESHashCompared(plainText []byte, cipherText []byte, key []byte) error {
	decrypted, err := AESDecryption(cipherText, key)
	if err != nil {
		return &entity.CError{
			Err:        err,
			StatusCode: 500,
		}
	}

	if string(plainText) != string(decrypted) {
		log.Println(string(plainText), string(decrypted))
		return &entity.CError{
			Err:        errors.New("error, password is invalid"),
			StatusCode: 400,
		}
	}

	return nil
}

package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("The key parameter was not passed to the environment variables")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encryp(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	fmt.Println(err)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decryp(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}

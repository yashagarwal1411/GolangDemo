package apputils

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	b64 "encoding/base64"
)

func Encrypt(plaintext []byte, key []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	sEnc := b64.URLEncoding.EncodeToString([]byte(gcm.Seal(nonce, nonce, plaintext, nil)))
	return sEnc, nil
}

func Decrypt(sEnc string, key []byte) ([]byte, error) {
	ciphertext, _ := b64.URLEncoding.DecodeString(sEnc)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

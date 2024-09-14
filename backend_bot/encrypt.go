package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"log"
)

func encryptFlag(flag string, key string) string {
	buf := AESEncrypt([]byte(flag), []byte(key))
	if buf == nil {
		return "failed to generate cipher, report to admins"
	}
	return base64.StdEncoding.EncodeToString(buf)
}

func AESEncrypt(content []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return nil
	}
	iv := make([]byte, 16)
	ecb := cipher.NewCBCEncrypter(block, iv)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func generateSecretKey() string {
	t := make([]byte, 8)
	rand.Read(t)
	return hex.EncodeToString(t)
}

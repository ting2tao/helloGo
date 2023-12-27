package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("AAAAAAAAAAAAAA") // 16 bytes key
	ciphertext, _ := base64.StdEncoding.DecodeString("AAAAAAAAAAAAAAAAAAAAA=")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%s\n", ciphertext)
}

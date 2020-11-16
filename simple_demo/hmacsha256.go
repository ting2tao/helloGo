package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	message = "zoekwui1hnmg49x5fwzf5la0ml5dziwn&&1542355862990"
	secret  = "gywzffojtnzl0vd6kcut8fcgyud5wg49"
)

//HMac_Sha256加密
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {

	fmt.Println(ComputeHmacSha256(message, secret))

}

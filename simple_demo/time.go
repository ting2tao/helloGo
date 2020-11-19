package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano() / 1e6)
	fmt.Println(time.Now())
	sellout := "sellout"
	fmt.Println(sellout)
	fmt.Println(hex.EncodeToString([]byte("A")))

	zeroTime, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	expiration := zeroTime.AddDate(0, 0, 1).Sub(time.Now())
	fmt.Println(expiration)
}

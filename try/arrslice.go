package main

import (
	"fmt"
	"reflect"
)

func main() {
	joinAvatar := []string{
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/02b5e2970926b4401e0bb5e4a942f335.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/02d828d6c62119ec0f481254183c6574.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/02e167d9ae38638b408159ca66ea9b62.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/03459e33424e0cef85b43b14fd036038.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/05f14576b6d0e621dcc34d2e65302698.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/06124a06f5664d74531a26abb32f0e8f.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/06a3ef84b78b6b48fadbbd6acf74b274.jpg",
		"https://xiniao-upload.oss-accelerate.aliyuncs.com/AI/071ae2d5cd0f48bbfa64e2549977e584.jpg",
	}

	useAvatar := joinAvatar[0:8]
	fmt.Println(useAvatar)

	var intNum int8 = 64
	fmt.Println(reflect.TypeOf(intNum))
	fmt.Println(intNum)
}

package main

import (
	"fmt"
	"strconv"
)

func main()  {
	for i:=0;i<10;i++{
		fmt.Println("i的值"+ strconv.Itoa(i))
	}
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}


	str3 := "中国"
	fmt.Printf("The length of str3 is: %d\n", len(str3))
	for ix :=0; ix < len(str3); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str3[ix])
	}
}

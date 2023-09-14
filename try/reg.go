package main

import (
	"fmt"
	"regexp"
)

func main() {
	newName := "（已于[20230823171100]被合并）sadad"
	newBasicName := regexp.MustCompile(`（[^（）]*）`).ReplaceAllString(newName, "")
	fmt.Println(newBasicName)
}

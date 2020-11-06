package main

import (
	"fmt"
)

type React struct {
	Width, Height float64
}

func (react React) area() float64 {
	return react.Height * react.Width
}
func main() {
	react := React{10, 30}
	fmt.Println("width", react.Width, "height:", react.Height, "react area:", react.area())
	countdown := "HA"
	fmt.Println("keyword", countdown)
}

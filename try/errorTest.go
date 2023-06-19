package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("数据已存在")
	err2 := fmt.Errorf("保存新数据失败: %w", err)

	if errors.Is(err2, err) {
		fmt.Println("EQ")
	} else {
		fmt.Println("NE")
	}
}

package main

import (
	"fmt"
	"os"
)

func writeStringToFile(filename, content string) error {
	// 将字符串转换为字节数组
	data := []byte(content)

	// 将字节数组写入文件
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("已成功将内容写入文件：%s\n", filename)
	return nil
}

func appendStringToFile(filename, content string) error {
	// 打开文件以追加内容
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将字符串转换为字节数组
	data := []byte(content + "\n") // 添加换行符

	// 写入文件
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	fmt.Printf("已成功将内容追加到文件：%s\n", filename)
	return nil
}

func main() {
	filename := "D:/workSpace/goworkspace/src/hello/try/copy.txt"
	content := "Hello, World!"

	err := appendStringToFile(filename, content)
	if err != nil {
		fmt.Printf("写入文件时出错：%v\n", err)
	}
	err = appendStringToFile(filename, content)
	if err != nil {
		fmt.Printf("写入文件时出错：%v\n", err)
	}
}

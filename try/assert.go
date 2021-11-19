package main

import (
	"fmt"
	"math"
)

func main() {
	projectBA := make(map[string]interface{})
	projectBA["project_type"] = 1
	projectType, ok := projectBA["project_type"].(string)
	if !ok {
		fmt.Println("断言失败")
		//continue
	}
	math.Pow(2, 2)
	math.Sqrt(54)
	fmt.Println(projectType)
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(54))
}

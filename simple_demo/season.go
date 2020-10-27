package main

import "fmt"

func main() {
	season := season(16)
	fmt.Println("当前季节为：" + season)
	fall(7)
}

func season(month int) string {
	str := ""
	switch month {
	case 1, 2, 3:
		str = "春"
	case 4, 5, 6:
		str = "夏"
	case 7, 8, 9:
		str = "秋"
	case 10, 11, 12:
		str = "冬"
	default:
		return "没有这个季节"
	}
	return str + "季"
}

func  fall(k int)  {

	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
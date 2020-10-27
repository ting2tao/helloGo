package main

import "fmt"

func main() {
	i := 10
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i <= 0 {
			break
		}
	}
	for4()
	for5()
	for6()
	goto1()
	for7()
	for8()
}
func for4()  {
	for i:=0; i<3; i++ {
		for j:=0; j<10; j++ {
			if j>5 {
				break
			}
			print(j)
		}
		print("  ")
	}
}
func for5()  {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
}


/**
 标签
 */
func for6(){
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
func goto1 () {
	a := 1
	b := 9
	goto TARGET // compile error
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}

func for7()  {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 { break }
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}

func for8()  {
	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}
}
package main

var a2 string

func main() {
	a2 = "G"
	print(a2)
	f1()
}

func f1() {
	a2 := "O"
	print(a2)
	f2()
}

func f2() {
	print(a2)
}
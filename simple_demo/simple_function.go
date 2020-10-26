package main

import "fmt"
type A struct {
	id int
	name string

}

func main() {
	fmt.Printf("Multiply: 2*5*6 = %d\n",Multiply(2,5,6))

	fmt.Printf("funcB = %d\n",funcB(A{1,"hh"}))
}
func Multiply(a int , b int , c int) int{
	return a * b * c
}

func funcA(a *A)  {

	b = a
	fmt.Printf("funcA: = %s\n",b)
}



func funcB(a A) A {

	b = &a
	return b
}
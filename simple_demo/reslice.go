package main

import "fmt"

func main() {
	s1:=make([]int,3,6)
	fmt.Printf("%p\n",s1)
	s1=append(s1, 1,2,3,5)
	fmt.Printf("%v %p \n",s1,s1)

	a:=[]int{1,2,3,4,5}
	s3:=a[2:5]
	s4:=a[1:3]
	fmt.Println("s3:",s3,"s4",s4)
	fmt.Println("s4 len:",len(s4),"s4 cap",cap(s4))
	s4=append(s4, 1,25,5) // 容量不够 加倍
	fmt.Println("s4 len:",len(s4),"s4 cap",cap(s4))
	s3[0] = 10
	fmt.Println("s3:",s3,"s4",s4,"a",a)

	copy(s3,s4)
	fmt.Println(s3)


	//s5 = s3[0:]
	s5:=a[:]
	fmt.Printf("%v\n",s5)
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	//release函数中error的值并不会被defer的return返回，因为匿名返回值在defer执行前就已经声明好并复制为nil。
	//correctRelease函数能够修改返回值是因为闭包的特性，defer中的err是实际的返回值err地址引用，指向的是同一个变量。
	//defer修改程序返回值error一般用在和recover搭配中，上述的情况属于滥用defer的一种情况，其实error函数值可以直接在程序的return中修改，不用defer
	{
		err := release()
		fmt.Println(err)
	}

	{
		err := correctRelease()
		fmt.Println(err)
	}
}

func release() error {
	defer func() error {
		return errors.New("error")
	}()

	return nil
}

func correctRelease() (err error) {
	defer func() {
		err = errors.New("error")
	}()
	return nil
}

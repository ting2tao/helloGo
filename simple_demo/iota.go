package main

import "fmt"

const (
	//
	// 1-登录验证码  2-绑定手机号码 3-实名认证 4-提现 5-换绑手机号码  6-注销手机号码
	i = iota
	ForLogin
	// 2-绑定手机号码
	BindPhone
	// 3-实名认证
	Verified = 6
	// 4-提现
	CashOut
	// 5-换绑手机号码
	ChangePhone = iota
	// 6-注销手机号码
	Deregister
)

func main() {
	fmt.Println(ForLogin)
	fmt.Println(BindPhone)
	fmt.Println(Verified)
	fmt.Println(CashOut)
	fmt.Println(ChangePhone)
	fmt.Println(Deregister)
}

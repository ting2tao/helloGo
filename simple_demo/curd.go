package main

import (
	"fmt"
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Sex  uint
}
var db *gorm.DB
var users []User
func main() {
	fmt.Println("Hello JUNJUNXIA")
	dsn := "root:lastmoon@tcp(127.0.0.1:3306)/hello_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql",dsn)
	defer db.Close()
	db.AutoMigrate(&User{})

 	//user := User{Name:"Shi Chuntao2",Sex:1}
	//res := db.Create(&user)
	//fmt.Println(res.RowsAffected)
	//fmt.Println(user.Name)
	//fmt.Println(user.ID)
	//db.Find(&User{})
	//fmt.Println(user)
	//var users []User
	//db.Find(&users)
	//fmt.Println(users)
	//// query one
	//user := new (User)
	//db.First(user,1)
	//fmt.Println(user)
	db.Delete(&User{}, "10")
	//del()
}

func sel(){

	fmt.Println("Hello World")
	// query all
	//var users []User
	db.Find(&users)
	//fmt.Println(users)
	//// query one
	//user := new (User)
	//db.First(user,1)
	//fmt.Println(user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	//var user *User
	//user = new(User)
	//fmt.Println(user)
	//resF := db.First(user)
	//fmt.Println(resF.RowsAffected)
}

func update(u *User){

}

func del(){
	var testResult User
	db.Where("name = ?", "Shi Chuntao2").First(&testResult)
	fmt.Println("result: ", testResult)
	//db.Delete(&User{}, "11")
}
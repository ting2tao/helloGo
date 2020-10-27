package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Sex  uint
}
var db *gorm.DB

func main() {
	dsn := "root:lastmoon@tcp(127.0.0.1:3306)/hello_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql",dsn)
	db.AutoMigrate(&User{})

 	user := User{Name:"Shi Chuntao2",Sex:1}
	res := db.Create(&user)
	fmt.Println(res.RowsAffected)
	fmt.Println(user.Name)
	fmt.Println(user.ID)
	db.Find(&User{})
	fmt.Println(user)
	//Sel()
}

func Sel(){
	var user User

	// 可以
	db.First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	//var user *User
	//user = new(User)
	//fmt.Println(user)
	//resF := db.First(user)
	//fmt.Println(resF.RowsAffected)
}

func update(u *User){

}
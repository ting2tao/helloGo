package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type Apple struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	gorm.Model
}

func main() {

	db, _ = gorm.Open("mysql", "root:sct123@tcp(127.0.0.1:3306)/fastadmin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Apple{})

	r := gin.Default()

	r.POST("/getAll", GetProjects)
	r.POST("/get", GetPerson)
	r.POST("/add", CreatePerson)
	r.POST("/update", UpdatePerson)
	r.POST("/del", DeletePerson)

	r.Run("localhost:8080")
}

func GetProjects(c *gin.Context) {
	var people []Apple
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(people)
	} else {
		c.JSON(200, people)
	}
}

func GetPerson(c *gin.Context) {
	// 使用bind 获取json的参数
	var person Apple
	c.ShouldBindJSON(&person)
	id:= person.ID

	fmt.Println(id)
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func CreatePerson(c *gin.Context) {
	var person Apple
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(200, person)
}

func UpdatePerson(c *gin.Context) {
	var person Apple
	//var person2 Apple
	//c.ShouldBindJSON(&person)
	//id:= person2.ID
	//fmt.Println(id)
	c.ShouldBindJSON(&person)
	if err := db.First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	fmt.Println(person)

	db.Save(&person)
	c.JSON(200, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Apple
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
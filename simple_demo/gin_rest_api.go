package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB
var err error

type Person struct {
	gorm.Model
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
}

func main() {
	// NOTE: See we're using = to assign the global var
	// instead of := which would assign it only in this function
	// db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "root:lastmoon@tcp(127.0.0.1:3306)/hello_go?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	r := gin.Default()

	r.GET("/", GetProjects)
	r.GET("/people/:id", GetPerson)
	r.POST("/people", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)

	r.Run("127.0.0.1:8080")
}

func GetProjects(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, people)
	}
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func CreatePerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(200, person)
}

func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(200, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
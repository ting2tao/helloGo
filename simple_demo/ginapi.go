package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	//r.GET("/",func(c *gin.Context){
	//	c.String(200,"Hello Go")
	//})

	m1 := make(map[string]map[string]string)
	m1["data_obj"] = make(map[string]string)
	m1["data_obj"]["id"] = "2"
	datetime := "2020"
	m1["data_obj"]["name"] = datetime

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Hello Go",
			"data": m1,
		})
	})
	r.Run(":8082")
}

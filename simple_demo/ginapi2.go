package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
func	main(){

	r :=gin.Default()
	//r.GET("/",func(c *gin.Context){
	//	c.String(200,"Hello Go")
	//})

	m1:=make(map[string]map[string]string)
	//m1["data_list"] = make(map[string]string)
	//m1["data_list"]["id"] = "2"
	//m1["data_list"]["name"] = "史纯涛"

	//m1["data_obj"] = make(map[string]string)
	//m1["data_obj"]["id"] = "6"
	//m1["data_obj"]["name"] = "T的Q"

	for i:=0;i<6;i++{
		var item = make(map[string]string)
		j := strconv.Itoa(i)
		item["id"] = j
		item["name"] = fmt.Sprintf("史纯涛"+j)
		var _,ok = item["id"]
		fmt.Println(j)
		if ok{
			fmt.Println(item)
		}
		m1[j] = item
		fmt.Println(m1)
		//delete(item,"id")
		//delete(item,"name")
	}
	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"Hello Go",
			"data": m1,
		})
	})
	r.Run()
}

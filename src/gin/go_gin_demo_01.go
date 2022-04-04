package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.Writer.WriteString("hello gin")
	})

	// gin接收get请求中的参数
	r.GET("/hello2", func(c *gin.Context) {
		message := c.Query("message")
		c.Writer.WriteString("hello gin " + message)
	})

	// gin 解析post请求的表单数据
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		c.Writer.WriteString("message:" + message)
	})

	// gin 解析post请求的json数据
	// POST /json_post HTTP/1.1
	// Content-Type: application/json
	// User-Agent: PostmanRuntime/7.29.0
	// Accept: */*
	// Cache-Control: no-cache
	// Postman-Token: acc145c6-fc1a-4570-90a7-38c34b411744
	// Host: localhost:8080
	// Accept-Encoding: gzip, deflate, br
	// Connection: keep-alive
	// Content-Length: 56
	// {
	// "username":"admin",
	// "password":"1qaz2wsx"
	// }
	r.POST("/json_post", func(c *gin.Context) {
		json := make(map[string]interface{})
		c.BindJSON(&json)
		c.JSON(http.StatusOK, gin.H{
			"username" : json["username"],
			"password" : json["password"],
		})
	})

	// gin 解析post请求的json数据2
	// POST /json_post HTTP/1.1
	// Content-Type: application/json
	// User-Agent: PostmanRuntime/7.29.0
	// Accept: */*
	// Cache-Control: no-cache
	// Postman-Token: acc145c6-fc1a-4570-90a7-38c34b411744
	// Host: localhost:8080
	// Accept-Encoding: gzip, deflate, br
	// Connection: keep-alive
	// Content-Length: 56
	// {
	// "username":"admin",
	// "password":"1qaz2wsx"
	// }
	r.POST("/json_post2", func(c *gin.Context) {
		json := User{}
		c.BindJSON(&json)
		c.JSON(http.StatusOK, gin.H{
			"json2 username" : json.UserName,
			"json2 password" : json.Password,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

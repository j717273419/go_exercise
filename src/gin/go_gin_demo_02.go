package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "gin",
		})
	})
	r.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// gin路由分组
	g := r.Group("/group")
	g.GET("/user", func(c *gin.Context) {
		c.Writer.WriteString("user")
	})
	g.GET("/user2", func(c *gin.Context) {
		c.Writer.WriteString("user2")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

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

	g := r.Group("/group")
	//path中的参数
	// http://localhost:8080/group/user:charlie
	// user :charlie
	g.GET("/user:username", func(c *gin.Context) {
		username := c.Param("username")
		c.Writer.WriteString("user " + username)
	})
	
	//path中的参数
	// http://localhost:8080/group/user2/charlie
	// user2 charlie
	g.GET("/user2/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.Writer.WriteString("user2 " + username)
	})

		//path中的参数
	// http://localhost:8081/group/user3/charlie/msg
	// charlie user3 /msg
	g.GET("/user3/:username/*action", func(c *gin.Context) {
		username := c.Param("username")
		action := c.Param("action")
		c.Writer.WriteString(username + " user3 " + action)
	})
	//path中的参数
	// http://localhost:8080/group/user/:charlie
	// user :charlie
	g.GET("/user3/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.Writer.WriteString("user2 " + username)
	})

	// query中的参数
	// http://localhost:8081/admin
	// Hello kimi 
	// http://localhost:8081/admin?name=charlie&age=18
	// Hello charlie 18
	r.GET("/admin", func(c *gin.Context) {
		firstname := c.DefaultQuery("name", "kimi") // 获取query中的name，没有的话就为kim
		lastname := c.Query("age") // 获取query中的age
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

	})
	r.Run("localhost:8081") // 监听并在 http://localhost:8081 上启动服务
}

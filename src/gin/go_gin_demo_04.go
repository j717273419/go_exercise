package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin验证
	r.POST("/user", func(c *gin.Context) {
		var person = &Person{}
		if err := c.ShouldBind(person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    person,
		})
	})

	r.Run("localhost:8081") // 监听并在 http://localhost:8081 上启动服务
}
 
type Person struct {
	Name string `json:"name" binding:"required"`      // json格式从name取值，并且该值为必须的
	Age  int    `json:"age" binding:"required,gt=20"` // json格式从age取值，并且该值为必须的，且必须大于20
}

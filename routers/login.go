package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(e *gin.Engine)  {
	// {} 是书写规范
	//e.Group("/login")
	{
		e.GET("/login", loginHandler)
		e.GET("/submit", submit)
	}
}

// 定义接收数据的结构体
type Login struct {

	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func loginHandler(c *gin.Context)  {
	fmt.Println("1111")
	data:= map[string]interface{}{
		"name":"ddd",
		"msg":"ddd",
		"id":"ddd",
	}
	c.JSON(http.StatusOK, gin.H{"error": data})
}
func submit(c *gin.Context)  {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func loadShop(e *gin.Engine)  {
	e.GET("/hello", helloHandler)
	e.GET("/goods", goodsHandler)
	e.GET("/checkout", checkoutHandler)
}
func helloHandler(c *gin.Context)  {}
func goodsHandler(c*gin.Context)  {}
func checkoutHandler(c*gin.Context)  {}
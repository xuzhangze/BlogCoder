package view

import "github.com/gin-gonic/gin"

func HelloHandle(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg":"Hello World!",
	})
}

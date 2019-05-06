package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xuzhangze/BlogCoder/view"
)

func main() {
	r := gin.Default()
	r.GET("/blog/v1/hello", view.HelloHandle)
}

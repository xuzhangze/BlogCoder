package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

func TextEdit(c *gin.Context) {
	page := c.DefaultPostForm("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "edit.html", gin.H{"code" : 1})
		return
	}

	title := c.DefaultPostForm("title", "")
	text := c.DefaultPostForm("text", "")
	if len(title)==0 || len(text)==0 {
		logger.Warn("Parameter is invalid")
		c.HTML(http.StatusOK, "failed.html", gin.H{"code" : 0})
		return
	}


}

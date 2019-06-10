package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func TextModifyHandle(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if len(idStr) == 0 {
		logger.Warn("Invalid blog ID, blog ID: %v", idStr)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "invalid blog id",
		})
		return
	}

	blog, err := controller.GetBlogInfoByID(id)
	if err != nil {
		logger.Error("Get blog error, blog ID: %v, err: %v", id, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get blog info err",
		})
		return
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"id" : id,
		"text" : blog.GetText(),
	})
}

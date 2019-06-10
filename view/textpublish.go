package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func TextPublishHandle(c *gin.Context)  {
	idStr := c.DefaultQuery("id", "")
	uidStr, _ := c.Cookie("uid")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	uid, _ := strconv.ParseInt(uidStr, 10, 64)

	if len(idStr) == 0 {
		logger.Warn("Invalid blog ID, blog ID: %v", idStr)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "invalid blog id",
		})
		return
	}

	blog, err := controller.TextPublish(id, 1)
	if err != nil {
		logger.Error("Blog publish error, blog ID: %v, err: %v", id, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "unknow error",
		})
		return
	}
	user, err := controller.PersonalInfo(uid)
	if err != nil {
		logger.Warn("Get user personal info error, user ID: %v, err: %v", uid, err)
	}

	c.HTML(http.StatusOK, "detailpage.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"title" : blog.GetTitle(),
		"author" : user.GetUname(),
		"publish" : "yes",
		"id" : blog.GetId(),
		"text" : blog.GetText(),
		"star" : blog.GetStar(),
		"browse" : blog.GetBrowse(),
	})
}

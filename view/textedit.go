package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"github.com/xuzhangze/BlogCoder/middle"
	"net/http"
	"strconv"
)

func TextEditHandle(c *gin.Context) {
	uIdStr, err := c.Cookie("uid")
	if err != nil {
		logger.Warn("Get cookie error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get cookie error",
		})
	}

	page := c.DefaultQuery("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"code" : 1,
			"msg" : "success",
			"id" : 666,
			"text" : "Hello xblog!",
		})
		return
	}

	title := c.DefaultQuery("title", "")
	text := c.DefaultQuery("text", "")
	idStr := c.DefaultQuery("id", "")
	if len(title)==0 || len(text)==0 || len(idStr)==0 {
		logger.Warn("Parameter is invalid")
		c.HTML(http.StatusOK, "failed.html", gin.H{"code" : 0})
		return
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)
	uId, _ := strconv.ParseInt(uIdStr, 10, 64)

	var blog middle.TextInfo
	if len(idStr) < 9 {
		blog, err = controller.TextEdit(uId, title, text)
	} else {
		blog, err = controller.Updatetext(id, uId, title, text)
	}
	if err != nil {
		logger.Error("Edit blog error, user ID: %v, err: %v", uId, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "create blog error",
		})
		return
	}

	user, err := controller.PersonalInfo(uId)
	if err != nil {
		logger.Error("Get user info error, user ID: %v, err: %v", uId, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get user info error",
		})
		return
	}

	c.HTML(http.StatusOK, "detailpage.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"title" : blog.GetTitle(),
		"author" : user.GetUname(),
		"publish" : "no",
		"id" : blog.GetId(),
		"text" : blog.GetText(),
		"star" : blog.GetStar(),
		"browse" : blog.GetBrowse(),
	})
}

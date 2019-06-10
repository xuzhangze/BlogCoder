package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func TextDeleteHandle(c *gin.Context) {
	uidStr, _ := c.Cookie("uid")
	idStr := c.DefaultQuery("id", "")
	if len(idStr) == 0 {
		logger.Warn("Parameter is invalid")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "invalid parameter",
		})
		return
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)
	if err := controller.TextDelete(id); err != nil {
		logger.Error("Blog delete error, blog ID: %v, err: %v", id, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "unknow error",
		})
		return
	}

	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	user, err := controller.PersonalInfo(uid)
	if err != nil{
		logger.Error("Get user info error, user ID: %v, err: %v", uid, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get user info error",
		})
		return
	}

	c.HTML(http.StatusOK, "personalpage.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"title" : user.GetUname(),
	})
}

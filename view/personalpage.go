package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func UserPersonalPageHandle(c *gin.Context) {
	uidStr, err := c.Cookie("uid")
	if err != nil {
		logger.Warn("Get user cookie error")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get coockie error",
		})
		return
	}

	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	user, err := controller.PersonalInfo(uid)
	if err != nil {
		logger.Error("Get user info error, user ID: %v, err: %v", uid, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "unknow error",
		})
		return
	}

	c.HTML(http.StatusOK, "personalpage.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"title" : user.GetUname(),
	})
}

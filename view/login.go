package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func LoginHandle(c *gin.Context) {
	page := c.DefaultQuery("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "login.html", gin.H{"code" : 1})
		return
	}

	name := c.DefaultQuery("name", "")
	password := c.DefaultQuery("password", "")
	if len(name)==0 || len(password)==0 {
		logger.Warn("Parameter is invalid")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "parameter is invalid",
		})
		return
	}

	var uId int64
	var err error
	if uId, err = controller.Login(name, password); err != nil {
		logger.Warn("login error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "unknow error",
		})
		return
	}

	c.SetCookie("uid", strconv.FormatInt(uId, 10), 86400, "/", "148.70.137.99:8080", false, true)

	c.HTML(http.StatusOK, "personalpage.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"title" : name,
	})
}

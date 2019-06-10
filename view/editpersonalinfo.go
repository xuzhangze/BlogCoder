package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"github.com/xuzhangze/BlogCoder/middle"
	"net/http"
	"strconv"
)

func EditPersonalInfoHandle(c *gin.Context) {
	uIdStr, err := c.Cookie("uid")
	if err != nil {
		logger.Warn("User ID does not exist in request handler")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get cookie error",
		})
		return
	}

	page := c.DefaultQuery("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "editpersonalinfo.html", gin.H{
			"code" : 1,
			"msg" : "success",
		})
		return
	}

	password := c.DefaultQuery("password", "")
	if len(password)==0 {
		logger.Warn("Parameter is invalid")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "parameter is invalid",
		})
		return
	}

	uId, _ := strconv.ParseInt(uIdStr, 10, 64)
	var user middle.UserInfo
	if user, err = controller.UpdatePersonalInfo(uId, password); err != nil {
		logger.Error("Service error, err: %v", err)
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

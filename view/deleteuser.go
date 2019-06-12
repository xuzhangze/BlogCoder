package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"github.com/xuzhangze/BlogCoder/utils"
	"net/http"
	"strconv"
)

func DeleteUserHandle(c *gin.Context) {
	uidStr := c.DefaultQuery("uid", "0")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)

	_, err := controller.UpdatePersonalInfo(uid, utils.INITPASSWORD)
	if err != nil {
		logger.Error("Modify user info error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "unknow error",
		})
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"code" : 1,
		"msg" : "success",
	})
}

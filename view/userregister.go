package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
)

func UserRegisterHandle(c *gin.Context) {
	index := c.DefaultQuery("page", "")
	if len(index) != 0 {
		c.HTML(http.StatusOK, "register.html", gin.H{"code":1})
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

	if err := controller.UserRegister(name, password); err != nil {
		logger.Error("Call UserRegister error, name: %v, err: %v", name, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "unknow error",
		})
		return
	}

	c.HTML(http.StatusOK, "registersuccess.html", gin.H{
		"code" : 1,
		"msg" : "success",
	})
}
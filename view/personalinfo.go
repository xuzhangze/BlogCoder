package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func PersonalInfoHandle(c *gin.Context) {
	fmt.Println("request: ", *c.Request)
	id, err := c.Request.Cookie("uid")
	fmt.Println("id: ", id)
	fmt.Println("err: ", err)
	uIdStr, err := c.Cookie("uid")
	fmt.Println("uid: ", uIdStr)
	fmt.Println("err: ", err)
	if err != nil {
		logger.Warn("Get cookie error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "uid not exist",
		})
		return
	}

	uId, _ := strconv.ParseInt(uIdStr, 10, 64)
	user, err := controller.PersonalInfo(uId)
	if err != nil {
		logger.Error("service error: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "service error",
		})
	}

	c.HTML(http.StatusOK, "personalinfoedit.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"name" : user.GetUname(),
		"password" : user.GetPassword(),
	})
}

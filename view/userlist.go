package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"net/http"
	"strconv"
)

func UserListHandle(c *gin.Context) {
	uidStr, err := c.Cookie("uid")
	if err != nil {
		logger.Warn("Get cookie error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get cookie error",
		})
		return
	}

	page := c.DefaultQuery("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "userlist.html", gin.H{
			"code" : 1,
			"msg" : "success",
		})
		return
	}

	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	if uid != 4422517098958290944 {
		logger.Error("This uid is invalid")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "user id error",
		})
		return
	}

	users, err := controller.GetAllUsers()
	if err != nil {
		logger.Error("Get all user info error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get user info error",
		})
		return
	}

	datas := []map[string]string{}
	for _, user := range users {
		data := map[string]string{}
		data["uid"] = strconv.FormatInt(user.GetUid(), 10)
		data["uname"] = user.GetUname()
		data["password"] = user.GetPassword()

		datas = append(datas, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg" : "success",
		"data" : datas,
	})
}

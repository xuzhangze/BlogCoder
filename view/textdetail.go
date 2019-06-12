package view

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"github.com/xuzhangze/BlogCoder/middle"
	"net/http"
	"strconv"
)

func TextDetailHandle(c *gin.Context) {
	uidStr, _ := c.Cookie("uid")
	idStr := c.DefaultQuery("id", "")
	if len(idStr) == 0 {
		logger.Error("Invalid param")
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 1,
			"msg" : "invalid param",
		})
		return
	}
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	user, err := controller.PersonalInfo(uid)
	if err != nil {
		logger.Error("Get user info error, user ID: %v, err: %v", uid, err)
		user = middle.UserInfo{}
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := controller.GetBlogInfoByID(id)
	if err != nil {
		logger.Error("Get blog info error, blog ID: %v, err: %v", id, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 1,
			"msg" : "get blog info error",
		})
		return
	}

	extraJs := blog.GetExtra()
	extra := map[string]interface{}{}
	if err = json.Unmarshal([]byte(extraJs), &extra); err != nil {
		logger.Error("Extra json unmarshal error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 1,
			"msg" : "unknoe error",
		})
		return
	}

	flag := float64(1)
	publish := "no"
	if extra["publish"] == flag {
		publish = "yes"
	}

	c.HTML(http.StatusOK, "detailpage.html", gin.H{
		"code" : 1,
		"msg" : "success",
		"title" : blog.GetTitle(),
		"author" : user.GetUname(),
		"publish" : publish,
		"id" : blog.GetId(),
		"text" : blog.GetText(),
		"star" : blog.GetStar(),
		"browse" : blog.GetBrowse(),
	})
}

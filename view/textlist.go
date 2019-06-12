package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"github.com/xuzhangze/BlogCoder/utils"
	"net/http"
	"strconv"
)

func TextListHandle(c *gin.Context) {
	uidStr, _ := c.Cookie("uid")
	if len(uidStr) < 15 {
		logger.Error("Invalid user ID, user ID: %v", uidStr)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "invalid user id",
		})
		return
	}

	page := c.DefaultQuery("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "textlist.html", gin.H{
			"code" : 1,
			"msg" : "success",
		})
		return
	}

	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	blogs, err := controller.GetBlogsInfo(uid)
	if err != nil {
		logger.Error("Get blogs error, user ID: %v, err: %v", uid, err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 1,
			"msg" : "get blogs error",
		})
		return
	}

	datas := []map[string]string{}
	for _, blog := range blogs {
		data := map[string]string{}
		idStr := strconv.FormatInt(blog.GetId(), 10)
		data["id"] = idStr
		data["title"] = blog.GetTitle()
		text := blog.GetText()
		if len(text) > utils.TEXTLENGTH {
			text = text[:utils.TEXTLENGTH]
		}
		data["text"] = text
		datas = append(datas, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg" : "success",
		"data" : datas,
	})
}

package view

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/controller"
	"github.com/xuzhangze/BlogCoder/utils"
	"net/http"
	"strconv"
)

func TextFeedsHandle(c *gin.Context) {
	page := c.DefaultQuery("page", "")
	if len(page) != 0 {
		c.HTML(http.StatusOK, "feeds.html", gin.H{
			"code" : 1,
			"msg" : "success",
		})
		return
	}

	uidStr, err := c.Cookie("uid")
	if err != nil {
		uidStr = "0"
	}
	uid, _ := strconv.ParseInt(uidStr, 10, 64)

	blogs, err := controller.GetAllBlogsInfo(uid)
	if err != nil {
		logger.Error("Get all blog error, err: %v", err)
		c.HTML(http.StatusOK, "failed.html", gin.H{
			"code" : 0,
			"msg" : "get all blog error",
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

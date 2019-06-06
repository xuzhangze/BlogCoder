package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuzhangze/BlogCoder/view"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*/*/*")
	r.GET("/xblog/hello", view.HelloHandle)
	r.GET("/xblog/admin/v1/user/register_page", view.UserRegisterPage)
	r.GET("/xblog/admin/v1/user/register", view.UserRegisterHandle)
	r.GET("/xblog/admin/v1/user/login", view.LoginHandle)
	r.GET("/xblog/admin/v1/user/personal/info", view.PersonalInfoHandle)
	r.GET("/xblog/admin/v1/user/edit/personalinfo", view.EditPersonalInfoHandle)
	r.POST("/xblog/adminv1/text/edit", view.TextEdit)

	r.Run()   //http://148.70.137.99:8080/xblog/admin/v1/user/login?page=0
}

package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterPage(c *gin.Context) {
	//s, _ := exec.LookPath(os.Args[0])
	//fmt.Println(s)
//
	//c.JSON(http.StatusOK, gin.H{
	//	"code" : 0,
	//	"msg" : "success",
	//	"data" : s,
	//})
	c.HTML(http.StatusOK, "register.html", gin.H{"code":0})
}

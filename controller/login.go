package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"github.com/xuzhangze/BlogCoder/utils"
)

func Login(name, password string) (int64, error) {
	userinfo := middle.UserInfo{
		Uname: &name,
		Password: &password,
	}

	user, err := userinfo.Select()
	if err != nil {
		logger.Error("Query userinfo error, err: %v", err)
		return 0, err
	}
	userPassword := user.GetPassword()
	if userPassword != password {
		logger.Warn("Input password error, input password: %v", password)
		return 0, utils.ERRORPASSWORD
	}

	return user.GetUid(), nil
}

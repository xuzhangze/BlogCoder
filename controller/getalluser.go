package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func GetAllUsers() ([]middle.UserInfo, error) {
	userinfo := middle.UserInfo{}
	users, err := userinfo.SelectAll()
	if err != nil {
		logger.Error("Select user table error, err: %v", err)
		return []middle.UserInfo{}, err
	}

	return users, nil
}

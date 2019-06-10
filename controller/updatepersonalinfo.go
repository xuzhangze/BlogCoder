package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func UpdatePersonalInfo(uId int64, password string) (middle.UserInfo, error) {
	usertext := middle.UserInfo{
		Uid:&uId,
		Password:&password,
	}

	if err := usertext.UpdateByUserID(); err != nil {
		logger.Error("Update user personal info error, user ID: %v, err: %v", uId, err)
		return middle.UserInfo{}, err
	}

	user, err := usertext.SelectByUserId()
	if err != nil {
		logger.Error("Query user info error, user ID: %v, err: %v", uId, err)
		return middle.UserInfo{}, err
	}

	return user, nil
}

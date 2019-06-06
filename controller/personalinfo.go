package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"github.com/xuzhangze/BlogCoder/utils"
)

func PersonalInfo(uId int64) (middle.UserInfo, error) {
	if uId <= 0 {
		return middle.UserInfo{}, utils.ERRORINVALIDUSERID
	}

	userinfo := middle.UserInfo{
		&uId,
		nil,
		nil,
		nil,
	}

	user, err := userinfo.SelectByUserId()
	if err != nil {
		logger.Error("This user ID is invalid")
		return middle.UserInfo{}, utils.ERRORINVALIDUSERID
	}

	return user, nil
}

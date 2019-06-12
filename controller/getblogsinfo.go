package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func GetBlogsInfo(uid int64) ([]middle.TextInfo, error) {
	textinfo := middle.TextInfo{
		Uid: &uid,
	}

	blogs, err := textinfo.SelectByUserID()
	if err != nil {
		logger.Error("get blogs error, err: %v", err)
		return []middle.TextInfo{}, err
	}

	return blogs, nil
}

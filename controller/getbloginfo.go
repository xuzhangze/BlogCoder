package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func GetBlogInfoByID(id int64) (middle.TextInfo, error) {
	var blog middle.TextInfo
	var err error
	textinfo := middle.TextInfo{
		Id: &id,
	}

	blog, err = textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return blog, err
	}

	return blog, nil
}

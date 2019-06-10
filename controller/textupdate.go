package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func Updatetext(id int64, title, text string) (middle.TextInfo, error) {
	textinfo := middle.TextInfo{
		Id: &id,
	}
	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	blog.Title = &title
	blog.Text = &text
	if err = blog.Update(); err != nil {
		logger.Error("Blog update error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	return blog, nil
}

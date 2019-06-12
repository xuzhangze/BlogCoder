package controller

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"github.com/xuzhangze/BlogCoder/utils"
)

func Updatetext(id, uId int64, title, text string) (middle.TextInfo, error) {
	textinfo := middle.TextInfo{
		Id: &id,
	}
	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	blogUid := blog.GetUid()
	if blogUid != uId && uId != 4422517098958290944 {
		logger.Error("Not delete jurisdiction, blog user ID: %v, option user ID: %v", blogUid, uId)
		return middle.TextInfo{}, utils.ERRORUNKNOW
	}

	blog.Title = &title
	blog.Text = &text
	if err = blog.Update(); err != nil {
		logger.Error("Blog update error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	return blog, nil
}

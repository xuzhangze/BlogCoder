package controller

import (
	"encoding/json"
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

	extraStr := blog.GetExtra()
	extra := map[string]interface{}{}
	if err := json.Unmarshal([]byte(extraStr), &extra); err != nil {
		logger.Error("Json unmarshal error, err: %v", err)
		return middle.TextInfo{}, err
	}
	extra["publish"] = 0

	extraJs, err := json.Marshal(extra)
	if err != nil {
		logger.Error("Json marshal error, err: %v", err)
		return middle.TextInfo{}, err
	}
	extraStr = string(extraJs)

	blog.Title = &title
	blog.Text = &text
	blog.Extra = &extraStr
	if err = blog.Update(); err != nil {
		logger.Error("Blog update error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	return blog, nil
}

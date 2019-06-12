package controller

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"github.com/xuzhangze/BlogCoder/utils"
)

func TextPublish(id, uid int64, publish int32) (middle.TextInfo, error) {
	textinfo := middle.TextInfo{
		Id: &id,
	}

	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	blogUid := blog.GetUid()
	if blogUid != uid && uid != 4422517098958290944 {
		logger.Error("Not delete jurisdiction, blog user ID: %v, option user ID: %v", blogUid, uid)
		return middle.TextInfo{}, utils.ERRORUNKNOW
	}

	extraStr := blog.GetExtra()
	extra := map[string]interface{}{}
	err = json.Unmarshal([]byte(extraStr), &extra)
	if err != nil {
		logger.Error("Blog's extra unmarshal error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	extra["publish"] = publish
	extraJs, err := json.Marshal(extra)
	if err != nil {
		logger.Error("Blog' extra marshal error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	extraStr = string(extraJs)
	blog.Extra = &extraStr
	if err = blog.Update(); err != nil {
		logger.Error("Blog update error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	return blog, nil
}

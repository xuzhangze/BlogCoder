package controller

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func TextPublish(id int64, publish int32) (middle.TextInfo, error) {
	textinfo := middle.TextInfo{
		Id: &id,
	}

	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
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

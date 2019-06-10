package controller

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func TextDelete(id int64) error {
	textinfo := middle.TextInfo{
		Id: &id,
	}
	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return err
	}

	extra := map[string]interface{}{}
	if err = json.Unmarshal([]byte(blog.GetExtra()), &extra); err != nil {
		logger.Error("Blog json unmarshal error, blog ID: %v, err: %v", id, err)
		return err
	}

	newStatus := int32(0)
	extra["status"] = &newStatus

	extraJs, err := json.Marshal(extra)
	if err != nil {
		logger.Error("Blog' extra json marshal error, blog ID: %v, err: %v", id, err)
		return err
	}
	extraStr := string(extraJs)
	blog.Extra = &extraStr
	if err = blog.Update(); err != nil {
		logger.Error("Blog update error, blog ID: %v, err: %v", id, err)
		return err
	}

	return nil
}

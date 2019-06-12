package controller

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"github.com/xuzhangze/BlogCoder/utils"
)

func TextDelete(id, uid int64) error {
	textinfo := middle.TextInfo{
		Id: &id,
	}
	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Blog select error, blog ID: %v, err: %v", id, err)
		return err
	}

	blogUid := blog.GetUid()
	if blogUid != uid && uid != 4422517098958290944 {
		logger.Error("Not delete jurisdiction, blog user ID: %v, option user ID: %v", blogUid, uid)
		return utils.ERRORUNKNOW
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

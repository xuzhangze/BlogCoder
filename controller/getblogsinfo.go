package controller

import (
	"encoding/json"
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

	resp := []middle.TextInfo{}
	flag := float64(1)
	for _, blog := range blogs {
		extraStr := blog.GetExtra()
		extra := map[string]interface{}{}
		if err := json.Unmarshal([]byte(extraStr), &extra); err != nil{
			logger.Warn("Json unmarshal error, err: %v", err)
			continue
		}
		if status, ok := extra["status"]; ok {
			if status == flag {
				resp = append(resp, blog)
			}
		}
	}

	return resp, nil
}

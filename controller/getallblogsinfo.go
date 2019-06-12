package controller

import (
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"reflect"
)

func GetAllBlogsInfo(uid int64) ([]middle.TextInfo, error) {
	textinfo := middle.TextInfo{}
	blogs, err := textinfo.SelectAll()
	if err != nil {
		logger.Error("Blog select error, err: %v", err)
		return []middle.TextInfo{}, err
	}

	if uid == 4422517098958290944 {
		return blogs, nil
	}

	flag := float64(1)
	var res []middle.TextInfo
	for _, blog := range blogs {
		extraJs := blog.GetExtra()
		extra := map[string]interface{}{}
		if err := json.Unmarshal([]byte(extraJs), &extra); err != nil {
			logger.Warn("Json unmarshal error, err: %v", err)
			continue
		}
		if status, ok := extra["status"]; ok {
			if status == flag {
				if publish, ok := extra["publish"]; ok {
					if publish == flag {
						res = append(res, blog)
					}
				}
			}
		}
	}

	return res, nil
}

func printtype(i interface{}) {
	fmt.Println("type:", reflect.TypeOf(i))
}

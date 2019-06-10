package controller

import (
	"encoding/json"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
)

func TextEdit(uId int64, title, text string) (middle.TextInfo, error) {
	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(100000000000000000, 0)
	id, err := currWorker.NextId()
	if err != nil {
		logger.Error("Get blog ID error, err: %v", err)
		return middle.TextInfo{}, err
	}

	m := map[string]int32{
		"status" : 1,
		"publish" : 0,
	}
	extraJs, err := json.Marshal(m)
	if err != nil {
		logger.Error("Json marshal error, err: %v", err)
		return middle.TextInfo{}, err
	}
	extra := string(extraJs)

	textinfo := middle.TextInfo{
		Id: &id,
		Uid: &uId,
		Title: &title,
		Text: &text,
		Extra: &extra,
	}

	if err := textinfo.Insert(); err != nil {
		logger.Error("Blog save error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	blog, err := textinfo.Select()
	if err != nil {
		logger.Error("Select blog error, blog ID: %v, err: %v", id, err)
		return middle.TextInfo{}, err
	}

	return blog, nil
}

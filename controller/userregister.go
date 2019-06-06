package controller

import (
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/middle"
	"github.com/xuzhangze/BlogCoder/utils"
)

func UserRegister(name, password string) error {
	user := &middle.UserInfo{}
	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(100000000000000000, 0)
	uid, err := currWorker.NextId()
	if err != nil {
		logger.Error("Get user id error, err: %v", err)
		return err
	}

	user.Uid = &uid
	user.Uname = &name
	user.Password = &password

	if _, err := user.Select(); err == nil {
		logger.Warn("This name already exist")
		return utils.ERRORALREADYEXIST
	}

	if err := user.Insert(); err != nil {
		logger.Error("User infomation save error, name: %v, err: %v", name, err)
		return err
	}

	return nil
}

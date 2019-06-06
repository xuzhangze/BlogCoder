package middle

import (
	"fmt"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/wonderivan/logger"
	"testing"
)

func _TestUserInfo_Insert(t *testing.T) {
	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(100000000000000000, 0)
	uId, err := currWorker.NextId()
	if err != nil {
		logger.Error("Get user id error, err: %v", err)
		return
	}
	uname := "wangwu"
	password := "3366"
	user := UserInfo{&uId, &uname, &password, nil}
	if err = user.Insert(); err != nil {
		t.Error("insert to db error")
	}
}

func _TestUserInfo_Update(t *testing.T) {
	name := "wangwu"
	password := "6677"
	user := UserInfo{
		Uname:&name,
		Password:&password,
	}
	if err := user.Update(); err != nil {
		t.Error("update db error")
	}
}

func TestUserInfo_Select(t *testing.T) {
	name := "zhaoliu"
	user := UserInfo{
		Uname:&name,
	}

	userinfo, err := user.Select()
	if err != nil {
		t.Error("select db usertable error")
		return
	}
	fmt.Println(userinfo.GetUid(), userinfo.GetUname(), userinfo.GetPassword(), userinfo.GetExtra())
}

func TestTextInfo_Insert(t *testing.T) {
	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(100000000000000000, 0)
	id, err := currWorker.NextId()
	if err != nil {
		logger.Error("Get user id error, err: %v", err)
		return
	}
	uid, err := currWorker.NextId()
	if err != nil {
		logger.Error("Get user id error, err: %v", err)
		return
	}
	title := "second blog"
	text := "this is second blog's text"
	star := 1
	browse := 1
	textinfo := &TextInfo{&id, &uid, &title, &text, &star, &browse, nil}
	if err = textinfo.Insert(); err != nil {
		t.Error("insert value into table textinfo")
	}
}

func TestTextInfo(t *testing.T) {
	textinfo := &TextInfo{}
	if textinfo.Id == nil {
		fmt.Println("id ok")
	}
	if textinfo.Uid == nil {
		fmt.Println("uid ok")
	}
	if textinfo.Title == nil {
		fmt.Println("title ok")
	}
	if textinfo.Text == nil {
		fmt.Println("text ok")
	}
	if textinfo.Star == nil {
		fmt.Println("star ok")
	}
	if textinfo.Browse == nil {
		fmt.Println("browse ok")
	}
	if textinfo.Extra == nil {
		fmt.Println("extra ok")
	}
}
package middle

import (
	"fmt"
	"github.com/wonderivan/logger"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	err := GetUserInfo()
	if err != nil {
		t.Error("call GetUserInfo error")
	}
}

func _TestInsertValueToUserInfoTable(t *testing.T) {
	err := InsertValueToUserInfoTable(1, "zhangsan", "123", "")
	if err != nil {
		t.Error("call InsertValueToUserInfoTable error")
	}
}

func _TestUpdateValueToUserInfoTable(t *testing.T) {
	err := UpdateValueToUserInfoTable("zhangsan", "789")
	if err != nil {
		t.Error("call UpdateValueToUserInfoTable error")
	}
}

func TestSelectOfUserinfoTable(t *testing.T) {
	users, err := SelectOfUserinfoTable("zhangsan")
	if err != nil {
		t.Error("call SelectOfUserinfoTable error")
	}
	user := UserInfo{}
	if len(users) == 0 {
		logger.Warn("Result is empty")
		return
	} else {
		user = users[0]
	}
	fmt.Println(user.GetUid(), user.GetUname(), user.GetPassword(), user.GetExtra())
	
	users, err = SelectOfUserinfoTable("lisi")
	if err != nil {
		t.Error("call SelectOfUserinfoTable error")
	}
	if len(users) == 0 {
		logger.Warn("Result is empty")
		return
	} else {
		user = users[0]
	}
	fmt.Println(user.GetUid(), user.GetUname(), user.GetPassword(), user.GetExtra())
}
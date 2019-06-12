package middle

import (
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/utils"
)

type UserInfo struct {
	Uid *int64 `json:"uid"`
	Uname *string `json:"uname"`
	Password *string `json:"password"`
	Extra *string `json:"extra"`
}

func (userinfo *UserInfo)GetUid() int64 {
	if userinfo.Uid == nil {
		return 0;
	}

	return *userinfo.Uid
}

func (userinfo *UserInfo)GetUname() string {
	if userinfo == nil {
		return ""
	}

	return *userinfo.Uname
}

func (userinfo *UserInfo)GetPassword() string {
	if userinfo.Password == nil {
		return ""
	}

	return *userinfo.Password
}

func (userinfo *UserInfo)GetExtra() string {
	if userinfo.Extra == nil {
		return ""
	}

	return *userinfo.Extra
}

func (userinfo *UserInfo)Insert() error {
	if err := InsertValueToUserInfoTable(userinfo.GetUid(), userinfo.GetUname(), userinfo.GetPassword(), userinfo.GetExtra()); err != nil {
		logger.Error("UserInfo insert to db error, err: %v", err)
		return err
	}

	return nil
}

func (userinfo *UserInfo)Update() error {
	if err := UpdateValueToUserInfoTable(userinfo.GetUname(), userinfo.GetPassword()); err != nil {
		logger.Error("Update userinfo error, err: %v", err)
		return err
	}

	return nil
}

func (userinfo *UserInfo)UpdateByUserID() error {
	if err := UpdateValueToUserInfoTableByUserID(userinfo.GetUid(), userinfo.GetPassword()); err != nil {
		logger.Error("Update userinfo error, err: %v", err)
		return err
	}

	return nil
}

func (userinfo *UserInfo)Select() (UserInfo, error) {
	users, err := SelectOfUserinfoTable(userinfo.GetUname())
	if err != nil {
		logger.Error("Query userinfo talbe error, err: %v", err)
		return UserInfo{}, err
	}
	if len(users) == 0 {
		logger.Warn("This element not found")
		return UserInfo{}, utils.ERRORNOTFOUND
	}

	return users[0], nil
}

func (userinfo *UserInfo)SelectAll() ([]UserInfo, error) {
	users, err := SelectOfUserinfoTableAll()
	if err != nil {
		logger.Error("Query userinfo talbe error, err: %v", err)
		return []UserInfo{}, err
	}
	if len(users) == 0 {
		logger.Warn("This element not found")
		return []UserInfo{}, utils.ERRORNOTFOUND
	}

	return users, nil
}

func (userinfo *UserInfo)SelectByUserId() (UserInfo, error) {
	users, err := SelectOfUserinfoTableByUserID(userinfo.GetUid())

	if err != nil {
		logger.Error("Query userinfo talbe error, err: %v", err)
		return UserInfo{}, err
	}
	if len(users) == 0 {
		logger.Warn("This element not found")
		return UserInfo{}, utils.ERRORNOTFOUND
	}

	return users[0], nil
}

type TextInfo struct {
	Id *int64 `json:"id"`
	Uid *int64 `json:"uid"`
	Title *string `json:"title"`
	Text *string `json:"text"`
	Star *int `json:"star"`
	Browse *int `json:"browse"`
	Extra *string `json:"extra"`
}

func (textinfo *TextInfo)GetId() int64 {
	if textinfo.Id == nil {
		return int64(0)
	}

	return *textinfo.Id
}

func (textinfo *TextInfo)GetUid() int64 {
	if textinfo.Uid == nil {
		return int64(0)
	}

	return *textinfo.Uid
}

func (textinfo *TextInfo)GetTitle() string {
	if textinfo.Title == nil {
		return ""
	}

	return *textinfo.Title
}

func (textinfo *TextInfo)GetText() string {
	if textinfo.Text == nil {
		return ""
	}

	return *textinfo.Text
}

func (textinfo *TextInfo)GetStar() int {
	if textinfo.Star == nil {
		return 0
	}

	return *textinfo.Star
}

func (textinfo *TextInfo)GetBrowse() int {
	if textinfo.Browse == nil {
		return 0
	}

	return *textinfo.Browse
}

func (textinfo *TextInfo)GetExtra() string {
	if textinfo.Extra == nil {
		return ""
	}

	return *textinfo.Extra
}

func (textinfo *TextInfo)Insert() error {
	if err := InsertValueToTextInfoTable(textinfo.GetId(), textinfo.GetUid(), textinfo.GetStar(), textinfo.GetBrowse(), textinfo.GetTitle(), textinfo.GetText(), textinfo.GetExtra()); err != nil {
		logger.Error("Insert value into table textinfo error, err: %v", err)
		return err
	}

	return nil
}

func (textinfo *TextInfo)Update() error {
	if textinfo.Id == nil || textinfo.Uid == nil {
		logger.Error("Struct is empty")
		return utils.ERRORINVALIDELEMENT
	}

	textinfos, err := SelectOfTextInfoTable(textinfo.GetId())
	if err != nil {
		logger.Error("Select table textinfo error, id: %v, err: %v", textinfo.GetId(), err)
		return err
	}
	if len(textinfos) == 0 {
		logger.Error("Id is invalid, id: %v", textinfo.GetId())
		return utils.ERRORINVALIDELEMENT
	}

	textInfo := textinfos[0]
	if textinfo.Title != nil {
		textInfo.Title = textinfo.Title
	}
	if textinfo.Text != nil {
		textInfo.Text = textinfo.Text
	}
	if textinfo.Star != nil {
		textInfo.Star = textinfo.Star
	}
	if textinfo.Browse != nil {
		textInfo.Browse = textinfo.Browse
	}
	if textinfo.Extra != nil {
		textInfo.Extra = textinfo.Extra
	}

	if err = UpdateValueToTextInfoTable(textInfo.GetId(), textInfo.GetStar(), textInfo.GetBrowse(), textInfo.GetTitle(), textInfo.GetText(), textInfo.GetExtra()); err != nil {
		logger.Error("Update table textinfo error, err: %v", err)
		return err
	}

	return nil
}

func (textinfo *TextInfo)Select() (TextInfo, error) {
	textinfos, err := SelectOfTextInfoTable(textinfo.GetId())
	if err != nil {
		logger.Error("Select table textinfo error, err: %v", err)
		return TextInfo{}, err
	}
	if len(textinfos) == 0 {
		logger.Warn("Query table result is empty, id: %v", textinfo.GetId())
		return TextInfo{}, utils.ERRORNOTFOUND
	}

	return textinfos[0], nil
}

func (textinfo *TextInfo)SelectByUserID() ([]TextInfo, error) {
	textinfos, err := SelectOfTextInfoTableByUserID(textinfo.GetUid())
	if err != nil {
		logger.Error("Select table textinfo error, err: %v", err)
		return []TextInfo{}, err
	}
	if len(textinfos) == 0 {
		logger.Warn("Query table result is empty, id: %v", textinfo.GetId())
		return []TextInfo{}, utils.ERRORNOTFOUND
	}

	return textinfos, nil
}

func (textinfo *TextInfo)SelectAll() ([]TextInfo, error) {
	textinfos, err := SelectOfTextInfoTableAll()
	if err != nil {
		logger.Error("Select table textinfo error, err: %v", err)
		return []TextInfo{}, err
	}
	if len(textinfos) == 0 {
		logger.Warn("Query table result is empty, id: %v", textinfo.GetId())
		return []TextInfo{}, utils.ERRORNOTFOUND
	}

	return textinfos, nil
}
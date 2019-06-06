package middle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"github.com/xuzhangze/BlogCoder/module"
)

var Blogdb *sql.DB = nil

func init() {
	if Blogdb == nil {
		var err error
		Blogdb, err = module.GetBlogDB()
		if err != nil {
			logger.Error("Connect mysql error, err: %v", err)
		}
	}
}

func checkBlogdb() error {
	if Blogdb == nil {
		var err error
		Blogdb, err = module.GetBlogDB()
		if err != nil {
			logger.Error("Connect DB error, err: %v", err)
			return err
		}
	}

	return nil
}

func GetUserInfo() error {
	db, err := module.GetUserDB()
	if err != nil {
		logger.Error("Connect mysql error, err: %v", err)
		return err
	}

	rows,err := db.Query("SELECT * FROM t1")
	if err != nil {
		logger.Error("Get query result error, err: %v", err)
		return err
	}

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			logger.Error("Traversal query result error, err: %v", err)
			return err
		}
		fmt.Println(id, ":", name)
	}

	return nil
}

func InsertValueToUserInfoTable(uId int64, name, password, extra string) error {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return err
	}

	//currWorker := &idworker.IdWorker{}
	//currWorker.InitIdWorker(100000000000000000, 0)
	//uId, err := currWorker.NextId()
	//if err != nil {
	//	logger.Error("Get user id error, err: %v", err)
	//	return err
	//}

	if len(extra) == 0 {
		m := map[string]string{name : password}
		mJson, err := json.Marshal(m)
		if err != nil {
			logger.Error("Json Marshal error, err: %v, value: %v", err, m)
			return err
		}
		extra = string(mJson)
	}

	stmt, err := Blogdb.Prepare("INSERT userinfo SET uid=?,uname=?,password=?,extra=?")
	if err != nil {
		logger.Error("Start insert to userinfo error, err: %v", err)
		return err
	}
	res, err := stmt.Exec(uId, name, password, extra)
	if err != nil {
		logger.Error("Give value to userinfo error, err: %v", err)
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		logger.Error("Insert into userinfo error, err: %v", err)
		return err
	}

	return nil
}

func UpdateValueToUserInfoTable(name, password string) error {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return err
	}

	stmt, err := Blogdb.Prepare("UPDATE userinfo SET password=? WHERE uname=?")
	if err != nil {
		logger.Error("Start update userinfo error, err: %v", err)
		return err
	}
	res, err := stmt.Exec(password, name)
	if err != nil {
		logger.Error("Give value to userinfo error, err: %v, uname: %v", err, name)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		logger.Error("Update value to userinfo error, err: %v", err)
		return err
	}

	return nil
}

func UpdateValueToUserInfoTableByUserID(uId int64, password string) error {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return err
	}

	stmt, err := Blogdb.Prepare("UPDATE userinfo SET password=? WHERE uname=?")
	if err != nil {
		logger.Error("Start update userinfo error, err: %v", err)
		return err
	}
	res, err := stmt.Exec(password, uId)
	if err != nil {
		logger.Error("Give value to userinfo error, err: %v, uname: %v", err, uId)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		logger.Error("Update value to userinfo error, err: %v", err)
		return err
	}

	return nil
}

func SelectOfUserinfoTable(name string) ([]UserInfo, error) {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return nil, err
	}

	rows,err := Blogdb.Query("SELECT * FROM userinfo WHERE uname=?", name)
	if err != nil {
		logger.Error("Get query result error, err: %v", err)
		return nil, err
	}

	result := []UserInfo{}
	for rows.Next() {
		var uid int64
		var uname, password, extra string
		user := UserInfo{&uid, &uname, &password, &extra}
		if err = rows.Scan(user.Uid, user.Uname, user.Password, user.Extra); err != nil {
			logger.Error("Traversal query result error, err: %v", err)
			return nil, err
		}
		result = append(result, user)
	}

	return result, nil
}

func SelectOfUserinfoTableByUserID(uId int64) ([]UserInfo, error) {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return nil, err
	}

	rows,err := Blogdb.Query("SELECT * FROM userinfo WHERE uid=?", uId)
	if err != nil {
		logger.Error("Get query result error, err: %v", err)
		return nil, err
	}

	result := []UserInfo{}
	for rows.Next() {
		var uid int64
		var uname, password, extra string
		user := UserInfo{&uid, &uname, &password, &extra}
		if err = rows.Scan(user.Uid, user.Uname, user.Password, user.Extra); err != nil {
			logger.Error("Traversal query result error, err: %v", err)
			return nil, err
		}
		result = append(result, user)
	}

	return result, nil
}

func InsertValueToTextInfoTable(id, uId int64, star, browse int, title, text, extra string) error {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return err
	}

	if len(extra) == 0 {
		m := map[string]int64{title : id}
		mJson, err := json.Marshal(m)
		if err != nil {
			logger.Error("Json Marshal error, err: %v, value: %v", err, m)
			return err
		}
		extra = string(mJson)
	}

	stmt, err := Blogdb.Prepare("INSERT textinfo SET id=?,uid=?,title=?,text=?,star=?,browse=?,extra=?")
	if err != nil {
		logger.Error("Start insert to textinfo error, err: %v", err)
		return err
	}
	res, err := stmt.Exec(id, uId, title, text, star, browse, extra)
	if err != nil {
		logger.Error("Give value to textinfo error, err: %v", err)
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		logger.Error("Insert into textinfo error, err: %v", err)
		return err
	}

	return nil
}

func UpdateValueToTextInfoTable(id int64, star, browse int, title, text, extra string) error {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return err
	}

	stmt, err := Blogdb.Prepare("UPDATE textinfo SET title=?,text=?,star=?,browse=?,extra=? WHERE id=?")
	if err != nil {
		logger.Error("Start update userinfo error, err: %v", err)
		return err
	}
	res, err := stmt.Exec(title, text, star, browse, extra, id)
	if err != nil {
		logger.Error("Give value to userinfo error, err: %v, title: %v", err, title)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		logger.Error("Update value to userinfo error, err: %v", err)
		return err
	}

	return nil
}

func SelectOfTextInfoTable(id int64) ([]TextInfo, error) {
	if err := checkBlogdb(); err != nil {
		logger.Error("DB connect error, err: %v", err)
		return nil, err
	}

	rows,err := Blogdb.Query("SELECT * FROM userinfo WHERE id=?", id)
	if err != nil {
		logger.Error("Get query result error, err: %v", err)
		return nil, err
	}

	result := []TextInfo{}
	for rows.Next() {
		var id, uid int64
		var star, browse int
		var title, text, extra string
		textinfo := TextInfo{&id, &uid, &title, &text, &star, &browse, &extra}
		if err = rows.Scan(textinfo.Id, textinfo.Uid, textinfo.Title, textinfo.Text, textinfo.Star, textinfo.Browse, textinfo.Extra); err != nil {
			logger.Error("Traversal query result error, err: %v", err)
			return nil, err
		}
		result = append(result, textinfo)
	}

	return result, nil
}
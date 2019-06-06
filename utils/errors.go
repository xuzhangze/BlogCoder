package utils

import "errors"

var (
	ERRORINVALIDELEMENT = errors.New("invalid element of struct")
	ERRORINVALIDPARAMETER = errors.New("parameter is invalid")
	ERRORUNKNOW = errors.New("unknow error")
	ERRORNOTFOUND = errors.New("element not found")
	ERRORALREADYEXIST = errors.New("already exist")
	ERRORPASSWORD = errors.New("password is failed")
	ERRORINVALIDUSERID = errors.New("user id is invalid")
)

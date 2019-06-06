package module

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:cjmxzz1314@/test_xu?charset=utf8")
}

func GetBlogDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:cjmxzz1314@/blogdb?charset=utf8")
}
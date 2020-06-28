package db

import (
	"database/sql"
	"douyu-point/common"
	"douyu-point/global"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(url string) {
	var err error
	global.DB, err = sql.Open("mysql", url)
	common.CheckErr(err)
	err = global.DB.Ping()
	common.CheckErr(err)
}

func InsertData(uid string, id string, change string) {
	_, err := global.DB.Exec("INSERT INTO points(uid, id, point) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE point=point+?,id=?", uid, id, change, change, id)
	common.CheckErr(err)
}
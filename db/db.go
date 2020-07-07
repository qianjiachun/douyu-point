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

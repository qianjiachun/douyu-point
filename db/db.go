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

func QueryByUid(uid string) []global.UserInfo {
	var items []global.UserInfo
	var item global.UserInfo
	row := global.DB.QueryRow("select * from points where uid = ?", uid)
	err := row.Scan(&item.Uid, &item.Id, &item.Point, &item.UpdateTime)
	common.CheckErr(err)
	items = append(items, item)
	return items
}

func QueryById(id string) []global.UserInfo {
	var items []global.UserInfo
	var item global.UserInfo
	rows, _ := global.DB.Query("select * from points where id like ?", "%" + id + "%")
	for rows.Next() {
		err := rows.Scan(&item.Uid, &item.Id, &item.Point, &item.UpdateTime)
		common.CheckErr(err)
		items = append(items, item)
	}
	return items
}
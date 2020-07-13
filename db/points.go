package db

import (
	"database/sql"
	"douyu-point/common"
	"douyu-point/global"
)

func InsertUserInfo(uid string, id string, change string) {
	_, err := global.DB.Exec("INSERT INTO points(uid, id, point) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE point=point+?,id=?", uid, id, change, change, id)
	common.CheckErrNoExit(err)
}

func QueryUserInfoByUid(uid string) []global.UserInfo {
	var items []global.UserInfo
	var item global.UserInfo
	row := global.DB.QueryRow("select * from points where uid = ?", uid)
	_ = row.Scan(&item.Uid, &item.Id, &item.Point, &item.UpdateTime)
	items = append(items, item)
	return items
}

func QueryUserInfoById(id string) []global.UserInfo {
	var items []global.UserInfo
	var item global.UserInfo
	rows, _ := global.DB.Query("select * from points where id like ? order by update_time desc", "%"+id+"%")
	for rows.Next() {
		_ = rows.Scan(&item.Uid, &item.Id, &item.Point, &item.UpdateTime)
		items = append(items, item)
	}
	return items
}

func UpdateUserPointByUid(tx *sql.Tx, uid string, point string) bool {
	// 事务函数，请在事务内使用
	_, err := tx.Exec("update points set point = ? where uid = ?", point, uid)
	return common.CheckErrRollback(err, tx)
}

func QueryUserInfoByUid_Work(tx *sql.Tx, uid string) (global.UserInfo, bool) {
	// 事务函数，请在事务内使用
	var item global.UserInfo
	row := tx.QueryRow("select * from points where uid = ?", uid)
	err := row.Scan(&item.Uid, &item.Id, &item.Point, &item.UpdateTime)

	return item, common.CheckErrRollback(err, tx)
}

func QueryPointList() []global.UserInfo {
	var items []global.UserInfo
	var item global.UserInfo
	rows, _ := global.DB.Query("select * from points order by point desc limit 20")
	for rows.Next() {
		_ = rows.Scan(&item.Uid, &item.Id, &item.Point, &item.UpdateTime)
		items = append(items, item)
	}
	return items
}

package db

import (
	"database/sql"
	"douyu-point/common"
	"douyu-point/global"
)

func QueryItemByPage(offset string) []global.Item {
	var items []global.Item
	var item global.Item
	rows, _ := global.DB.Query("select * from items order by update_time desc limit 10 offset ?", offset)
	for rows.Next() {
		_ = rows.Scan(&item.Id, &item.Name, &item.Description, &item.Pic, &item.Price, &item.Num, &item.UpdateTime)
		items = append(items, item)
	}
	return items
}
func QueryItem() []global.Item {
	var items []global.Item
	var item global.Item
	rows, _ := global.DB.Query("select * from items order by update_time desc")
	for rows.Next() {
		_ = rows.Scan(&item.Id, &item.Name, &item.Description, &item.Pic, &item.Price, &item.Num, &item.UpdateTime)
		items = append(items, item)
	}
	return items
}

func QueryItemById(tx *sql.Tx, id string) (global.Item, bool) {
	// 事务函数，请在事务内使用
	// 内部使用，所以返回单个item
	var item global.Item
	row := tx.QueryRow("select * from items where id = ?", id)
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Pic, &item.Price, &item.Num, &item.UpdateTime)
	return item, common.CheckErrRollback(err, tx)
}

func UpdateItemNumById(tx *sql.Tx, id string, num string) bool {
	// 事务函数，请在事务内使用
	_, err := tx.Exec("update items set num = ? where id = ?", num, id)
	return common.CheckErrRollback(err, tx)
}

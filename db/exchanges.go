package db

import (
	"database/sql"
	"douyu-point/common"
	"douyu-point/global"
)

func InsertExchange(tx *sql.Tx, uid, id, item_id, item_name, item_price, info string) bool {
	// 事务函数，请在事务内使用
	_, err := tx.Exec("INSERT INTO exchanges(uid, id, item_id, item_name, item_price, info) VALUES (?, ?, ?, ?, ?, ?)", uid, id, item_id, item_name, item_price, info)
	return common.CheckErrRollback(err, tx)
}

func QueryExchangeByUidAndOffset(uid string, offset string) []global.Exchange {
	var items []global.Exchange
	var item global.Exchange
	rows, _ := global.DB.Query("select * from exchanges where uid = ? order by update_time desc limit 10 offset ?", uid, offset)
	for rows.Next() {
		_ = rows.Scan(&item.Status, &item.Uid, &item.Id, &item.Item_id, &item.Item_name, &item.Item_price, &item.Info, &item.UpdateTime)
		items = append(items, item)
	}
	return items
}

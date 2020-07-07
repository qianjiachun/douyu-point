package db

import (
	"database/sql"
	"douyu-point/common"
)

func InsertExchange(tx *sql.Tx, uid, id, item_id, item_name, item_price, info string) bool {
	// 事务函数，请在事务内使用
	_, err := tx.Exec("INSERT INTO exchanges(uid, id, item_id, item_name, item_price, info) VALUES (?, ?, ?, ?, ?, ?)", uid, id, item_id, item_name, item_price, info)
	return common.CheckErrRollback(err, tx)
}

package apis

import (
	"douyu-point/apis/room"
	"douyu-point/common"
	"net/http"
)

func Init_apis() {
	// 这里是HTTP接口的程序入口
	// 在这里可以实现相应的API接口
	http.HandleFunc("/douyu/point/5189167/query_by_uid", room.Api_queryByUid)
	http.HandleFunc("/douyu/point/5189167/exchange", room.Api_exchange)
	http.HandleFunc("/douyu/point/5189167/query_item", room.Api_queryItem)
	http.HandleFunc("/douyu/point/5189167/query_exchange", room.Api_queryExchange)
	http.HandleFunc("/douyu/point/5189167/point_list", room.Api_pointList)
	err := http.ListenAndServe(":27999", nil)
	common.CheckErr(err)
}

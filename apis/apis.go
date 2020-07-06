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
	err := http.ListenAndServe(":27999", nil)
	common.CheckErr(err)
}


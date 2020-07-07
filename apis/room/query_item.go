package room

import (
	"douyu-point/apis/apis_common"
	"douyu-point/common"
	"douyu-point/db"
	"douyu-point/global"
	"encoding/json"
	"fmt"
	"net/http"
)

func Api_queryItem(writer http.ResponseWriter, request *http.Request) {
	// post数据: token offset
	// token 斗鱼token
	// offset 位置

	writer.Header().Set("Access-Control-Allow-Origin", "*")             // 跨域 "*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")             //返回数据格式是json

	var ret *global.ItemJson
	ret = new(global.ItemJson)

	dyToken := request.PostFormValue("token")
	offset := request.PostFormValue("offset")
	isValid := apis_common.VerifyDyToken(dyToken)

	if isValid {
		ret.Error = 0
		ret.Msg = "success"
		ret.Data = db.QueryItemByPage(offset)
	} else {
		ret.Error = 1
		ret.Msg = "invalid token"
		ret.Data = nil
	}

	bytes, err := json.Marshal(ret)
	common.CheckErrNoExit(err)

	_, _ = fmt.Fprint(writer, string(bytes))
}

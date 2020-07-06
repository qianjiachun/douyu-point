package room

import (
	"douyu-point/apis/common"
	common2 "douyu-point/common"
	"douyu-point/db"
	"douyu-point/global"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Api_queryByUid(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")             // 跨域 "*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")             //返回数据格式是json

	var ret *global.UserInfoJson
	var uid string

	ret = new(global.UserInfoJson)

	dyToken := request.PostFormValue("token")
	isValid := common.VerifyDyToken(dyToken)

	if isValid {
		ret.Error = 0
		ret.Msg = "success"
		tmpArr := strings.Split(dyToken, "_")
		uid = tmpArr[0]
		ret.Data = db.QueryByUid(uid)
	} else {
		ret.Error = 1
		ret.Msg = "invalid token"
		ret.Data = nil
	}

	bytes, err := json.Marshal(ret)
	common2.CheckErr(err)

	_, _ = fmt.Fprint(writer, string(bytes))
}

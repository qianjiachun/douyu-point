package common

import (
	"douyu-point/common"
)

func VerifyDyToken(token string) bool {
	// 用于判断斗鱼的token是否有效
	var ret bool
	content := common.HttpPost("https://pcapi.douyucdn.cn/japi/tasksys/ytxb/box", "token=" + token)
	if common.GetStrMiddle(content, `"error":`, `,`) == "0" {
		ret = true
	} else {
		ret = false
	}
	return ret
}

package room

import (
	"douyu-point/apis/apis_common"
	"douyu-point/common"
	"douyu-point/db"
	"douyu-point/global"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
得到token，id，item_id，info => 检测token是否有效 => 有效则获取UID
=> 开启mysql事务
=> 根据item_id到items表里去获得物品的信息
=> 检查物品num是否足够 => 不足够的话就返回error:物品数量不足
=> 物品数量够，根据uid到points表里获取该用户的point => 检查积分够不够 => 不够的话就返回error:用户积分不足
=> 如果足够，说明积分够了，物品数量也充足
=> 扣除points表中用户积分(update) => 扣除items表中物品数量(num) => 插入到exchanges表中去
=> 提交mysql事务 => 返回success
*/
func Api_exchange(writer http.ResponseWriter, request *http.Request) {
	// post数据:token, item_id, id, info
	// token 斗鱼token
	// item_id 物品的id
	// id 用户名
	// info 兑换备注
	var ret *global.ItemExchangeJson
	ret = new(global.ItemExchangeJson)

	writer.Header().Set("Access-Control-Allow-Origin", "*")             // 跨域 "*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")             //返回数据格式是json

	var uid string
	dyToken := request.PostFormValue("token")
	item_id := request.PostFormValue("item_id")
	id := request.PostFormValue("id")
	info := request.PostFormValue("info")

	isValid := apis_common.VerifyDyToken(dyToken)
	if isValid {
		tmpArr := strings.Split(dyToken, "_")
		uid = tmpArr[0]
	} else {
		returnJson(&writer, ret, 1, "invalid token", nil)
		return
	}

	var newPoint string
	var newNum string

	tx, err := global.DB.Begin()
	common.CheckErrNoExit(err)

	itemInfo, itemInfoRet := db.QueryItemById(tx, item_id)
	if itemInfoRet == false {
		returnJson(&writer, ret, 1, "无效的物品id", nil)
		return
	}

	if itemInfo.Num <= 0 {
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback() Error:" + err.Error())
			return
		}
		returnJson(&writer, ret, 1, "物品数量不足", nil)
		return
	}
	// 此时num足够
	newNum = strconv.Itoa(itemInfo.Num - 1)

	userInfo, userInfoRet := db.QueryUserInfoByUid_Work(tx, uid)
	if userInfoRet == false {
		return
	}
	if userInfo.Id != id {
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback() Error:" + err.Error())
			return
		}
		returnJson(&writer, ret, 1, "用户uid与id不匹配", nil)
		return
	}

	if userInfo.Point-itemInfo.Price < 0 {
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback() Error:" + err.Error())
			return
		}
		returnJson(&writer, ret, 1, "用户积分不足", nil)
		return
	}
	// 此时用户积分足够
	newPoint = strconv.Itoa(userInfo.Point - itemInfo.Price)

	if db.UpdateItemNumById(tx, item_id, newNum) == false {
		returnJson(&writer, ret, 1, "更新物品数量失败", nil)
		return
	}
	if db.UpdateUserPointByUid(tx, uid, newPoint) == false {
		returnJson(&writer, ret, 1, "更新用户积分失败", nil)
		return
	}
	item_price := strconv.Itoa(itemInfo.Price)
	if db.InsertExchange(tx, uid, id, item_id, itemInfo.Name, item_price, info) == false {
		returnJson(&writer, ret, 1, "更新用户兑换记录失败", nil)
		return
	}

	err = tx.Commit()
	if common.CheckErrRollback(err, tx) == false {
		returnJson(&writer, ret, 1, "兑换失败", nil)
		return
	}

	userInfo.Point = userInfo.Point - itemInfo.Price
	userInfo.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	var userInfos []global.UserInfo
	userInfos = append(userInfos, userInfo)
	returnJson(&writer, ret, 0, "兑换成功", userInfos)
}

func returnJson(writer *http.ResponseWriter, ret *global.ItemExchangeJson, error int, msg string, data []global.UserInfo) {
	ret.Error = error
	ret.Msg = msg
	ret.Data = data
	bytes, err := json.Marshal(ret)
	common.CheckErrNoExit(err)

	_, _ = fmt.Fprint(*writer, string(bytes))
}

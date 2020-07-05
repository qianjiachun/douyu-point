package msg_main

import (
	"douyu-point/common"
	"douyu-point/db"
	"douyu-point/global"
	"strconv"
	"strings"
	"time"
)

func Init_msg_main(data string) {
	var msgType string

	msgType = common.GetFieldValue(data, "type")
	for i := 0; i < len(global.Rules.Data); i++ {
		if global.Rules.Data[i].Type == msgType && global.Rules.Data[i].Type != "" {
			// 匹配到对应的type字段
			handleFields(data, global.Rules.Data[i].Fields)
			break
		}
	}
}

func handleFields(data string, fields []global.Field) {
	for i := 0; i < len(fields); i++ {
		var fieldCnt string
		var fieldValue string // 字段值
		var cnt int           // 相乘值

		fieldValue = common.GetFieldValue(data, fields[i].Name)
		fieldCnt = fields[i].Cnt
		if fieldCnt == "" {
			cnt = 1
		} else {
			cnt, _ = strconv.Atoi(common.GetFieldValue(data, fieldCnt))
		}

		handleRules(data, fieldValue, cnt, fields[i].Rules, fields[i].Default)

	}
}

func handleRules(data string, fieldValue string, cnt int, fieldRules []global.Rule, fieldDeafult global.Default) {
	matchedNum := 0 // 满足条件的个数。当所有rule都不满足，则执行default的rule
	for i := 0; i < len(fieldRules); i++ {
		// 先判断有没有开播
		item := fieldRules[i]
		if item.IsLive == true && global.IsLive == false {
			continue
		}

		// 判断是不是模糊匹配
		var isMatch bool
		if item.Fuzzy {
			// 模糊匹配
			if strings.Index(fieldValue, item.Value) != -1 {
				isMatch = true
			} else {
				isMatch = false
			}
		} else {
			// 完全匹配
			if fieldValue == item.Value {
				isMatch = true
			} else {
				isMatch = false
			}
		}

		if isMatch {
			matchedNum++
			// 符合rule，开始执行积分操作
			var tempChange string // 变化值
			var isExist bool
			tempType := common.GetFieldValue(data, "type") // type名
			tempUid := common.GetFieldValue(data, "uid")   // 用户uid
			tempId := common.GetFieldValue(data, "nn")     // 用户id

			ruleName := tempType + "_" + item.Value
			var limitNum int
			var nextTime int64

			isExist = false
			limitNum = 0
			nextTime = 0
			if _, ok := global.List[ruleName]; ok {
				if _, ok := global.List[ruleName][tempUid]; ok {
					isExist = true
					limitNum = global.List[ruleName][tempUid].Count // 已变化的次数
					nextTime = global.List[ruleName][tempUid].NextTime
				}
			} else {
				global.List[ruleName] = make(map[string]*global.InfoUid)
			}
			if time.Now().Unix() < nextTime {
				return
			}

			if isExist {
				global.List[ruleName][tempUid].NextTime = int64(item.Cd) + time.Now().Unix()
			} else {
				global.List[ruleName][tempUid] = &global.InfoUid{Count: 0, NextTime: int64(fieldDeafult.Cd) + time.Now().Unix()}
			}
			if item.Limit > 0 {
				// 需要限制
				// 格式：list[ruleName][uid] = num
				//limitNum := global.List[ruleName][tempUid].Count // 已变化的次数
				if limitNum < item.Limit {
					// 满足限制条件
					var newCnt int
					if limitNum+cnt <= item.Limit {
						newCnt = cnt
					} else {
						newCnt = item.Limit - limitNum
					}
					global.List[ruleName][tempUid].Count = limitNum + newCnt
					tempChange = strconv.Itoa(newCnt * item.Change)
					// 插入到数据库
					db.InsertData(tempUid, tempId, tempChange)
				}
			} else {
				// 无限制
				tempChange = strconv.Itoa(cnt * item.Change)
				// 插入到数据库
				db.InsertData(tempUid, tempId, tempChange)
			}

		}

	}

	if matchedNum == 0 {
		// 这里执行default规则
		// 判断是否启用
		if fieldDeafult.Enable == false {
			return
		}
		// 先判断开没开播
		if fieldDeafult.IsLive == true && global.IsLive == false {
			return
		}

		var tempChange string // 变化值
		var isExist bool
		tempType := common.GetFieldValue(data, "type") // type名
		tempUid := common.GetFieldValue(data, "uid")   // 用户uid
		tempId := common.GetFieldValue(data, "nn")     // 用户id
		ruleName := tempType + "_" + "default"
		var limitNum int
		var nextTime int64

		isExist = false
		limitNum = 0
		nextTime = 0
		if _, ok := global.List[ruleName]; ok {
			if _, ok := global.List[ruleName][tempUid]; ok {
				isExist = true
				limitNum = global.List[ruleName][tempUid].Count // 已变化的次数
				nextTime = global.List[ruleName][tempUid].NextTime
			}
		} else {
			global.List[ruleName] = make(map[string]*global.InfoUid)
		}

		if time.Now().Unix() < nextTime {
			return
		}

		if isExist {
			global.List[ruleName][tempUid].NextTime = int64(fieldDeafult.Cd) + time.Now().Unix()
		} else {
			global.List[ruleName][tempUid] = &global.InfoUid{Count: 0, NextTime: int64(fieldDeafult.Cd) + time.Now().Unix()}
		}
		if fieldDeafult.Limit > 0 {

			if limitNum < fieldDeafult.Limit {
				// 满足限制条件

				var newCnt int
				if limitNum+cnt <= fieldDeafult.Limit {
					newCnt = cnt
				} else {
					newCnt = fieldDeafult.Limit - limitNum
				}
				global.List[ruleName][tempUid].Count = limitNum + newCnt
				tempChange = strconv.Itoa(newCnt * fieldDeafult.Change)
				// 插入到数据库
				db.InsertData(tempUid, tempId, tempChange)
			}
		} else {
			// 无限制
			tempChange = strconv.Itoa(cnt * fieldDeafult.Change)
			// 插入到数据库
			db.InsertData(tempUid, tempId, tempChange)
		}

	}
}

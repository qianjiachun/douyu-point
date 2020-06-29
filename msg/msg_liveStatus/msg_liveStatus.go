package msg_liveStatus

import (
	"douyu-point/common"
	"douyu-point/global"
	"fmt"
	"time"
)

func Init_msg_liveStatus(data string) {
	if common.GetFieldValue(data, "type") == "rss" {
		t := time.Now()
		if common.GetFieldValue(data, "ss") == "1" {
			global.IsLive = true
			fmt.Println(t.Format("2006-01-02 15:04:05") + " : 房间已开播")
		} else {
			global.IsLive = false
			fmt.Println(t.Format("2006-01-02 15:04:05") + " : 房间已下播")
		}
	}
}

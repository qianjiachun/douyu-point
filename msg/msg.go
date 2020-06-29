package msg

import (
	"douyu-point/msg/msg_liveStatus"
	"douyu-point/msg/msg_main"
)

func Init_msg(data string) {
	// 这个包用于处理斗鱼服务端返回的信息
	// 这个函数是入口
	msg_liveStatus.Init_msg_liveStatus(data)
	msg_main.Init_msg_main(data)
}

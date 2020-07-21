package main

import (
	"douyu-point/apis"
	"douyu-point/client"
	"douyu-point/common"
	"douyu-point/db"
	"douyu-point/global"
	"douyu-point/msg"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/yudeguang/ratelimit"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	global.List = make(map[string]map[string]*global.InfoUid)
	loadConfig()
	loadRules()
	fmt.Println("======配置项载入完毕======")
	fmt.Println("直播间号:" + global.Config.RoomID)
	fmt.Println("获取开播状态中...")
	ret := common.GetStrMiddle(common.HttpGet("https://www.douyu.com/swf_api/h5room/"+global.Config.RoomID), `show_status":"`, `",`)
	fmt.Println("开播状态：" + ret)
	if ret == "1" {
		global.IsLive = true
	} else {
		global.IsLive = false
	}

	db.Connect(global.Config.MysqlURL)
	fmt.Println("=> 连接数据库完毕")
	dyConn := client.DouyuClient{Rid: global.Config.RoomID}
	dyConn.Connect(func(data string) {
		msg.Init_msg(data)
	})
	fmt.Println("=> 连接斗鱼服务器完毕")

	go cmdPanel()

	c := cron.New()
	err := c.AddFunc("0 0 0 * * ?", func() { // 每天0点重置limit
		t := time.Now()
		global.List = make(map[string]map[string]*global.InfoUid)
		fmt.Println(t.Format("2006-01-02 15:04:05") + " : limit重置完毕")
	})
	common.CheckErr(err)
	c.Start()

	fmt.Println("=> 服务启动成功")
	setRateLimit()
	apis.Init_apis()

}

func loadConfig() {
	f, err := ioutil.ReadFile("config.json")
	common.CheckErr(err)
	err = json.Unmarshal(f, &global.Config)
	common.CheckErr(err)

}
func loadRules() {
	f, err := ioutil.ReadFile(global.Config.Rules)
	common.CheckErr(err)
	err = json.Unmarshal(f, &global.Rules)
}

func cmdPanel() {
	var cmd string
	for {
		_, _ = fmt.Scanf("%s\n", &cmd)
		if cmd == "" {
			continue
		}
		if cmd == "reload" {
			cmd = ""
			loadRules()
			fmt.Println("rules重载完毕")

		} else if cmd == "help" {
			cmd = ""
			fmt.Println("reload : 重新载入rules文件")
			fmt.Println("resetLimit : 重置所有limit次数")
			fmt.Println("saveLimit : 保存limit状态")
			fmt.Println("loadLimit : 恢复limit状态")
			fmt.Println("logLimit : 输出limit数据")
			fmt.Println("exit : 退出程序")
		} else if cmd == "resetLimit" {
			cmd = ""
			global.List = make(map[string]map[string]*global.InfoUid)
			fmt.Println("limit重置完毕")
		} else if cmd == "saveLimit" {
			cmd = ""
			data, err := json.Marshal(global.List)
			common.CheckErrNoExit(err)
			f, err := os.OpenFile("./limit_bak", os.O_RDWR|os.O_CREATE, 0600)
			common.CheckErrNoExit(err)
			_, err = f.Write(data)
			common.CheckErrNoExit(err)
			_ = f.Close()
			println("limit保存成功")
		} else if cmd == "loadLimit" {
			cmd = ""
			if common.IsFileExist("./limit_bak") {
				bytes, err := ioutil.ReadFile("./limit_bak")
				common.CheckErrNoExit(err)
				err = json.Unmarshal(bytes, &global.List)
				common.CheckErrNoExit(err)
				println("limit恢复成功")
			} else {
				println("未找到limit_bak")
			}
		} else if cmd == "logLimit" {
			cmd = ""
			data, err := json.Marshal(global.List)
			common.CheckErr(err)
			println(string(data))
		} else if cmd == "exit" {
			cmd = ""
			os.Exit(0)
		} else {
			cmd = ""
			fmt.Println("无效的命令，请尝试输入help获取命令")
		}
	}
}
func setRateLimit() {
	global.RateLimit = ratelimit.NewRule()
	global.RateLimit.AddRule(time.Second*10, 5)
	global.RateLimit.AddRule(time.Hour*24, 500)
}

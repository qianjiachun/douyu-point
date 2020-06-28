## 斗鱼直播间积分系统
为你的直播间增加积分功能吧！

### 特性
1. 长久地在低端的服务器上运行
2. 热更新规则
3. 配置灵活

### 注意事项
1. 软件重启后limit会重置，也就是说每天限制的次数会重置
2. 软件每天0点自动重置limit
3. 修改rules后，请输入reload然后回车，此步操作可热更新积分规则
4. 如果遇到特殊情况可以输入resetLimit手动重置每天限制的次数
5. 本项目只完成了原型，更多的功能和HTTP接口请自行实现

### 使用步骤
1. 需要准备一台服务器，并且部署MySQL数据库
2. 按照下方要求手动设置数据库
3. 将本软件上传至服务器，并在同一目录下创建config.json文件，config配置项见下方
4. 再创建rules.json文件，放到config.json设置的路径下
5. rules.json的配置项见下方，为了方便配置可访问[json在线编辑](http://json.la/online.html)对rules进行配置
6. 全部配置好后，运行软件即可


### MySQL数据库
1. 创建一个table
```
points // 用户积分
```
2. points表（uid, id, point三个单列索引）
```
字段:
uid // 用户uid BIGINT(20) 主键 无默认值
id // 用户id VARCHAR(50)
point // 用户积分 INT(20)
update_time // 更新时间


创建代码:
CREATE TABLE `points` (
	`uid` BIGINT(20) NOT NULL DEFAULT '0',
	`id` VARCHAR(50) NULL DEFAULT NULL,
	`point` INT(20) NULL DEFAULT NULL,
	`update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`uid`),
	INDEX `index_id` (`id`),
	INDEX `index_point` (`point`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;


```

### config.json
```
// 配置属性一律使用小写+下划线形式
// 该文件为项目启动的配置文件
// 用于指定监听的直播间，数据库地址，规则地址
{
	room_id: "", // 房间号，必须是真实房间号
	mysql_url: "", // mysql数据库地址，格式必须为golang连接mysql的url格式
	rules: "./rules.json", // 规则地址
}
```

### rules.json
修改后，可输入reload重载规则  
``` json
// 用于锁定字段规则，字段来源斗鱼弹幕服务器第三方接入协议v1.6.2，字段名请与协议中的内容保持一致
// 设计目标：通过字段值寻址，最大化扩展性
{
	data: [
		{
			type: "chatmsg", // type字段，chatmsg表示弹幕，dgb表示礼物，uenter表示用户进入直播间
			fields: [
				{
					name: "txt", // 字段名，与接入协议保持一致，txt表示弹幕内容。
					cnt: "", // 数量字段，一般用于礼物信息，如果是礼物则填写"gfcnt"。填写""则默认为1，
					rules: [
						{
							value: "弹幕内容1", // 字段的值，注意一定是string类型，也就是带双引号（比如礼物id是纯数字也要用双引号）。这里表示type为chatmsg的里面的txt的值
                            fuzzy: false, // 是否模糊匹配
							change: 1, // 填写正负数，表示变化的的积分
							limit: 2, // 表示每日的次数上限，如果写1，表示一天只能计算一次，如果写-1则表示无限
                            isLive: true, // 是否在开播时才有效
						},
                        {
							value: "弹幕内容2",
                            fuzzy: false,
							change: 1,
							limit: 2,
                            isLive: true,
						},
					],
					default: { // 所有rules都不满足则执行default的规则
                        enable: false, // 用于开启default
						change: 10,
						limit: 1,
                        isLive: true,
					},
				}
			],
		}
	]
}

// 下面是模板，可依照模板定制个性化的积分规则
// 实现的规则如下：
// 1. 每天开播期间发送带有 #签到 的弹幕，增加1积分，每日限1次
// 2. 每天开播期间发送带有 自定义弹幕内容 的弹幕，增加1积分，每日限1次
// 3. 除上述两种弹幕，每天开播期间发送任意的弹幕，增加1积分，每日限1次
// 4. 赠送办卡(20002)，增加12积分，每日无限次
// 5. 赠送飞机(20003)，增加200积分，每日无限次
// 6. 每日开播期间进入直播间，增加1积分，每日限1次

{
    "data": [
        {
            "type": "chatmsg",
            "fields": [
                {
                    "name": "txt",
                    "cnt": "",
                    "rules": [
                        {
                            "value": "#签到",
                            "fuzzy": true,
                            "change": 1,
                            "limit": 1,
                            "isLive": true
                        },
                        {
                            "value": "自定义弹幕内容",
                            "fuzzy": true,
                            "change": 1,
                            "limit": 1,
                            "isLive": true
                        }
                    ],
                    "default": {
                        "enable": true,
                        "change": 1,
                        "limit": 1,
                        "isLive": true
                    }
                }
            ]
        },
        {
            "type": "dgb",
            "fields": [
                {
                    "name": "gfid",
                    "cnt": "gfcnt",
                    "rules": [
                        {
                            "value": "20002",
                            "fuzzy": false,
                            "change": 12,
                            "limit": -1,
                            "isLive": false
                        },
                        {
                            "value": "20003",
                            "fuzzy": false,
                            "change": 200,
                            "limit": -1,
                            "isLive": false
                        }
                    ],
                    "default": {
                        "enable": false,
                        "change": 0,
                        "limit": 0,
                        "isLive": true
                    }
                }
            ]
        },
        {
            "type": "uenter",
            "fields": [
                {
                    "name": "type",
                    "cnt": "",
                    "rules": [
                        {
                            "value": "uenter",
                            "fuzzy": false,
                            "change": 1,
                            "limit": 1,
                            "isLive": true
                        }
                    ],
                    "default": {
                        "enable": false,
                        "change": 0,
                        "limit": 0,
                        "isLive": true
                    }
                }
            ]
        }
    ]
}

```

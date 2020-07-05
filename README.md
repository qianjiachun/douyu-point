## 🚀 斗鱼直播间积分系统
为你的直播间增加积分功能吧！

## 特点
1. 能够长久地在低端的服务器上运行
2. 热更新规则
3. 配置灵活

## 声明
1. 本项目为个人兴趣开发，无任何盈利手段
2. 水平有限，代码质量不高，互联网同好可参考使用，**尊重作者劳动成果，引用请注明出处**
3. 若本项目有帮助到你，还请不要吝啬star❤
4. 作者: 小淳 / QQ: 189964430 / 邮箱: 189964430@qq.com

## 使用步骤
1. 需要准备一台服务器，并且部署MySQL数据库
2. 按照下方要求手动设置数据库
3. 将本软件上传至服务器，并在同一目录下创建config.json文件，config配置项见下方
4. 再创建rules.json文件，放到config.json设置的路径下
5. rules.json的配置项见下方，为了方便配置可访问[json在线编辑](http://json.la/online.html)对rules进行配置
6. 全部配置好后，运行软件即可

## 常见问题
### 如何快速部署服务?
- 不想配置规则，想直接运行看效果
1. 可以将项目中的dist文件夹上传到服务器，这个是项目的默认模板程序
2. 修改config.json中mysql_url
3. 使用下方的代码创建table
4. 运行douyu-point
```
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

### 如何修改积分规则(rules.json)?
1. 可按照模板/标准配置，手动或者访问[json在线编辑](http://json.la/online.html)对积分规则(rules.json)进行修改
2. 如果不是直接修改原rules文件，请修改完毕后，保存rules.json到config.json设置的路径下
3. 在程序面板输入reload，然后回车即可重载积分规则，且立即生效
- 请注意，修改积分规则不会重置次数(limit)

### 重启软件会有什么影响？
1. 首先并不建议重启软件，除非你要修改config.json
2. 重启软件会造成limit重置，也就是每日积分增加的次数限制都会重置
3. 软件目前没有实现服务自动重启功能。若遇到数据库崩溃、斗鱼服务器崩溃等不可抗因素导致软件退出运行，还请自行处理后续。


## 注意事项
1. 软件重启后limit会重置，也就是说每天限制的次数会重置
2. 软件每天0点自动重置limit
3. 修改rules后，请输入reload然后回车，此步操作可热更新积分规则
4. 如果遇到特殊情况可以输入resetLimit手动重置每天限制的次数
5. 本项目只完成了原型，更多的功能和HTTP接口请自行实现


## MySQL数据库
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

## config.json
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

## rules.json
修改后，可输入reload重载规则  
用于锁定字段规则，字段来源[斗鱼弹幕服务器第三方接入协议v1.6.2](https://wwa.lanzous.com/io5iMe4qojc)，字段名请与协议中的内容保持一致
```
// 设计目标：通过字段值寻址，最大化扩展性
{
    "data": [
        {
            "type": "chatmsg", // type字段，chatmsg表示弹幕，dgb表示礼物，uenter表示用户进入直播间
            "fields": [
                {
                    "name": "txt", // 字段名，与接入协议保持一致，txt表示弹幕内容。
                    "cnt": "", // 数量字段，一般用于礼物信息，如果是礼物则填写"gfcnt"。其他时候填写""即可
                    "rules": [
                        {
                            "value": "弹幕内容1", // 字段的值，注意一定是string类型，也就是带双引号（比如礼物id是纯数字也要用双引号）。这里表示type为chatmsg的里面的txt的值
                            "fuzzy": false, // 是否模糊匹配，若开启，只要弹幕内容包含value的值，就算符合
                            "change": 1, // 填写正负数，表示变化的的积分
                            "limit": 2, // 表示每日的次数上限，如果写1，表示一天只能计算一次，如果写0或-1则表示无限
                            "isLive": true, // 是否在开播时才有效
                            "cd": 3600 // 冷却时间，单位：秒。例如设置每个小时内只允许加分一次，就写3600。如果设置0，则没有冷却时间。
                        },
                        {
                            "value": "弹幕内容2",
                            "fuzzy": false,
                            "change": 1,
                            "limit": 2,
                            "isLive": true,
                            "cd": 0
                        }
                    ],
                    "default": { // 当上面所有rules都不满足时，则执行default的规则
                        "enable": false, // 是否开启default规则
                        "change": 10,
                        "limit": 1,
                        "isLive": true,
                        "cd": 0
                    }
                }
            ]
        }
    ]
}
```
## rules模板
``` json
// 下面是模板，可依照模板定制个性化的积分规则
// 实现的规则如下：
// 1. 每天开播期间发送带有 #签到 的弹幕，增加1积分，每日限1次
// 2. 每天发送带有 #打卡 的弹幕，增加1积分，每小时只能发送1次，每日限5次
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
              "isLive": true,
              "cd": 0
            },
            {
              "value": "#打卡",
              "fuzzy": true,
              "change": 1,
              "limit": 5,
              "isLive": false,
              "cd": 3600
            }
          ],
          "default": {
            "enable": true,
            "change": 1,
            "limit": 1,
            "isLive": true,
            "cd": 0
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
              "limit": 0,
              "isLive": false,
              "cd": 0
            },
            {
              "value": "20003",
              "fuzzy": false,
              "change": 200,
              "limit": 0,
              "isLive": false,
              "cd": 0
            }
          ],
          "default": {
            "enable": false,
            "change": 0,
            "limit": 0,
            "isLive": true,
            "cd": 0
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
              "isLive": true,
              "cd": 0
            }
          ],
          "default": {
            "enable": false,
            "change": 0,
            "limit": 0,
            "isLive": true,
            "cd": 0
          }
        }
      ]
    }
  ]
}
```

## 更新内容

### 2020年7月6日
1. 修复default规则有误的BUG
2. 优化代码

### 2020年7月5日
1. rule规则新增cd字段（单位：秒），用于设置某个规则的冷却时间，可实现例如：1小时内只能打一次卡

### 2020年6月29日
1. 修复无法变更直播间开播状态的BUG
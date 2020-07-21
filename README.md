## 🚀 斗鱼直播间积分系统
为你的直播间增加积分功能吧！

## 特点
1. 能够长久地在低端的服务器上运行
2. 热更新规则
3. 配置灵活
4. 安全完善的积分兑换机制

## 声明
1. 本项目为个人兴趣开发，无任何盈利手段
2. 水平有限，代码质量不高，互联网同好可参考使用，**尊重作者劳动成果，引用请注明出处**
3. 若本项目有帮助到你，还请不要吝啬star❤
4. 作者: 小淳 / QQ: 189964430 / 邮箱: 189964430@qq.com

## 使用步骤
1. 需要准备一台服务器，并且部署MySQL数据库
2. 按照下方要求手动设置数据库，
3. 将本软件上传至服务器，并在同一目录下创建config.json文件，config配置项见下方
4. 再创建rules.json文件，放到config.json设置的路径下
5. rules.json的配置项见下方，为了方便配置可访问[json在线编辑](http://json.la/online.html)对rules进行配置
6. 全部配置好后，运行软件即可

## 常见问题
### 如何快速部署服务?
- 不想配置规则，想直接运行看效果
1. 可以将项目中的dist文件夹上传到服务器，这个是项目的默认模板程序
2. 修改config.json中mysql_url
3. 部署mysql服务，使用下方的代码创建table，***并使得mysql数据库支持emoji***，这一步在下面mysql部分有详细步骤
4. 运行douyu-point

### 如何修改积分规则(rules.json)?
1. 可按照模板/标准配置，手动或者访问[json在线编辑](http://json.la/online.html)对积分规则(rules.json)进行修改
2. 如果不是直接修改原rules文件，请修改完毕后，保存rules.json到config.json设置的路径下
3. 在程序面板输入reload，然后回车即可重载积分规则，且立即生效
- 请注意，修改积分规则不会重置次数(limit)

### 重启软件会有什么影响？
1. 首先并不建议重启软件，除非你要修改config.json
2. 重启软件会造成limit重置，也就是每日积分增加的次数限制都会重置
3. 软件目前没有实现服务自动重启功能。若遇到数据库崩溃、斗鱼服务器崩溃等不可抗因素导致软件退出运行，还请自行处理后续。

### 如何安全的重启服务？
1. 保存limit数据，输入命令saveLimit
2. 输入exit提出程序
3. 重启程序
4. 输入命令loadLimit加载limit数据

### 命令有哪些？
1. 请输入help然后回车查看

### 如何设置/增加兑换的物品
1. 在数据库items表内按字段备注设置物品相关信息即可
2. 备注不清楚的可以查看下方mysql部署步骤中的代码
3. 图片地址请尽量使用图床保存

### 如何查看用户兑换物品的信息以及发货
1. 在数据库exchanges表内显示所有用户的兑换信息
2. 每个记录的status字段表示物品发货状态，0表示没有发货，1表示已发货
3. 发货请在后台手动设置status的值（0改成1）

-------------------------------

## 注意事项
1. 软件重启后limit会重置，也就是说每天限制的次数会重置
2. 软件每天0点自动重置limit
3. 修改rules后，请输入reload然后回车，此步操作可热更新积分规则
4. 如果遇到特殊情况可以输入resetLimit手动重置每天限制的次数
5. 本项目只完成了原型，更多的功能和HTTP接口请自行实现

-------------------------------
## MySQL数据库部署步骤
1. 创建table
```
points // 用户积分
items // 兑换物品
exchanges // 兑换记录
```
2. points表
```
CREATE TABLE `points` (
	`uid` BIGINT(20) NOT NULL DEFAULT '0',
	`id` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`point` BIGINT(20) NULL DEFAULT NULL,
	`update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`uid`),
	INDEX `index_id` (`id`),
	INDEX `index_point` (`point`)
)
COLLATE='utf8mb4_bin'
ENGINE=InnoDB
;

```

3. items表
```
CREATE TABLE `items` (
	`id` INT(11) NOT NULL DEFAULT '0' COMMENT '物品id',
	`name` VARCHAR(255) NULL DEFAULT NULL COMMENT '物品名称' COLLATE 'utf8mb4_unicode_ci',
	`description` VARCHAR(255) NULL DEFAULT NULL COMMENT '物品描述' COLLATE 'utf8mb4_unicode_ci',
	`pic` VARCHAR(255) NULL DEFAULT NULL COMMENT '物品图片地址',
	`price` BIGINT(20) NULL DEFAULT NULL COMMENT '物品价格',
	`num` INT(11) NULL DEFAULT NULL COMMENT '物品数量',
	`update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
)
COMMENT='兑换物品列表'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;

```

4. exchanges表
```
CREATE TABLE `exchanges` (
	`status` INT(11) NOT NULL DEFAULT '0' COMMENT '状态 0未处理 1已处理',
	`uid` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
	`id` VARCHAR(50) NULL DEFAULT NULL COMMENT '用户id' COLLATE 'utf8mb4_unicode_ci',
	`item_id` INT(11) NULL DEFAULT NULL COMMENT '物品id',
	`item_name` VARCHAR(255) NULL DEFAULT NULL COMMENT '物品名称' COLLATE 'utf8mb4_unicode_ci',
	`item_price` BIGINT(20) NULL DEFAULT NULL COMMENT '物品价格',
	`info` VARCHAR(255) NULL DEFAULT NULL COMMENT '兑换备注信息' COLLATE 'utf8mb4_unicode_ci',
	`update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	INDEX `index_id` (`id`),
	INDEX `index_item_id` (`item_id`),
	INDEX `index_uid` (`uid`),
	INDEX `index_status` (`status`)
)
COMMENT='兑换记录'
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
;



```
5. 支持emoji
[参考这篇文章设置](https://www.jianshu.com/p/48c3fbf28ea1)
```
1. 修改mysql配置文件my.cnf
[client]
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4

[mysqld]
character-set-client-handshake = FALSE
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init_connect='SET NAMES utf8mb4'

2. 重启mysql服务

3. 修改数据库的字符集
ALTER DATABASE douyu_point CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

4. golang连接mysql时带上以下参数
charset=utf8mb4&collation=utf8mb4_unicode_ci
例如: root:qianjiachun@/douyu_point?charset=utf8mb4&collation=utf8mb4_unicode_ci

```


## config.json
```
// 配置属性一律使用小写+下划线形式
// 该文件为项目启动的配置文件
// 用于指定监听的直播间，数据库地址，规则地址
{
	room_id: "", // 房间号，必须是真实房间号
	mysql_url: "", // mysql数据库地址，格式必须为golang连接mysql的url格式 注意带上参数charset=utf8mb4&collation=utf8mb4_unicode_ci以支持emoji
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

## API
当前系统默认对外提供的接口如下（端口默认27999）
### 查询用户积分
```
POST:
http://localhost:27999/douyu/point/5189167/query_by_uid
Header: {Content-Type: application/x-www-form-urlencoded}
Body: token=<斗鱼的token>
成功返回示例: {"error":0,"msg":"success","data":[{"uid":1825394,"id":"小淳丿","point":4,"update_time":"2020-07-07 22:31:32"}]}
```

### 查询兑换的物品列表
```
POST:
http://localhost:27999/douyu/point/5189167/query_item
Header: {Content-Type: application/x-www-form-urlencoded}
Body: token=<斗鱼的token>&offset=<limit是10>
成功返回示例: {"error":0,"msg":"success","data":[{"id":1,"name":"淳宝飞吻","description":"mua~","pic":"www.baidu.com","price":20,"num":99994,"update_time":"2020-07-07 22:31:32"},{"id":2,"name":"测试礼物","description":"222","pic":"www.baidu.com","price":1,"num":0,"update_time":"2020-07-07 18:04:14"},{"id":54,"name":"试试","description":"测试用的","pic":"www.baidu.com","price":1,"num":0,"update_time":"2020-07-07 22:14:19"},{"id":5777,"name":"免费的","description":"芜湖","pic":"www.baidu.com","price":0,"num":46,"update_time":"2020-07-07 22:13:50"}]}
```

### 兑换物品
```
POST:
http://localhost:27999/douyu/point/5189167/exchange
Header: {Content-Type: application/x-www-form-urlencoded}
Body: token=<斗鱼的token>&item_id=<兑换物品的id>&id=<用户的斗鱼昵称>&info=<备注信息,用于填写联系方式>
成功返回示例: {"error":0,"msg":"兑换成功","data":[{"uid":1825394,"id":"小淳丿","point":4,"update_time":"2020-07-07 23:19:16"}]}
```

### 查询用户物品兑换记录
```
POST:
http://localhost:27999/douyu/point/5189167/query_exchange
Header: {Content-Type: application/x-www-form-urlencoded}
Body: token=<斗鱼的token>&offset=<limit是10>
成功返回示例: {"error":0,"msg":"success","data":[{"status":0,"uid":1825394,"id":"小淳丿","item_id":1,"item_name":"测试礼物","item_pic":"","item_description":"","price":0,"info":"","update_time":""},{"status":1,"uid":1825394,"id":"小淳丿","item_id":1,"item_name":"测试礼物321","item_pic":"","item_description":"","price":0,"info":"","update_time":""}]}
```

## 更新内容

### 2020年7月21日
1. 新增limit状态序列化保存和加载命令，分别是saveLimit和loadLimit
2. 新增limit数据输出命令，logLimit
3. 积分榜返回数据个数由20增加到50 

### 2020年7月9日
1. 为接口增加ratelimit，限制IP访问频率
2. 删除exchanges表内description和pic字段
3. 修改了rateLimit规则
4. 接口返回按时间从新到旧排序

### 2020年7月8日
1. 数据库添加了支持emoji的规则，详细请看上方Mysql部署步骤
2. 新增查询用户物品兑换记录的接口

### 2020年7月7日
1. 新增物品兑换功能与一系列接口

### 2020年7月6日
1. 修复default规则有误的BUG
2. 优化代码
3. 新增查询用户积分接口

### 2020年7月5日
1. rule规则新增cd字段（单位：秒），用于设置某个规则的冷却时间，可实现例如：1小时内只能打一次卡

### 2020年6月29日
1. 修复无法变更直播间开播状态的BUG


--------------------------

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
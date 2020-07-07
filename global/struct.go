package global

/*
	内存中使用的字段
	用于底层积分变化逻辑
*/

type _Rules struct {
	Data []RuleRoot
}
type RuleRoot struct {
	Type   string `json:"type"`
	Fields []Field
}
type Field struct {
	Name    string `json:"name"`
	Cnt     string `json:"cnt"`
	Rules   []Rule
	Default Default
}
type Rule struct {
	Value  string `json:"value"`
	Fuzzy  bool   `json:"fuzzy"`
	Change int    `json:"change"`
	Limit  int    `json:"limit"`
	IsLive bool   `json:"isLive"`
	Cd     int    `json:"cd"`
}
type Default struct {
	Enable bool `json:"enable"`
	Change int  `json:"change"`
	Limit  int  `json:"limit"`
	IsLive bool `json:"isLive"`
	Cd     int  `json:"cd"`
}

type _Config struct {
	RoomID   string `json:"room_id"`
	MysqlURL string `json:"mysql_url"`
	Rules    string `json:"rules"`
}

type InfoUid struct {
	Count    int
	NextTime int64
}

/*
	数据库points表的字段
*/

type UserInfo struct { // points字段
	Uid        int    `json:"uid"`
	Id         string `json:"id"`
	Point      int    `json:"point"`
	UpdateTime string `json:"update_time"`
}

type UserInfoJson struct { // 返回给前端的结构
	Error int        `json:"error"`
	Msg   string     `json:"msg"`
	Data  []UserInfo `json:"data"`
}

/*
	数据库items表的字段
*/

type Item struct { // items字段
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Pic         string `json:"pic"`
	Price       int    `json:"price"`
	Num         int    `json:"num"`
	UpdateTime  string `json:"update_time"`
}

type ItemJson struct { // 返回给前端的结构
	Error int    `json:"error"`
	Msg   string `json:"msg"`
	Data  []Item `json:"data"`
}

type ItemExchangeJson struct { // 返回给前端的结构
	Error int        `json:"error"`
	Msg   string     `json:"msg"`
	Data  []UserInfo `json:"data"`
}

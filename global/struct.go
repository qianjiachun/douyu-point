package global

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

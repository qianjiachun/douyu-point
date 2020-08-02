package global

import (
	"database/sql"
	"github.com/yudeguang/ratelimit"
)

var Config _Config
var Rules _Rules
var IsLive bool
var DB *sql.DB
var List map[string]map[string]*InfoUid

var RateLimit *ratelimit.Rule

var GiftPrice map[string]int

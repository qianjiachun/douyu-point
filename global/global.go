package global

import "database/sql"

var Config _Config
var Rules _Rules
var IsLive bool
var DB *sql.DB
var List map[string]map[string]*InfoUid

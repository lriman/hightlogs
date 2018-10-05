package hlog

import (
	"time"
	"gopkg.in/mgo.v2"
)

// Log LEVEL's
const FATAL = 0
const ERROR = 1
const MAIN = 2
const EVENT = 3
const DETAILS = 4
const DEBUG = 5
const MINOR = 6

var db *mgo.Database
var MAX_LVL_DB = MINOR
var MAX_LVL_TV = MINOR

type Log struct {
	Info    string
	Date    time.Time
	Level   int
	Sender  string
}


func InitLogsDB(tmpdb *mgo.Database, max_lvl_db int, max_lvl_tv int){
	db = tmpdb
	MAX_LVL_DB = max_lvl_db
	MAX_LVL_TV = max_lvl_tv
}
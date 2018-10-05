package hlog

import (
	"log"
	"time"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"strings"
)


func SaveLog(sender string, lvl int, v ...interface{})  {

	info := fmt.Sprintln(v...)
	info = strings.TrimSuffix(info, "\n")

	if db != nil{
		l := Log{
			Info: info,
			Sender: sender,
			Level: lvl,
			Date: time.Now(),
		}

		if l.Level <= MAX_LVL_DB {
			db.C("logs").Insert(l)
		}
	}

	if lvl < MAX_LVL_TV {
		if lvl > 0 {
			if lvl > 1 {
				log.Println(sender, ":", info)
			} else {
				log.Println(sender, ":ERROR:", info)
			}
		} else {
			log.Fatalln(sender, ":FATAL:", info)
		}
	}
}

func GetLogs(sender string, lvl int) []Log {
	selector := bson.M{}
	if sender != "" {
		selector["sender"] = sender
	}
	if lvl > -1 {
		selector["level"] = bson.M{"$lte": lvl}
	}

	var logs []Log

	if err := db.C("logs").Find(selector).All(&logs); err == nil {
		return logs
	} else {
		return make([]Log, 0)
	}
}
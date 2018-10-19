package main

import (
	"github.com/jinzhu/gorm"
	"log/syslog"
	"log"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var globalDB *gorm.DB
var logger *log.Logger

func GetLogger() *log.Logger {
	return logger
}

func Init() {
	logger, _ = syslog.NewLogger(syslog.LOG_DEBUG, 1)

	db, err := gorm.Open("sqlite3", "bible.db")

	if err != nil {
		logger.Panic(err)
	}
	globalDB = db
}

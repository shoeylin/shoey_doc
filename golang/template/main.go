package main

import (
	"log"
	"os"
	"template/conf"
	"template/models"
)

func init() {

	// log "gopkg.in/natefinch/lumberjack.v2"
	conf.InitLog()

	// json config
	if !conf.InitConfig() {
		log.Println("read config error!!")
		os.Exit(1)
	}

	// mysql db
	for i, dbConnStr := range conf.Config.DBConnAry {
		models.GetDBConn(i+1, dbConnStr)

	}

	if models.ConnAry == nil {
		log.Println("connection db error!!")
		os.Exit(1)
	}

}

func main() {

	// recover panic log
	defer conf.PanicLog()

}

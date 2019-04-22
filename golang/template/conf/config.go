package conf

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Configuration struct {
	DBConnAry []string
}

var Config *Configuration

func InitLog() {
	log.SetOutput(io.MultiWriter(
		&lumberjack.Logger{
			Filename:   "./log/mergeKLine.log",
			MaxSize:    10, //megabytes
			MaxBackups: 10,
			MaxAge:     30, //days
			LocalTime:  true,
		},
		os.Stdout))
	log.SetFlags(log.Lmicroseconds | log.Ldate)
}

func InitConfig() bool {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = &Configuration{}
	err := decoder.Decode(&Config)
	if err != nil {
		log.Println("initConfig fail!!!!, err :", err)
		return false
	}
	return true
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	log.Printf("==> %s\n", string(buf[:n]))
}

func PanicLog() {
	// func() {
	if e := recover(); e != nil {
		PrintStack()
		os.Exit(1)
	}
	// }()
}

package models

import (
	"fmt"
	"log"
	"shoey_doc/golang/template/conf"

	"github.com/garyburd/redigo/redis"
)

func RedisHandle() {
	conn, err := redis.Dial("tcp", conf.Config.RedisConn)
	if err != nil {
		log.Println("cmd_redis connection error: ", err)
		return
	}
	// defer conn.Close()

	_, err = conn.Do("SELECT", "1")
	if err != nil {
		log.Println("select redis db error:", err)
		conn.Close()
		return
	}

	//參考 https://segmentfault.com/a/1190000007078961
	//set
	// bulk, err := redis.Values(conn.Do("smembers", "key"))
	//hash
	bulk, err := redis.Values(conn.Do("hmget", "HASH", "key"))
	if err != nil {
		log.Println("getRequest err: ", err)
	}

	var symbolAry []string

	for _, value := range bulk {
		fmt.Println("string(value.([]byte)", string(value.([]byte)))
		symbolAry = append(symbolAry, string(value.([]byte)))
	}

	fmt.Println(symbolAry)
}

package mongodb

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session

func Init() {
	Mongo("mongodb://192.168.6.85:27017")

	err := Session.Ping()
	if err != nil {
		fmt.Println(err)

	}
}

func Mongo(mgURL string) error {
	conn, err := mgo.Dial(mgURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	conn.DB("DB") // 帳密
	conn.SetMode(mgo.Monotonic, true)
	conn.SetPoolLimit(2000) // 限制連線數
	Session = conn
	return nil
}

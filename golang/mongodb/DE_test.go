package mongodb

import (
	"fmt"
)

type Test1 struct {
	id   string `bson:"_id" mapstructure:"_id"`
	name string `bson:"name"`
	img  string `bson:"avatar"`
}

type testmongo struct {
}

func (p *testmongo) InsertTest(InsertState []interface{}) {
	conn := Session.Copy()
	c := conn.DB("testDB").C("testC")
	for _, v := range InsertState {
		if err := c.Insert(v); err != nil {
			fmt.Println(err)
		}
	}
}

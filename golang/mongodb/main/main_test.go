package main

import (
	"shoey_doc/golang/mongodb"
)

func main() {
	ExampleMongoDBtest()
}

// ExampleMongoDBtest test description
func ExampleMongoDBtest() {
	// Test1
	mongodb.Init()
	// Output:
	//
	t := mongodb.Mgotest{}
	tStruct := &mongodb.Test1{id: 666666}
	t.gotestInsert(tStruct)

}

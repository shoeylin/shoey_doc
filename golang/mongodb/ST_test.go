package mongodb

type Mgotest struct {
	testmongo
}

func (p *Mgotest) gotestInsert(in *Test1) {

	var InsertState []interface{}
	InsertState = append(InsertState, in)
	p.InsertTest(InsertState)
}

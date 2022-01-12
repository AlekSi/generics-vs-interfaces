package interfaces

type Int int

func (Int) bsontype() {}

type String string

func (String) bsontype() {}

package interfaces

// Array is an ordered collection of values of any BSON type, including other arrays.
type Array struct {
	s []BSONType
}

func (*Array) bsontype() {}

func NewArray() *Array {
	return &Array{
		s: make([]BSONType, 0),
	}
}

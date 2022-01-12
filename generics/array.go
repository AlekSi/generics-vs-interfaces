package generics

// Array is an ordered collection of values of any BSON type, including other arrays.
type Array struct {
	s []any // can't use BSONType here
}

func NewArray() *Array {
	return &Array{
		s: make([]any, 0),
	}
}

package generics

type BSONType interface {
	int | string | *Array | *Document
}

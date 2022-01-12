package generics

type Document struct {
	m    map[string]any
	keys []string
}

// Document is an ordered collection of key/value pairs,
// where key is string and value is any BSON type, including other documents.
func NewDocument() *Document {
	return &Document{
		m:    make(map[string]any), // can't use BSONType here
		keys: make([]string, 0),
	}
}

func MakeDocument[T BSONType](key string, value T) *Document {
	d := NewDocument()
	DocumentSet(d, key, value)
	return d
}

func MakeDocument2[T1, T2 BSONType](key1 string, value1 T1, key2 string, value2 T2) *Document {
	d := MakeDocument(key1, value1)
	DocumentSet(d, key2, value2)
	return d
}

func MakeDocument3[T1, T2, T3 BSONType](key1 string, value1 T1, key2 string, value2 T2, key3 string, value3 T3) *Document {
	d := MakeDocument2(key1, value1, key2, value2)
	DocumentSet(d, key3, value3)
	return d
}

// Methods cannot have type parameters.
func DocumentSet[T BSONType](d *Document, key string, value T) {
	_, ok := d.m[key]
	if !ok {
		d.keys = append(d.keys, key)
	}
	d.m[key] = value
}

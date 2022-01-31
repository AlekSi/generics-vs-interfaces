package interfaces

import "fmt"

// Document is an ordered collection of key/value pairs,
// where key is string and value is any BSON type, including other documents.
type Document struct {
	m    map[String]BSONType
	keys []String
}

func (*Document) bsontype() {}

func NewDocument() *Document {
	return &Document{
		m:    make(map[String]BSONType),
		keys: make([]String, 0),
	}
}

func MakeDocument(pairs ...BSONType) (*Document, error) {
	l := len(pairs)
	if l%2 != 0 {
		return nil, fmt.Errorf("invalid number of arguments: %d", l)
	}

	doc := &Document{
		m:    make(map[String]BSONType, l/2),
		keys: make([]String, 0, l/2),
	}

	for i := 0; i < l; i += 2 {
		key, ok := pairs[i].(String)
		if !ok {
			return nil, fmt.Errorf("invalid key type: %T", pairs[i])
		}
		doc.Set(key, pairs[i+1])
	}

	return doc, nil
}

func (d *Document) Set(key String, value BSONType) {
	if _, ok := d.m[key]; !ok {
		d.keys = append(d.keys, key)
	}
	d.m[key] = value
}

func (d *Document) Has(key String) bool {
	_, ok := d.m[key]
	return ok
}

package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	d := NewDocument()

	// 1. It would be nice to disallow that at compile time.
	d.Set("qux", nil)

	// 2. It would be nice to write `d.Set("foo", 42)`` instead, using built-in types.
	d.Set("foo", Int(42))

	// 3. It would be nice to make it more efficient by avoiding boxing.
	d.Set("bar", String("baz"))

	d.Set("document", NewDocument())
	d.Set("array", NewArray())

	// 4. It would be nice to disallow that at compile time.
	_, err := MakeDocument(String("invalid number of arguments"))
	assert.Error(t, err)

	// 5. It would be nice to disallow that at compile time.
	_, err = MakeDocument(
		Int(42), String("invalid key type"),
	)
	assert.Error(t, err)

	// 6. That works with a single function.
	MakeDocument(String("foo"), Int(42), String("bar"), String("baz"))

	// 7. Works as expected.
	assert.True(t, d.Has("foo"))
	assert.Equal(t, Int(42), d.Get("foo"))
	assert.False(t, d.Has("baz"))
	assert.Equal(t, nil, d.Get("baz"))
}

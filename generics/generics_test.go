package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	d := NewDocument()

	// 1. Nice - does not compile.
	// DocumentSet(d, "qux", nil)

	// 2. Kinda nice, but method would be nicer.
	DocumentSet(d, "foo", 42)

	// 3. Nice, no boxing.
	DocumentSet(d, "bar", "baz")

	DocumentSet(d, "document", NewDocument())
	DocumentSet(d, "array", NewArray())

	// 4. Nice - does not compile.
	// MakeDocument("invalid number of arguments")

	// 5. Nice - does not compile.
	// MakeDocument(42, "invalid key type")

	// 6. That requires a separate function - there are not variadic template parameters.
	MakeDocument2("foo", 42, "bar", "baz")

	// 7. The caller should know the type.
	assert.True(t, d.Has("foo"))
	assert.Equal(t, 42, DocumentGet[int](d, "foo"))
	assert.Equal(t, "", DocumentGet[string](d, "foo"))
	assert.False(t, d.Has("baz"))
	assert.Equal(t, 0, DocumentGet[int](d, "baz"))
	assert.Equal(t, "", DocumentGet[string](d, "baz"))
}

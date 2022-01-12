package generics

import "testing"

func TestSet(t *testing.T) {
	d := NewDocument()

	// 1. Nice - does not compile.
	// DocumentSet(d, "qux", nil)

	// 2. Kinda nice, but method would be nicer.
	DocumentSet(d, "foo", 42)

	// 3. Nice, no boxing.
	DocumentSet(d, "foo", "baz")

	DocumentSet(d, "document", NewDocument())
	DocumentSet(d, "array", NewArray())

	// 4. Nice - does not compile.
	// MakeDocument("invalid number of arguments")

	// 5. Nice - does not compile.
	// MakeDocument(42, "invalid key type")

	// 6. That requires a separate function - there are not variadic template parameters.
	MakeDocument2("foo", 42, "bar", "baz")
}

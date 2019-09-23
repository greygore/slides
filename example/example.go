/*
Package example is just an empty example of what a package documentation can look like.

	// It can contain preformatted code blocks.

Sections

The package documentation can also contain sections
*/
package example

const (
	// MyFalseConstant is a generic constant value
	MyFalseConstant = false
)

var (
	// MyFoo is a generic variable value
	MyFoo = Foo{}
)

// Foo is a simple struct exported and unexported fields.
type Foo struct {
	Bar int // Every foo has a bar
	baz string
}

// NewFoo is a contructor func that returns a new Foo.
func NewFoo(bar int) Foo {
	return Foo{Bar: bar}
}

// Baz returns a Foo's baz value
func (f *Foo) Baz() string {
	return f.baz
}

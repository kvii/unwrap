package unwrap

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleUnwrap() {
	e1 := errors.New("e1")
	e2 := fmt.Errorf("e2 -> %w", e1)
	e3 := fmt.Errorf("e3 -> %w", e2)
	fmt.Println(Unwrap(e3))

	//Output: e1
}

func TestUnwrap(t *testing.T) {
	a := A{}
	b := B{a}
	c := C{b}

	w := Unwrap[I](c)
	if w != a {
		t.Fatalf("w != rw, %+v", w)
	}
}

type I interface {
	F()
}

type A struct{}

func (A) F() {}

type B struct{ I }

func (b B) Unwrap() I { return b.I }

type C struct{ I }

func (c C) Unwrap() I { return c.I }

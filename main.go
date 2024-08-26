package main

import (
	"fmt"
)

type X struct {
	V int
}

func (x X) outP() {
	fmt.Println(x.V)
}
func main() {

	var f AB = Foo{}
	g, ok := f.(BC)
	_ = g
	fmt.Println(ok)
}

type Foo struct{}

func (f Foo) A() {}
func (f Foo) B() {}
func (f Foo) C() {}

type AB interface {
	A()
	B()
}
type BC interface {
	B()
	C()
}

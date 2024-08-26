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
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.([]rune)
	fmt.Println(f, ok)
	g := i.(float64) // panic
	fmt.Println(g)
}

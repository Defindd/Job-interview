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
	sl1 := make([]int8, 5, 5)
	sl2 := make([]int64, 5, 5)
	sl1 = append(sl1, 1)
	sl2 = append(sl2, 2)
	fmt.Println(cap(sl1), cap(sl2)) //16,10-дело в классах и размерах памяти, которое может выделить алокатор памяти в го
	//https://go.dev/src/runtime/sizeclasses.go

}

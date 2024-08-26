package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 4)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	test(&mySlice)
	fmt.Println(mySlice)
}
func test(sl *[]int) {
	*sl = append(*sl, 1000)
}

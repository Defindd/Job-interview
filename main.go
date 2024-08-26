package main

import (
	"fmt"
)

func main() {
	dfas := "qwerty"
	runeDfas := []rune(dfas)
	for i := 0; i < len(runeDfas); i++ {
		fmt.Println('a', 'A')
		runeDfas[i] += 'A' - 'a'
	}
	fmt.Println(string(runeDfas))
}

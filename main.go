package main

import (
	"fmt"
	"sync"
)

type X struct {
	V int
}

func (x X) outP() {
	fmt.Println(x.V)
}
func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	f := map[int]int{}
	wg.Add(2)
	go func() {
		mu.Lock()
		defer func() {
			mu.Unlock()
			wg.Done()
		}()
		f[1] = 2

	}()
	go func() {
		mu.Lock()
		defer func() {
			mu.Unlock()
			wg.Done()
		}()
		f[1] = 100
	}()
	wg.Wait()
	fmt.Println(f[1])
}

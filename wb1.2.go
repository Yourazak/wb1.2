package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, n := range array {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			fmt.Println(count * count)
		}(n)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"github.com/LaBeloi/concurrency/sliceFuncs"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	start := time.Now()

	wg.Add(len(n))
	for i, s := range n {
		go func(slice []int, index int) {
			defer wg.Done()
			sliceFuncs.TotalOfSlice(slice, index)
		}(s, i)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

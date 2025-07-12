package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		mapa = make(map[int]string)
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			mapa[n] = "value" + strconv.Itoa(n)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(mapa)
}

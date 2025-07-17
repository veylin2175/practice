package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)

	// альтернатива через атомики
	var cnt int32
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&cnt, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(cnt)
}

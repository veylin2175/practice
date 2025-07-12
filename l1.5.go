package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var n time.Duration
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Введите положительное число")
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(ctx, time.Second*n)
	defer cancel()

	wg := &sync.WaitGroup{}
	toProcess, processed := make(chan int), make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, toProcess, processed)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			toProcess <- i
		}
		close(toProcess)
	}()

	go func() {
		wg.Wait()
		close(processed)
	}()

	var counter int
	for resultValue := range processed {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond * 10)
			processed <- value
		}
	}
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	resChan := make(chan int, len(slice))

	// Решение для версии Go 1.22 и позже (+ синхронизация вывода за счет каналов)
	for _, v := range slice {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resChan <- v * v
		}()
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for v := range resChan {
		fmt.Println(v)
	}

	fmt.Println("---------")

	// Решение для версии Go до 1.22 (избежание замыкания) (+ реализация без синхронизации)
	for _, v := range slice {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(v)
	}
	wg.Wait()
}

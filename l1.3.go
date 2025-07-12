package main

import (
	"fmt"
	"sync"
)

func main() {
	var wCount int
	_, err := fmt.Scan(&wCount)
	if err != nil || wCount <= 0 {
		fmt.Println("Некорректное число воркеров")
		return
	}
	wg := sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int), make(chan int)

	for i := 0; i < wCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workers(numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	for res := range processedNumbers {
		fmt.Println(res)
	}
}

func workers(toProcess <-chan int, processed chan<- int) {
	for {
		v, ok := <-toProcess
		if !ok {
			return
		}
		processed <- v
	}
}

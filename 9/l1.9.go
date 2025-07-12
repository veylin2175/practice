package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int), make(chan int) // числа до результата и числа после результата
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for _, v := range arr {
			numbersToProcess <- v
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)
}

func worker(toProcess <-chan int, processed chan<- int) {
	for {
		value, ok := <-toProcess
		if !ok {
			return
		}
		processed <- value * 2
	}
}

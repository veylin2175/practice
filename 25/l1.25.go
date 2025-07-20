package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.After(10 * time.Second)

	fmt.Println("Запуск остановки")
	sleep(timer)
	fmt.Println("Остановка завершена")
}

func sleep(timer <-chan time.Time) {
	select {
	case <-timer:
		return
	}
}

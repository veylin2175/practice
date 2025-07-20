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

	fmt.Println("Второй запуск")
	sleep2(time.Second * 10)
	fmt.Println("Второе время прошло")
}

func sleep(timer <-chan time.Time) {
	select {
	case <-timer:
		return
	}
}

// Еще одна реализация, с передачей time. Duration вместо канала (как в оригинальном time. Sleep)
func sleep2(timer time.Duration) {
	<-time.After(timer)
}

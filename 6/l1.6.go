package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	go exitByCondition() // 1 способ

	ctx, cancel := context.WithCancel(context.Background()) // 2 способ
	go exitViaContext(ctx)
	time.Sleep(time.Second)
	cancel()

	go exitViaGoexit() // 3 способ

	data := make(chan int) // 4 способ
	go exitChannels(data)
	data <- 1
	data <- 2
	close(data)

	go panicExit() // 5 способ

	go timeoutExit() // 6 способ

	go signalExit() // 7 способ
}

// exitByCondition - выход по условию (горутина завершается сама по себе по завершении)
func exitByCondition() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

// exitViaContext - выход через контекст (отмена контекста завершает работу горутины)
func exitViaContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст отменен")
			return
		default:
			fmt.Println("Горутина все еще работает")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// exitViaGoexit - принудительное завершение работы через runtime.Goexit()
func exitViaGoexit() {
	defer fmt.Println("Goexit также выполняет defer-функции")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 8 {
			runtime.Goexit()
		}
	}
}

// exitChannels закрывает через канал данных
func exitChannels(data chan int) {
	for v := range data {
		fmt.Println("Обработка значения", v)
	}
	fmt.Println("Канал закрыт")
}

// panicExit экстренно завершает горутину через панику
func panicExit() {
	defer fmt.Println("Действие перед паникой")
	panic("Экстренно завершение")
}

// timeoutExit завершение через таймаут (time. After)
func timeoutExit() {
	timer := time.After(3 * time.Second)
	for {
		select {
		case <-timer:
			fmt.Println("Время вышло")
			return
		default:
			fmt.Println("Работает...")
		}
	}
}

// signalExit завершает по получении сигнала ОС
func signalExit() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Ожидание сигнала")
	<-sig
	fmt.Println("Получен сигнал")
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Механизм Graceful Shutdown
	sigChan := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}

	// 1. Запускаем воркеры (для имитации работы горутин-воркеров)
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// 2. Получаем сигнал Ctrl+C
	signal.Notify(sigChan, syscall.SIGINT)

	sig := <-sigChan
	fmt.Printf("Пойман сигнал %v, идет завершение программы\n", sig)

	// 3. Через контекст передаем сигнал всем воркерам
	cancel()

	// 4. Через wg дожидаемся их завершения
	wg.Wait()

	// 5. main дождалась завершения работы всех воркеров
	fmt.Println("Программа корректно завершила свою работу")
}

func workers(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %v завершает работу\n", id)
			return
		default:
			fmt.Printf("Воркер %v еще работает...\n", id)
			time.Sleep(time.Second)
		}
	}
}

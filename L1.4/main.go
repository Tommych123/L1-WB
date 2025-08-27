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

// worker имитирует работу до тех пор, пока не придёт сигнал завершения через ctx
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d finish\n", id)
			return
		default:
			fmt.Printf("Worker %d work\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// создаём контекст с возможностью отмены, чтобы управлять завершением воркеров
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// канал для перехвата сигналов ОС (Ctrl+C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// запускаем 5 воркеров
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(ctx, i, wg)
	}

	// ждём сигнал прерывания
	<-signals
	fmt.Println("Get signal - finish")

	// отменяем контекст → это сигнал для всех воркеров завершиться
	cancel()

	// ждём завершения всех горутин
	wg.Wait()
	fmt.Println("All workers finish")
}

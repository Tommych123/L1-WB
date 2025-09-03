package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	mode := flag.String("mode", "condition", "Способ остановки горутины: condition | channel | context | goexit") // выбор режима по которому завершится горутина
	flag.Parse()

	switch *mode { // развилка где выбирается по какому режиму он завершится
	case "condition":
		stopByCondition()
	case "channel":
		stopByChannel()
	case "context":
		stopByContext()
	case "goexit":
		stopByGoexit()
	default:
		fmt.Println("Неверно введенный режим")
	}
}

// 1) Завершение по условию
func stopByCondition() {
	go func() {
		for i := 0; i < 5; i++ { // условие в виде счетчика
			fmt.Println("Работаем") // имитируем работу
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("Горутина завершилась по условию")
	}()
	time.Sleep(3 * time.Second) // задержка чтобы main не завершился раньше нашей горутины
}

// 2) Завершение через канал уведомления
func stopByChannel() {
	ch := make(chan struct{}) // создание канала для завершения
	go func() {
		for {
			select {
			case <-ch: // ждем закрытие канала
				fmt.Println("Горутина завершилась с помощью канала уведомления")
				return
			default:
				fmt.Println("Работа") // имитируем работу
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	close(ch)                   // закрываем канал
	time.Sleep(3 * time.Second) // задержка чтобы main не завершился раньше нашей горутины
}

// 3) Завершение через контекст
func stopByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // ждем сигнал отмены через контекст
				fmt.Println("Горутина завершилась через контекст")
				return
			default:
				fmt.Println("Работа") // имитируем работу
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	cancel()                    // завершаем контекст
	time.Sleep(3 * time.Second) // задержка чтобы main не завершился раньше нашей горутины
}

// 4) Завершение через runtime.Goexit()
func stopByGoexit() {
	go func() {
		defer fmt.Println("Дефер был завершен до выхода из горутины") // запускаем defer для завершения горутины
		fmt.Println("Горутина начала работу")                         // имитируем работу
		runtime.Goexit()                                              // Завершение горутины
		fmt.Println("Горутина должна завершиться до этого вывода")    // Эта строка не должна напечататься!
	}()
	time.Sleep(3 * time.Second) // задержка чтобы main не завершился раньше нашей горутины
}

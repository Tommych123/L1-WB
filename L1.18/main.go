package main

import (
	"fmt"
	"sync"
)

// Counter — структура счётчика с защитой для конкурентного доступа
type Counter struct {
	mu    sync.Mutex // mutex для безопасного изменения count
	count int        // текущее значение счётчика
}

// Increment безопасно увеличивает счётчик на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	counter := Counter{}  // создаём новый счётчик
	var wg sync.WaitGroup // группа ожидания для синхронизации горутин

	for i := 0; i < 5; i++ { // запускаем 5 горутин
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment() // безопасный инкремент
		}()
	}

	wg.Wait()                  // ждём завершения всех горутин
	fmt.Println(counter.count) // выводим итоговое значение счётчика
}

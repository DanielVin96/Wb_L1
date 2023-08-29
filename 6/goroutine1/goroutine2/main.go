package main

import (
	"context"
	"fmt"
	"time"
)

// Использование context для остановки горутины
func myFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // После вызова cancel() в горутину передается Done() и она завершается
			fmt.Println("goroutine stopped") // имитация случая если потеряно соединение сервера с клиентом
			return
		default:
			fmt.Println("goroutine is running") // если соелинение стабильно то горутина продолжает работу
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // Создание контекста
	go myFunc(ctx)                                          // Запуск горутины
	time.Sleep(3 * time.Second)
	cancel() // Через 3 сек выполняю метод cancel
	time.Sleep(time.Second)
}

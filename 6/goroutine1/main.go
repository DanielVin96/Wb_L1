package main

import (
	"fmt"
	"time"
)

func myFunc(stop chan bool) {
	for {
		select {
		case <-stop: // как только в канал приходит значение горутна прекращает работу
			fmt.Println("goroutine stopped")
			return
		default:
			fmt.Println("goroutine is running") // если из канала ничего не пришло, горутина продолжает работу
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stop := make(chan bool)
	go myFunc(stop) // Запуск горутины
	time.Sleep(3 * time.Second)
	stop <- true            // через 3 сек отправляем в канал значение
	time.Sleep(time.Second) // сон на сек чтобы горутина myFunc успела отработать прежде чем main завершит работу
}

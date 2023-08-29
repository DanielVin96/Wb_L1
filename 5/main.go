package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) { // Метод отправляющий данные в канал
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("отправлено:", i)
		time.Sleep(time.Second) // задержка между отправкой каждого значения
	}
	close(ch)
}

func getData(ch chan int) { // Метод получающий данные из канала
	for {
		value, ok := <-ch //в бесконечном цикле данные выводятся из канала, пока существуют в нем
		if ok {
			fmt.Println("получено:", value)
		} else {
			fmt.Println("Канал закрыт")
			return
		}
	}
}
func main() {
	N := 5 // время, по истечению кот-го программа завершается

	ch := make(chan int) // Создание канала

	go sendData(ch) //параллельное выполнение метода отправления
	go getData(ch)  // параллельное выполнение метода получения
	time.Sleep(time.Duration(N) * time.Second)
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int) // Cоздание каналов
	ch2 := make(chan int)

	go func(ch1 chan<- int) { // ан-я горутина, отправляющая числа в канал ch1
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}(ch1)

	go func(ch1 <-chan int, ch2 chan<- int) { // ан-я горутина принимающая числа в канал ch2 из ch1
		for num := range ch1 { // Цикл по ch1
			ch2 <- num * num // запись в канал сh2 чисел, возведенных в квадрат
		}
		close(ch2)
	}(ch1, ch2)

	for num := range ch2 { // Чтение из канала ch2
		fmt.Println(num)
	}

}

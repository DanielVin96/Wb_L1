package main

import (
	"fmt"
	"sync"
)

func SquareNumbers(num int, wg *sync.WaitGroup, Chan chan<- int) { // метод расчитывающий и записывающий квадраты чисел в канал
	square := num * num
	Chan <- square // помещение в канал квадрата числа
	wg.Done()
}

func Sum(Chan <-chan int) int { // Метод, расчитывающий суммирующий квадраты чисел из канала
	var wg sync.WaitGroup
	result := 0
	for num := range Chan {
		wg.Add(1)
		go func(num int) { // внутри цикла по каналу анониманые горутины выполняют вычисление
			result += num
			wg.Done()
		}(num)
	}
	wg.Wait()
	return result
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	SquareChan := make(chan int, len(nums)) // Буферизированный канал размером длины среза чисел

	for _, num := range nums { // В цикле идет вычисление квадратов параллельно
		wg.Add(1)
		go SquareNumbers(num, &wg, SquareChan)
	}

	wg.Wait()

	close(SquareChan)

	fmt.Println(Sum(SquareChan))
}

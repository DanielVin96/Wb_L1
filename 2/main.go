package main

import (
	"fmt"
	"sync"
)

func SquareNumbers(numbers []int) { // Метод, рассчитывающий квадраты чисел из среза
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) { // внутри цикла по срезу чисел анониманые горутины выполняют вычисление
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}

	wg.Wait()

}

func main() {
	nums := []int{2, 4, 6, 8, 10} // полученные данные
	SquareNumbers(nums)           // вызываю метод рассчитывания квадратов чисел
}

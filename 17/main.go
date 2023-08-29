package main

import "fmt"

func Binary(arr []int, item int) bool { // метод, реализующий алгоритм бинарного поиска
	low := 0             // первый элемент в срезе (его индекс)
	high := len(arr) - 1 // заносим в переменную последний элемент

	for low <= high { // идем циклом по срезу от начала до последнего элемента
		middle := (low + high) / 2 // находим индекс середины среза
		mid := arr[middle]         // помещаем середину в переменную
		if mid < item {            // если mid  меньше искомого числа
			low = middle + 1 // то сдвигаем mid на единицу и находим новый middle
		} else {
			high = middle - 1 // если больше сдвигаем mid в обратную сторону и находим новый middle
		}
	}
	if low == len(arr) || arr[low] != item { // если мы дошли до конца среза и не нашли item то вернем false
		return false
	}
	return true // если искомое число найдено вернем true
}
func main() {
	arr := []int{3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println((Binary(arr, 5)))
}

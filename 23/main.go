package main

import "fmt"

func DeleteElem(slice []int, i int) []int { // Метод удлаения элемента по индексу переданному в i
	for idx := range slice {
		if idx == i {
			slice = append(slice[:idx], slice[idx+1:]...) //Добавление в слайс элементов до и после idx, таким способом удаляем i-й элемент
		}
	}
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(DeleteElem(slice, 3))
}

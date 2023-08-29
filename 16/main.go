package main

import "fmt"

func Sort(arr []int) []int {
	l := len(arr)
	if l < 2 { // Если в срезе 1 элемент возвращаем срез, его не нужно отсортировывать
		return arr
	}
	left := make([]int, 0)  // создаю пустой срез для чисел меньше опорного эелемента
	right := make([]int, 0) // и срез для чисел больше опорного, чтобы в дальнейшем их отсортировать

	mid := arr[0] // опорный элемент

	for _, elem := range arr[1:] { // Идем циклом по arr

		if elem < mid {
			left = append(left, elem) // И добавляем в left элементы меньше опорного
		} else {
			right = append(right, elem) // И в right элементы больше опорного
		}
	}
	arr = append(Sort(left), mid)     // добавляем в срез arr отсторированный left и опорный эл-нт
	arr = append(arr, Sort(right)...) // Затем отсортированный right

	return arr // Возвращаем отсторированный срез

}

func main() {
	arr := []int{5, 4, 3, 6, 2, 1}
	fmt.Println(Sort(arr))
}

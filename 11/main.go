package main

import "fmt"

func Intersection(nums1, nums2 []int) []int {
	intersect := []int{}             // срез, куда будут включены схожие значения
	similarity := make(map[int]bool) // map, куда записыаются в кач-ве ключей знач-я 1-го множества

	for _, num := range nums1 { // запись в map значений первого множества в кач-ве ключей
		similarity[num] = true
	}

	for _, num := range nums2 { // цикл по 2-му множеству, для сравнения со значением, если знач-е true, значит значения схожи
		if similarity[num] {
			intersect = append(intersect, num) // добавление схожестей в пустой срез
		}
	}

	return intersect
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{4, 6, 1, 2, 9}
	fmt.Println(Intersection(nums1, nums2))
}

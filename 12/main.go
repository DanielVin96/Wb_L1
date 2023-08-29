package main

import "fmt"

func Availability(word string, myWords []string) bool { // метод, который проверяет наличие слова из множества в моем множестве
	for _, elem := range myWords {
		if word == elem {
			return true
		}
	}
	return false
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	myWords := []string{"car", "house"}

	for _, word := range words { // В цикле по множеству вызываю метод проверки наличия слова
		if !Availability(word, myWords) {
			myWords = append(myWords, word) // если слова нет в моем множестве то добавляю его свой набор
		}
	}
	fmt.Println(myWords)
}

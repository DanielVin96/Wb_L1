package main

import "fmt"

func Inverted(word string) string {
	empty := []rune{}
	runesWord := []rune(word)                  // Конвертирую строку в срез рун, для избежания проблем при итерации по Unicode
	for i := len(runesWord) - 1; i >= 0; i-- { // Иду циклом в обратном направлении и добавляю в пустой срез рун
		empty = append(empty, runesWord[i])
	}
	return string(empty) // Возвращаю обратную строку

}

func main() {
	fmt.Println(Inverted("главрыба"))
}

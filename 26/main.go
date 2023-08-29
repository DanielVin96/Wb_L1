package main

import "fmt"

func unique(str string) bool { // Метод, возвращающий bool  в зависимости есть или нет уникального значения
	runeStr := []rune(str)                // конвертирую строку в срез рун
	for i := 0; i < len(runeStr)-1; i++ { // Иду вложенными циклами по срезу,
		for j := i + 1; j < len(runeStr); j++ { // сравнивая каждый символ с последующими в строке
			if runeStr[i] == runeStr[j] { // если есть совпадение , значит строка не уникальна и верну false
				return false
			}
		}
	}
	return true // если совпадений не найдено, верну true
}

func main() {
	str := "abCdefAaf"
	fmt.Println(unique(str))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	empty := []string{}                    // Пустой срез, для помещения новой строки
	words := strings.Split(str, " ")       // создаю срез строк из str
	for i := len(words) - 1; i >= 0; i-- { // циклом в обртаном направлении по срезу строк записываю в пустую переменную
		empty = append(empty, words[i])
	}

	fmt.Println(strings.Join(empty, " ")) // Соединяю срез строк в строку

}

package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groupTemp := make(map[int][]float64)

	//делим текущую температуру на 10
	//округляем результат до ближайшего целого числа с помощью функции math.Floor
	// а затем умножаем результат на 10, чтобы получить число, кратное 10.
	for _, temp := range temperatures {
		idx := int(math.Floor(temp/10)) * 10          //  вычисляем индекс группы, в которую должна быть добавлена текущая температура.
		groupTemp[idx] = append(groupTemp[idx], temp) //  добавлем  текущую температуру
	}

	fmt.Println(groupTemp)
}

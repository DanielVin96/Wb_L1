package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) { // Метод, принимающий в кач-ве аргумента тип Duration в переменной d
	<-time.After(d) // Функция After отправляет текущее время по возвращаемому каналу после задержки, но в данной реал- ции это значение не используется
}

func main() {
	fmt.Println("Старт")
	Sleep(4 * time.Second)
	fmt.Println("Финиш")
}

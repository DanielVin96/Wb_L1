package main

import "fmt"

func main() {
	a := 1048576 // переменная 2 в 20 cтепени
	b := 1048576 // 2 в 20 cтепени

	sum := a + b // Сложение
	fmt.Printf("Результат сложения : %d\n", sum)

	sub := a - b // Вычитание
	fmt.Printf("Результат вычитания : %d\n", sub)

	multi := a * b // Умножение
	fmt.Printf("Результат умножения :%d\n", multi)

	div := a / b // Деление
	fmt.Printf("Результат деления : %d\n", div)
}

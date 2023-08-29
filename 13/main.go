package main

import "fmt"

func main() {
	num1, num2 := 2, 4
	num1, num2 = num2, num1 // меняю местами 2-е переменные без создания третьей
	fmt.Println(num1, num2)
}

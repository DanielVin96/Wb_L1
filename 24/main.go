package main

import (
	"fmt"
	"math"
)

type Points struct {
	x float64
	y float64
}

func MakePoints(x, y float64) *Points { // Метод, создающий структуру типа Points, в соответсвии с параметрами
	return &Points{x, y}
}

func Distance(p1, p2 Points) float64 { // Метод нахождения расстояния
	return math.Sqrt(math.Pow(p2.y-p1.y, 2) + math.Pow(p2.x-p1.x, 2)) //Применяется формула расстояния
}

func main() {
	p1 := *MakePoints(1, 2)
	p2 := *MakePoints(3, 5)
	dist := Distance(p1, p2)
	fmt.Println(dist)
}

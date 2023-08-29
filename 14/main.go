package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f interface{} = 5.0
	switch f.(type) {
	case int:
		fmt.Println("переменная типа int")
	case float64:
		fmt.Println("переменная типа float64")
	case string:
		fmt.Println("переменная типа string")
	case bool:
		fmt.Println("переменная типа bool")
	default:
		fmt.Println("переменная типа :", reflect.TypeOf(f))
	}
}

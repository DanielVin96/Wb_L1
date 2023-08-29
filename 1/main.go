package main

import "fmt"

type Person struct { // родительская структура
	Name       string
	Surname    string
	Patronymic string
}

func (p *Person) SayHello() string { // метод, который приветсвует
	return fmt.Sprintf("Привет, меня зовут %s %s %s", p.Surname, p.Name, p.Patronymic)
}

type Subsidiary struct { // дочерняя структура, наследующая Person
	Person
}

func main() {
	s := &Subsidiary{
		Person{
			Name:       "Петр",
			Surname:    "Смирнов",
			Patronymic: "Игоревич",
		},
	}
	fmt.Println(s.SayHello())
}

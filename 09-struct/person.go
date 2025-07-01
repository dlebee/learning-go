package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello", p.Name, "you are", p.Age, "years old")
}

func (p *Person) haveBirthday() {
	p.Age++
}

func (p Person) CombinedAge(secondPerson Person) int {
	return p.Age + secondPerson.Age
}

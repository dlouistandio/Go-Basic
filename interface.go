package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello2(value HasName) {
	fmt.Println("Hallo", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "David"}
	sayHello2(person)

	animal := Animal{Name: "Kucing"}
	sayHello2(animal)
}

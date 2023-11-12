package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var david Customer
	david.Name = "David Louis Tandio"
	david.Address = "Jl.Mahoni"
	david.Age = 18

	fmt.Println(david)
	fmt.Println(david.Address)

	louis := Customer{
		Name:    "Louis",
		Address: "Jl.Kaytang",
		Age:     20,
	}
	fmt.Println(louis)

	tandio := Customer{"Tandio", "Jl.cemara", 24}
	fmt.Println(tandio)

	david.sayHello("Louis")
}

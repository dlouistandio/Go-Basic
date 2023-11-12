package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Agus"
	sayHelloTo("David", "Louis")
	sayHelloTo(firstName, "Louis")
}

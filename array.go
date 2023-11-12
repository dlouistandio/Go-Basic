package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "David"
	name[1] = "Louis"
	name[2] = "Tandio"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(name))
	fmt.Println(len(values))
}

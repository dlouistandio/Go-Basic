package main

import "fmt"

func main() {

	counter := 0
	minus := 3

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	decrement := func() {
		fmt.Println("decrement")
		minus--
	}

	increment()
	increment()
	decrement()
	decrement()
	fmt.Println(counter)
	fmt.Println(minus)
}

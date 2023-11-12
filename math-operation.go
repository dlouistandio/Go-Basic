package main

import "fmt"

func main() {
	var result = 5 + 5
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i %= 3 // i = i % 2
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}

package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	fmt.Println(getHello("Agus"))
	fmt.Println(getHello(""))
}

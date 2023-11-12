package main

import "fmt"

func getFullName() (string, string, string) {
	return "David", "Louis", "Tandio"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

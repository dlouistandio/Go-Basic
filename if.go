package main

import "fmt"

func main() {
	name := "Louis T"

	if name == "David" {
		fmt.Println("Hallo David")
	} else if name == "Louis" {
		fmt.Println("Halo Louis")
	} else {
		fmt.Println("Ini siapa?")
	}

	//Short statement pada if
	if length := len(name); length <= 5 {
		fmt.Println("Nama kurang dari 5 huruf")
	} else {
		fmt.Println("Nama lebih dari 5 huruf")
	}
}

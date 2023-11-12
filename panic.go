package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApplications(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

func main() {
	runApplications(true)
}

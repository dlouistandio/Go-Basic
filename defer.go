package main

import "fmt"

func logging() {
	fmt.Println("selesai manggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Mulai app")
}

func main() {
	runApp()
}

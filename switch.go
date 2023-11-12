package main

import "fmt"

func main() {
	name := "David"

	switch name {
	case "David":
		fmt.Println("Halo David")
	case "Louis":
		fmt.Println("Halo Louis")
	default:
		fmt.Println("Ini siapa ya?")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama sangat panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah sesuai")
	}
}

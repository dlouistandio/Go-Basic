package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//pass by value
	//address1 := Address{"Banjarmasin", "Kalsel", "Indonesia"}
	//address2 := address1 //copy value
	//
	//address2.City = "Jogja"
	//fmt.Println(address1)value tidak berubah
	//fmt.Println(address2)value berubah jadi jogja

	//pass by reference
	address1 := Address{"Banjarmasin", "Kalsel", "Indonesia"}
	address2 := &address1 //pointer ke address1

	address2.City = "Jogja"
	fmt.Println(address1)
	fmt.Println(address2)
}

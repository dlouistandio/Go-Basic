package main

import "fmt"

type Alamat2 struct {
	City, Province, Country string
}

func main() {
	address1 := new(Alamat2)
	address2 := address1

	address2.Country = "Malay"

	fmt.Println(address1)
	fmt.Println(address2)
}

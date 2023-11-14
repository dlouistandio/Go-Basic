package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func main() {
	address1 := Alamat{"Banjarmasin", "Kalsel", "Indonesia"}
	address2 := &address1 //pointer

	address2.City = "Jogja"
	*address2 = Alamat{"Bandung", "Jabar", "Indo"} //asterisk operator
	fmt.Println(address1)
	fmt.Println(address2)

}

package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "David",
		"address": "Banjarmasin",
	}

	person["Nambah value"] = "Tambah key"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar go lang"
	book["author"] = "David"
	book["ups"] = "salah dong"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}

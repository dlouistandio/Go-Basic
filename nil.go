package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("David")

	if data == nil {
		fmt.Println("Map masih kosong")
	} else {
		fmt.Println(data["name"])
	}

}

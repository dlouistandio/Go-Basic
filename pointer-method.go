package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	david := Man{"David"}
	david.married()

	fmt.Println(david.Name)
}

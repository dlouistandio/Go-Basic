package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtp NoKTP = "3831237123"
	var marriedStatus Married = true
	fmt.Println(noKtp)
	fmt.Println(marriedStatus)
}

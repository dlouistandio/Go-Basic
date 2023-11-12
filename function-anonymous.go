package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("U are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		array := [...]string{name}
		for _, name := range array {
			switch {
			case name == "Root":
				return true
			case name == "Admin":
				return true
			}
		}
		return false
	}

	registerUser("Admin", blacklist)
	registerUser("David", blacklist)
}

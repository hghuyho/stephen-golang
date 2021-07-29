package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	huy := person{
		firstName: "Huy",
		lastName:  "HO",
		contactInfo: contactInfo{
			email:   "hghuyho@secman.vn",
			zipCode: 700000,
		},
	}
	huy.updateName("Mia")
	huy.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

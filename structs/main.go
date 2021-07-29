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
	// huy := person{lastName: "Huy", firstName: "HO"}
	// huy = person{lastName: "Huy", firstName: "HOO"}

	// var huy person
	// huy.firstName = "Huy"
	// huy.lastName = "HO"

	huy := person{
		firstName: "Huy",
		lastName:  "HO",
		contactInfo: contactInfo{
			email:   "hghuyho@secman.vn",
			zipCode: 700000,
		},
	}

	fmt.Println(huy)
	fmt.Printf("%+v", huy)

}

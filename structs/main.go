package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// huy := person{lastName: "Huy", firstName: "HO"}
	// huy = person{lastName: "Huy", firstName: "HOO"}

	var huy person
	huy.firstName = "Huy"
	huy.lastName = "HO"

	fmt.Println(huy)
	fmt.Printf("%+v", huy)

}

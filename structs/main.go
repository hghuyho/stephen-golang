package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	huy := person{lastName: "Huy", firstName: "HO"}
	// huy = person{lastName: "Huy", firstName: "HOO"}
	fmt.Println(huy)

}

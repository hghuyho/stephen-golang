package main

import "fmt"

func main() {
	ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenOrOdd string
	for i, n := range ns {
		if i%2 == 0 {
			evenOrOdd = "is even"
		} else {
			evenOrOdd = "is odd"
		}

		fmt.Println(n, evenOrOdd)

	}
}

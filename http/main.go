package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	bs := make([]byte, 100000)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
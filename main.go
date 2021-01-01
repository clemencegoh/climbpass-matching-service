package main

import (
	"net/http"
)

func main() {
	// Listen and Serve
	err := http.ListenAndServe(":8080", InitRouter("prod"))
	if err != nil {
		panic(err)
	}
}

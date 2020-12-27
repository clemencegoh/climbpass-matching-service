package main

import (
	"net/http"
)

func main() {
	// Listen and Serve
	http.ListenAndServe(":8080", InitRouter())
}

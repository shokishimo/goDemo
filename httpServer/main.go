package main

import (
	"fmt"
	"net/http"
)

func normalResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", normalResponse)

	fmt.Println("Server in running on port 80")
	http.ListenAndServe("localhost:80", nil)
}

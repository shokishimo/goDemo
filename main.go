package main

import (
	"fmt"
	"log"
	"net/http"
)

func normalResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", normalResponse)

	fmt.Println("Server in running on port 80")
	e := http.ListenAndServe("localhost:80", nil)
	if e != nil {
		log.Fatal("Error in running the server")
	}
}

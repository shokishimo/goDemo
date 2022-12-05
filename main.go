package main

import (
	"fmt"
	"log"
	"net/http"

	middleWare "github.com/shokishimo/goDemo/middleWare"
)

var count int = 0

func main() {
	// Static file like css and JavaScript
	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	http.HandleFunc("/", middleWare.IndexFunc)
	http.HandleFunc("/start", middleWare.StartQuiz)
	http.HandleFunc("/next", middleWare.NextQuestion)
	http.HandleFunc("/answer", middleWare.ShowAnswer)
	//http.HandleFunc("/finish")

	// Start the server
	fmt.Println("Server in running on port 80")
	if e := http.ListenAndServe("localhost:80", nil); e != nil {
		log.Fatal("Error in running the server")
	}
}

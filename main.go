package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	qDataModel "github.com/shokishimo/goDemo/models"
)

var count int = 0

func main() {
	// Static file like css and JavaScript
	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/start", startQuiz)
	http.HandleFunc("/next", nextQuestion)
	http.HandleFunc("/answer", showAnswer)
	//http.HandleFunc("/finish")

	// Start the server
	fmt.Println("Server in running on port 80")
	if e := http.ListenAndServe("localhost:80", nil); e != nil {
		log.Fatal("Error in running the server")
	}
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, 0)
}

func startQuiz(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/start" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "The HTTP Method is not supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	count = 0
	tmpl := template.Must(template.ParseFiles("static/quizQuestion.html"))
	// get random question
	data := qDataModel.Data{Question: "low-key", Answer: "「控え目、地味」「秘密」「ちょっと、なんとなく」"}

	tmpl.Execute(w, data)
}

func nextQuestion(w http.ResponseWriter, r *http.Request) {
}

func showAnswer(w http.ResponseWriter, r *http.Request) {

}

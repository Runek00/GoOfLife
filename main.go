package main

import (
	"log"
	"net/http"
	"text/template"
)

const size = 10

var arr = [size][size]byte{}
var arr2 = [size][size]byte{}

func main() {
	initGame()
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/box.js", serveBox)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func serveBox(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "box.js")
}

package main

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	//Serve static files
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", mux)
}
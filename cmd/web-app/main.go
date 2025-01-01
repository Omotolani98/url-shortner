package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ShowHomePage)
	log.Fatal(http.ListenAndServe(":8050", nil))
}

func ShowHomePage(writer http.ResponseWriter, _ *http.Request) {
	tmp, err := template.ParseFiles("internal/views/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmp.Execute(writer, nil); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

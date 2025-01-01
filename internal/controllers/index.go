package controllers

import (
	"html/template"
	"net/http"
)

func ShowIndex(writer http.ResponseWriter, _ *http.Request) {
	tmp, err := template.ParseFiles("internal/views/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmp.Execute(writer, nil); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
